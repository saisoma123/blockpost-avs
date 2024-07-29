package operator

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	blockPostServiceManager "operator/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Layr-Labs/incredible-squaring-avs/metrics"
	"github.com/Layr-Labs/incredible-squaring-avs/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const AVS_NAME = "blockpost-avs"
const SEM_VER = "0.0.1"

type SignedMessage struct {
	MessageId *big.Int
	Message   string
	Signature []byte
}

type ValidatedMessage struct {
	MessageId   *big.Int
	Message     string
	messageHash []byte
}

// Operator type from Incredible Squaring AVS
type Operator struct {
	config                      types.NodeConfig
	logger                      logging.Logger
	ethClient                   eth.Client
	metricsReg                  *prometheus.Registry
	metrics                     metrics.Metrics
	nodeApi                     *nodeapi.NodeApi
	avsWriter                   *avsregistry.AvsRegistryChainWriter
	avsReader                   *avsregistry.AvsRegistryChainReader
	avsSubscriber               *avsregistry.AvsRegistryChainSubscriber
	eigenlayerReader            sdkelcontracts.ELReader
	eigenlayerWriter            sdkelcontracts.ELWriter
	blsKeypair                  *bls.KeyPair
	ecdsaKey                    *ecdsa.PrivateKey
	operatorId                  sdktypes.OperatorId
	operatorAddr                common.Address
	skWallet                    wallet.Wallet
	txMgr                       txmgr.TxManager
	newTaskCreatedChan          chan *blockPostServiceManager.BindingsMessageSubmitted
	blockPostServiceManagerAddr common.Address
}

// Operator config function from Incredible Squaring AVS
func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	var logLevel logging.LogLevel
	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}
	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}
	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient eth.Client
	if c.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewInstrumentedClient(c.EthWsUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	} else {
		ethRpcClient, err = eth.NewClient(c.EthRpcUrl)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewClient(c.EthWsUrl)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}
	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 c.EthRpcUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress,
		AvsName:                    AVS_NAME,
		PromMetricsIpPortAddress:   c.EigenMetricsIpPortAddress,
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		c.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return nil, err
	}
	sdkClients, err := clients.BuildAll(chainioConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, common.HexToAddress(c.OperatorAddress))

	avsWriter, err := avsregistry.BuildAvsRegistryChainWriter(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress), logger, ethRpcClient,
		txMgr)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := avsregistry.BuildAvsRegistryChainReader(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := avsregistry.BuildAvsRegistryChainSubscriber(common.HexToAddress(c.AVSRegistryCoordinatorAddress), ethWsClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	operator := &Operator{
		config:                      c,
		logger:                      logger,
		metricsReg:                  reg,
		metrics:                     avsAndEigenMetrics,
		nodeApi:                     nodeApi,
		ethClient:                   ethWsClient,
		avsWriter:                   avsWriter,
		avsReader:                   avsReader,
		avsSubscriber:               avsSubscriber,
		eigenlayerReader:            sdkClients.ElChainReader,
		eigenlayerWriter:            sdkClients.ElChainWriter,
		blsKeypair:                  blsKeyPair,
		ecdsaKey:                    operatorEcdsaPrivateKey,
		operatorAddr:                common.HexToAddress(c.OperatorAddress),
		skWallet:                    skWallet,
		txMgr:                       txMgr,
		newTaskCreatedChan:          make(chan *blockPostServiceManager.BindingsMessageSubmitted),
		blockPostServiceManagerAddr: common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		operatorId:                  [32]byte{0}, // this is set below

	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	operator.operatorId = operatorId
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil

}

// This computes the message hash, used for later validation
func ComputeMessageHash(message string) []byte {
	hash := crypto.Keccak256Hash([]byte(message))
	return hash.Bytes()
}

// Checks to see if the given hash is correct
func VerifyMessageIntegrity(message string, expectedHash []byte) bool {
	actualHash := ComputeMessageHash(message)
	return bytes.Equal(actualHash, expectedHash)
}

// Appends ethereum signature to hash for signature
func toEthSignedMessageHash(messageHash []byte) []byte {
	prefix := "\x19Ethereum Signed Message:\n32"
	prefixedHash := append([]byte(prefix), messageHash...)
	ethSignedMessageHash := crypto.Keccak256Hash(prefixedHash)
	return ethSignedMessageHash.Bytes()
}

// Processes the message from the event emission and computes the hash
func (o *Operator) ProcessNewMessageLog(newMessageLog *blockPostServiceManager.BindingsMessageSubmitted) ValidatedMessage {
	o.logger.Debug("Received new message", "message", newMessageLog)
	o.logger.Info("Received new message",
		"messageId", newMessageLog.MessageId,
		"message", newMessageLog.Message,
	)

	messageHash := ComputeMessageHash(newMessageLog.Message)

	validatedMessageStruct := ValidatedMessage{
		MessageId:   newMessageLog.MessageId,
		Message:     newMessageLog.Message,
		messageHash: messageHash,
	}
	return validatedMessageStruct
}

// Creates the signature for on chain verification of the operator sent message to the ServiceManager
func (o *Operator) SignValidatedMessage(validatedMessage ValidatedMessage) *SignedMessage {

	// Checks to see if the message was tampered with in transit
	isValid := VerifyMessageIntegrity(validatedMessage.Message, validatedMessage.messageHash)
	if !isValid {
		o.logger.Fatal("Message integrity verification failed", "messageId", validatedMessage.MessageId)
		return nil
	}

	// Appends eth signature to hash
	ethHash := toEthSignedMessageHash(validatedMessage.messageHash)

	// Sign the hash with the ECDSA private key
	signature, err := crypto.Sign(ethHash, o.ecdsaKey)
	if err != nil {
		o.logger.Fatal("Failed to sign validated message", "err", err)
		return nil
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	signedMessage := &SignedMessage{
		MessageId: validatedMessage.MessageId,
		Message:   validatedMessage.Message,
		Signature: signature,
	}

	o.logger.Debug("Signed validated message", "signedMessage", signedMessage)
	return signedMessage
}

// Sends the validated and signed message to the ServiceManager via the operator contract
func (o *Operator) SubmitSignedMessageToBlockchain(signedMessage *SignedMessage, bindings *blockPostServiceManager.Bindings) error {
	auth, err := bind.NewKeyedTransactorWithChainID(o.ecdsaKey, big.NewInt(17000))
	tx, err := bindings.StoreValidatedMessage(auth, signedMessage.MessageId, signedMessage.Message, signedMessage.Signature)
	if err != nil {
		o.logger.Error("Failed to create transaction for signed message", "err", err)
		return err
	}

	o.logger.Info("Submitted signed message to blockchain", "txHash", tx.Hash().Hex())
	return nil
}

// This starts a continously running loop in which the operator will validate
// and sign a message any time a SubmitMessage event is emitted from the ServiceManager
func (o *Operator) StartMessageProcessing(ctx context.Context) error {
	messageChan := make(chan *blockPostServiceManager.BindingsMessageSubmitted)

	// Creates bindings to deployed service manager for subscribing to event listening
	bindings, err := blockPostServiceManager.NewBindings(common.HexToAddress("0x52acEa39aBe44B5b5598279Ff507dF5721c2A616"), o.ethClient)
	if err != nil {
		o.logger.Fatalf("Failed to instantiate bindings for event watching: %v", err)
	}

	//o.RegisterOperatorWithAvs(o.ecdsaKey)

	messageIds := []*big.Int{}

	// Get the current block number
	currentBlock, err := o.ethClient.BlockNumber(ctx)
	if err != nil {
		o.logger.Fatal("Failed to get current block number", "err", err)
		return err
	}

	// The block range and continous changing of it is needed for event watching
	// to function properly, as allowing for too much blocks can crash the subscriber
	var fromBlock uint64 = currentBlock - 10000
	blockRange := uint64(5000)

	watchOpts := &bind.WatchOpts{
		Start:   &fromBlock,
		Context: ctx,
	}

	// Subscribed to event listening for submitted messages on-chain
	sub, err := bindings.WatchMessageSubmitted(watchOpts, messageChan, messageIds)
	if err != nil {
		o.logger.Fatal("Failed to subscribe to message events", "err", err)
		return fmt.Errorf("failed to subscribe to message events: %v", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			o.logger.Info("Context done, exiting")
			return nil
		case newMessageLog := <-messageChan:
			if newMessageLog == nil {
				o.logger.Fatal("Received nil message log")
				continue
			}
			o.logger.Info("Received new message log")

			// Process the new message log
			validatedMessage := o.ProcessNewMessageLog(newMessageLog)
			if &validatedMessage == nil {
				o.logger.Fatal("Validated message is nil")
				continue
			}

			// Sign the validated message
			signedMessage := o.SignValidatedMessage(validatedMessage)
			if signedMessage == nil {
				o.logger.Fatal("Failed to sign validated message")
				continue
			}

			/// Submits the message to the ServiceManager contract
			err = o.SubmitSignedMessageToBlockchain(signedMessage, bindings)
			if err != nil {
				o.logger.Fatal("Failed to submit signed message to blockchain", "err", err)
				continue
			}

		case err := <-sub.Err():
			o.logger.Error("Subscription error", "err", err)
			sub.Unsubscribe()

			watchOpts.Start = &fromBlock
			sub, err = bindings.WatchMessageSubmitted(watchOpts, messageChan, messageIds)
			if err != nil {
				return fmt.Errorf("failed to resubscribe to message events: %v", err)
			}
			continue
		}

		fromBlock += blockRange
		if fromBlock > currentBlock {
			currentBlock, err = o.ethClient.BlockNumber(ctx)
			if err != nil {
				o.logger.Fatal("Failed to get current block number", "err", err)
				return err
			}
		}
		time.Sleep(2 * time.Second)
	}
}
