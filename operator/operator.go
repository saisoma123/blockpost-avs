package operator

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func (o *Operator) ProcessNewMessageLog(newMessageLog *blockpost.ContractBlockPostMessageSubmitted) *blockpost.ContractBlockPostValidatedMessage {
	o.logger.Debug("Received new message", "message", newMessageLog)
	o.logger.Info("Received new message",
		"messageId", newMessageLog.MessageId,
		"message", newMessageLog.Message,
	)

	validatedMessage := newMessageLog.Message // Will add validation logic

	validatedMessageStruct := &blockpost.ContractBlockPostValidatedMessage{
		MessageId: newMessageLog.MessageId,
		Message:   validatedMessage,
	}
	return validatedMessageStruct
}

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

func (o *Operator) StartMessageProcessing(ctx context.Context) error {
	messageChan := make(chan *blockpost.MessageSubmitted)

	sub, err := o.blockpostContract.WatchMessageSubmitted(&bind.WatchOpts{}, messageChan)
	if err != nil {
		return fmt.Errorf("failed to subscribe to message events: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case newMessageLog := <-messageChan:
			// Process the new message log
			validatedMessage := o.ProcessNewMessageLog(newMessageLog)

			// Sign the validated message
			signedMessage, err := o.SignValidatedMessage(validatedMessage)
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
			sub, err = o.blockpostContract.WatchMessageSubmitted(&bind.WatchOpts{}, messageChan)
			if err != nil {
				return fmt.Errorf("failed to resubscribe to message events: %v", err)
			}
		}
	}
}
