package operator

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	blockPostServiceManager "operator/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/metrics"
	"github.com/Layr-Labs/incredible-squaring-avs/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
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
	Signature *bls.Signature
}

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient eth.Client
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg       *prometheus.Registry
	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader sdkelcontracts.ELReader
	eigenlayerWriter sdkelcontracts.ELWriter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
	operatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *blockPostServiceManager.BindingsMessageSubmitted
	// needed when opting in to avs (allow this service manager contract to slash operator)
	blockPostServiceManagerAddr common.Address
}

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
	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
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

	avsWriter, err := chainio.BuildAvsWriter(
		txMgr, common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress), ethRpcClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.BuildAvsSubscriber(common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress), ethWsClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
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
		ethClient:                   ethRpcClient,
		avsWriter:                   avsWriter,
		avsReader:                   avsReader,
		avsSubscriber:               avsSubscriber,
		eigenlayerReader:            sdkClients.ElChainReader,
		eigenlayerWriter:            sdkClients.ElChainWriter,
		blsKeypair:                  blsKeyPair,
		operatorAddr:                common.HexToAddress(c.OperatorAddress),
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

func (o *Operator) ProcessNewMessageLog(newMessageLog *blockPostServiceManager.BindingsMessageSubmitted) *blockPostServiceManager.BindingsMessageValidated {
	o.logger.Debug("Received new message", "message", newMessageLog)
	o.logger.Info("Received new message",
		"messageId", newMessageLog.MessageId,
		"message", newMessageLog.Message,
	)

	validatedMessage := newMessageLog.Message // Will add validation logic

	validatedMessageStruct := &blockPostServiceManager.BindingsMessageValidated{
		MessageId: newMessageLog.MessageId,
		Message:   validatedMessage,
	}
	return validatedMessageStruct
}

func (o *Operator) SignValidatedMessage(validatedMessage *blockPostServiceManager.BindingsMessageValidated) *SignedMessage {
	messageHash := crypto.Keccak256Hash([]byte(validatedMessage.Message))

	signature := o.blsKeypair.SignMessage(messageHash)

	signedMessage := &SignedMessage{
		MessageId: validatedMessage.MessageId,
		Message:   validatedMessage.Message,
		Signature: signature,
	}

	o.logger.Debug("Signed validated message", "signedMessage", signedMessage)
	return signedMessage
}

func (o *Operator) SubmitSignedMessageToBlockchain(signedMessage *SignedMessage) error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	auth := bind.NewKeyedTransactor(privateKey)
	bindings, err := blockPostServiceManager.NewBindings(o.blockPostServiceManagerAddr, o.ethClient)
	tx, err := bindings.StoreValidatedMessage(auth, signedMessage.MessageId, signedMessage.Message, signedMessage.Signature.Serialize())
	if err != nil {
		o.logger.Error("Failed to submit signed message to blockchain", "err", err)
		return err
	}

	o.logger.Info("Submitted signed message to blockchain", "txHash", tx.Hash().Hex())
	return nil
}

func (o *Operator) StartMessageProcessing(ctx context.Context) error {
	messageChan := make(chan *blockPostServiceManager.BindingsMessageSubmitted)

	bindings, err := blockPostServiceManager.NewBindings(o.blockPostServiceManagerAddr, o.ethClient)
	if err != nil {
		o.logger.Fatalf("Failed to instantiate bindings for event watching: %v", err)
	}

	watchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}

	messageIds := []*big.Int{}

	sub, err := bindings.WatchMessageSubmitted(watchOpts, messageChan, messageIds)

	for {
		select {
		case <-ctx.Done():
			return nil
		case newMessageLog := <-messageChan:
			// Process the new message log
			validatedMessage := o.ProcessNewMessageLog(newMessageLog)

			// Sign the validated message
			signedMessage := o.SignValidatedMessage(validatedMessage)
			if err != nil {
				o.logger.Fatal("Failed to sign validated message", "err", err)
				continue
			}

			// Submit the signed message to the blockchain
			err = o.SubmitSignedMessageToBlockchain(signedMessage)
			if err != nil {
				o.logger.Fatal("Failed to submit signed message to blockchain", "err", err)
				continue
			}

		case err := <-sub.Err():
			o.logger.Error("Subscription error", "err", err)
			sub.Unsubscribe()
			sub, err = bindings.WatchMessageSubmitted(watchOpts, messageChan, messageIds)
			if err != nil {
				return fmt.Errorf("failed to resubscribe to message events: %v", err)
			}
		}
	}
}
