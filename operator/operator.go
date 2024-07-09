package operator

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func (o *Operator) SignValidatedMessage(validatedMessage *blockpost.ContractBlockPostValidatedMessage) (*blockpost.ContractBlockPostSignedMessage, error) {
	messageHash := crypto.Keccak256Hash([]byte(validatedMessage.Message))

	// Use ECDSA keypair for signing the message
	signature, err := crypto.Sign(messageHash.Bytes(), o.ecdsaKeypair)
	if err != nil {
		return nil, err
	}

	signedMessage := &blockpost.ContractBlockPostSignedMessage{
		MessageId: validatedMessage.MessageId,
		Message:   validatedMessage.Message,
		Signature: signature,
	}

	o.logger.Debug("Signed validated message", "signedMessage", signedMessage)
	return signedMessage, nil
}

func (o *Operator) SubmitSignedMessageToBlockchain(signedMessage *blockpost.ContractBlockPostSignedMessage) error {
	auth := bind.NewKeyedTransactor(o.ecdsaKeypair)
	tx, err := o.blockpostContract.StoreValidatedMessage(auth, signedMessage.MessageId, signedMessage.Message, signedMessage.Signature)
	if err != nil {
		o.logger.Error("Failed to submit signed message to blockchain", "err", err)
		return err
	}

	o.logger.Info("Submitted signed message to blockchain", "txHash", tx.Hash().Hex())
	return nil
}
