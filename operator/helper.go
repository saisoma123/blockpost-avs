package operator

import (
	"context"
	"crypto/ecdsa"

	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
)

// This registers the operator to the EigenDA AVS for security and validation purposes
func (o *Operator) RegisterOperatorWithAvs(
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
) error {
	// Define parameters for registration
	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(1)}
	socket := "https://ethereum-holesky-rpc.publicnode.com"

	// Register the operator
	_, err := o.avsWriter.RegisterOperator(
		context.Background(),
		operatorEcdsaKeyPair,
		o.blsKeypair, quorumNumbers, socket,
	)
	if err != nil {
		o.logger.Errorf("Unable to register operator with AVS registry coordinator %v", err)
		return err
	}
	o.logger.Infof("Registered operator with AVS registry coordinator.")

	return nil
}
