// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retrieveMessage\",\"inputs\":[{\"name\":\"_messageId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeValidatedMessage\",\"inputs\":[{\"name\":\"_messageId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitMessage\",\"inputs\":[{\"name\":\"_message\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSubmitted\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageValidated\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561001057600080fd5b506040516126d53803806126d583398101604081905261002f9161013c565b6001600160a01b0380841660c052808316608052811660a052828282610053610064565b5050600160c9555061018992505050565b600054610100900460ff16156100d05760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015610122576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461013957600080fd5b50565b60008060006060848603121561015157600080fd5b835161015c81610124565b602085015190935061016d81610124565b604085015190925061017e81610124565b809150509250925092565b60805160a05160c0516124bb61021a600039600081816101c601528181610bd601528181610e710152610ef00152600081816106cc0152818161081e015281816108b501528181610ff00152818161116a01526112090152600081816104f7015281816105860152818161060601528181610b8201528181610e1501528181610f2b01526110c501526124bb6000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c8063886f1195116100ad578063a98fb35511610071578063a98fb35514610283578063e481af9d14610296578063f2cb613a1461029e578063f2fde38b146102b1578063fabc1cbc146102c457600080fd5b8063886f1195146102195780638da5cb5b1461022c5780639926ee7d1461023d5780639ba8fd0914610250578063a364f4da1461027057600080fd5b80635ac86ab7116100f45780635ac86ab71461017f5780635c975abb146101b25780636b3aa72e146101c4578063708b34fe146101fe578063715018a61461021157600080fd5b806310d67a2f14610126578063136439dd1461013b57806333cfb7b71461014e578063595c6a6714610177575b600080fd5b610139610134366004611ce1565b6102d7565b005b610139610149366004611d05565b610393565b61016161015c366004611ce1565b6104d2565b60405161016e9190611d1e565b60405180910390f35b610139610982565b6101a261018d366004611d7a565b609854600160ff9092169190911b9081161490565b604051901515815260200161016e565b6098545b60405190815260200161016e565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b03909116815260200161016e565b6101b661020c366004611e3a565b610a49565b610139610b63565b6097546101e6906001600160a01b031681565b6033546001600160a01b03166101e6565b61013961024b366004611e77565b610b77565b61026361025e366004611d05565b610c43565b60405161016e9190611f73565b61013961027e366004611ce1565b610e0a565b610139610291366004611e3a565b610ed1565b610161610f25565b6101396102ac366004611f86565b6112d2565b6101396102bf366004611ce1565b6114f9565b6101396102d2366004611d05565b61156f565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561032a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034e9190611ff3565b6001600160a01b0316336001600160a01b0316146103875760405162461bcd60e51b815260040161037e90612010565b60405180910390fd5b610390816116cb565b50565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156103db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ff919061205a565b61041b5760405162461bcd60e51b815260040161037e9061207c565b609854818116146104945760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161037e565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa15801561053e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056291906120c4565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156105cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f191906120dd565b90506001600160c01b038116158061068b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610662573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106869190612106565b60ff16155b156106a757505060408051600081526020810190915292915050565b60006106bb826001600160c01b03166117c2565b90506000805b8251811015610787577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061070b5761070b612123565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa15801561074f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077391906120c4565b61077d908361214f565b91506001016106c1565b5060008167ffffffffffffffff8111156107a3576107a3611d97565b6040519080825280602002602001820160405280156107cc578160200160208202803683370190505b5090506000805b84518110156109755760008582815181106107f0576107f0612123565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610865573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088991906120c4565b905060005b8181101561096a576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610903573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109279190612162565b6000015186868151811061093d5761093d612123565b6001600160a01b03909216602092830291909101909101528461095f816121d2565b95505060010161088e565b5050506001016107d3565b5090979650505050505050565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156109ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ee919061205a565b610a0a5760405162461bcd60e51b815260040161037e9061207c565b600019609881905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000609854600014610a9d5760405162461bcd60e51b815260206004820152601c60248201527f5061757361626c653a20636f6e74726163742069732070617573656400000000604482015260640161037e565b6000825111610aee5760405162461bcd60e51b815260206004820152601760248201527f4d6573736167652063616e6e6f7420626520656d707479000000000000000000604482015260640161037e565b60c980549081906000610b00836121d2565b9091555050600081815260cc60205260409081902080546001600160a01b031916331790555181907f880e27ad36e62ea827e957f86d06f355728f6630b504a86ca8c371962f929ff090610b55908690611f73565b60405180910390a292915050565b610b6b611885565b610b7560006118df565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610bbf5760405162461bcd60e51b815260040161037e906121eb565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610c0d9085908590600401612263565b600060405180830381600087803b158015610c2757600080fd5b505af1158015610c3b573d6000803e3d6000fd5b505050505050565b600081815260ca6020526040812080546060929190610c61906122ae565b905011610cb05760405162461bcd60e51b815260206004820152601960248201527f4d65737361676520494420646f6573206e6f7420657869737400000000000000604482015260640161037e565b600082815260cb602052604090205460ff16610d065760405162461bcd60e51b815260206004820152601560248201527413595cdcd859d9481b9bdd081d985b1a59185d1959605a1b604482015260640161037e565b600082815260cc60205260409020546001600160a01b03163314610d6c5760405162461bcd60e51b815260206004820152601960248201527f54686973206973206e6f7420796f7572206d6573736167652100000000000000604482015260640161037e565b600082815260ca602052604090208054610d85906122ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610db1906122ae565b8015610dfe5780601f10610dd357610100808354040283529160200191610dfe565b820191906000526020600020905b815481529060010190602001808311610de157829003601f168201915b50505050509050919050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e525760405162461bcd60e51b815260040161037e906121eb565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610eb657600080fd5b505af1158015610eca573d6000803e3d6000fd5b5050505050565b610ed9611885565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610e9c908490600401611f73565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fab9190612106565b60ff16905080600003610fcc57505060408051600081526020810190915290565b6000805b8281101561107757604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa15801561103f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061106391906120c4565b61106d908361214f565b9150600101610fd0565b5060008167ffffffffffffffff81111561109357611093611d97565b6040519080825280602002602001820160405280156110bc578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611121573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111459190612106565b60ff168110156112c857604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa1580156111b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111dd91906120c4565b905060005b818110156112be576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015611257573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061127b9190612162565b6000015185858151811061129157611291612123565b6001600160a01b0390921660209283029190910190910152836112b3816121d2565b9450506001016111e2565b50506001016110c3565b5090949350505050565b609854156113225760405162461bcd60e51b815260206004820152601c60248201527f5061757361626c653a20636f6e74726163742069732070617573656400000000604482015260640161037e565b600083815260cb602052604090205460ff16156113815760405162461bcd60e51b815260206004820152601960248201527f4d65737361676520616c72656164792076616c69646174656400000000000000604482015260640161037e565b60008260405160200161139491906122e8565b6040516020818303038152906040528051906020012090506000611405826040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905060006114138285611931565b90506001600160a01b038116331461146d5760405162461bcd60e51b815260206004820152601e60248201527f4d657373616765207369676e6572206973206e6f74206f70657261746f720000604482015260640161037e565b600086815260ca602052604090206114858682612351565b50600086815260cb60209081526040808320805460ff1916600117905560cc9091529081902054905187917f4357217b8431f5ad92bca34a3a90de7f3ce5e3eff6683ad82984c957e9c68b24916114e99189916001600160a01b0390911690612411565b60405180910390a2505050505050565b611501611885565b6001600160a01b0381166115665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161037e565b610390816118df565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115e69190611ff3565b6001600160a01b0316336001600160a01b0316146116165760405162461bcd60e51b815260040161037e90612010565b6098541981196098541916146116945760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161037e565b609881905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016104c7565b6001600160a01b0381166117595760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161037e565b609754604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1609780546001600160a01b0319166001600160a01b0392909216919091179055565b60606000806117d084611957565b61ffff1667ffffffffffffffff8111156117ec576117ec611d97565b6040519080825280601f01601f191660200182016040528015611816576020820181803683370190505b5090506000805b82518210801561182e575061010081105b156112c8576001811b935085841615611875578060f81b83838151811061185757611857612123565b60200101906001600160f81b031916908160001a9053508160010191505b61187e816121d2565b905061181d565b6033546001600160a01b03163314610b755760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161037e565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008060006119408585611982565b9150915061194d816119f0565b5090505b92915050565b6000805b82156119515761196c60018461243b565b909216918061197a8161244e565b91505061195b565b60008082516041036119b85760208301516040840151606085015160001a6119ac87828585611ba6565b945094505050506119e9565b82516040036119e157602083015160408401516119d6868383611c93565b9350935050506119e9565b506000905060025b9250929050565b6000816004811115611a0457611a0461246f565b03611a0c5750565b6001816004811115611a2057611a2061246f565b03611a6d5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161037e565b6002816004811115611a8157611a8161246f565b03611ace5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161037e565b6003816004811115611ae257611ae261246f565b03611b3a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161037e565b6004816004811115611b4e57611b4e61246f565b036103905760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161037e565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611bdd5750600090506003611c8a565b8460ff16601b14158015611bf557508460ff16601c14155b15611c065750600090506004611c8a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611c5a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611c8357600060019250925050611c8a565b9150600090505b94509492505050565b6000806001600160ff1b03831681611cb060ff86901c601b61214f565b9050611cbe87828885611ba6565b935093505050935093915050565b6001600160a01b038116811461039057600080fd5b600060208284031215611cf357600080fd5b8135611cfe81611ccc565b9392505050565b600060208284031215611d1757600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015611d5f5783516001600160a01b031683529284019291840191600101611d3a565b50909695505050505050565b60ff8116811461039057600080fd5b600060208284031215611d8c57600080fd5b8135611cfe81611d6b565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611dbe57600080fd5b813567ffffffffffffffff80821115611dd957611dd9611d97565b604051601f8301601f19908116603f01168101908282118183101715611e0157611e01611d97565b81604052838152866020858801011115611e1a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215611e4c57600080fd5b813567ffffffffffffffff811115611e6357600080fd5b611e6f84828501611dad565b949350505050565b60008060408385031215611e8a57600080fd5b8235611e9581611ccc565b9150602083013567ffffffffffffffff80821115611eb257600080fd5b9084019060608287031215611ec657600080fd5b604051606081018181108382111715611ee157611ee1611d97565b604052823582811115611ef357600080fd5b611eff88828601611dad565b82525060208301356020820152604083013560408201528093505050509250929050565b60005b83811015611f3e578181015183820152602001611f26565b50506000910152565b60008151808452611f5f816020860160208601611f23565b601f01601f19169290920160200192915050565b602081526000611cfe6020830184611f47565b600080600060608486031215611f9b57600080fd5b83359250602084013567ffffffffffffffff80821115611fba57600080fd5b611fc687838801611dad565b93506040860135915080821115611fdc57600080fd5b50611fe986828701611dad565b9150509250925092565b60006020828403121561200557600080fd5b8151611cfe81611ccc565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561206c57600080fd5b81518015158114611cfe57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000602082840312156120d657600080fd5b5051919050565b6000602082840312156120ef57600080fd5b81516001600160c01b0381168114611cfe57600080fd5b60006020828403121561211857600080fd5b8151611cfe81611d6b565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8082018082111561195157611951612139565b60006040828403121561217457600080fd5b6040516040810181811067ffffffffffffffff8211171561219757612197611d97565b60405282516121a581611ccc565b815260208301516bffffffffffffffffffffffff811681146121c657600080fd5b60208201529392505050565b6000600182016121e4576121e4612139565b5060010190565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261228d60a0840182611f47565b90506020840151606084015260408401516080840152809150509392505050565b600181811c908216806122c257607f821691505b6020821081036122e257634e487b7160e01b600052602260045260246000fd5b50919050565b600082516122fa818460208701611f23565b9190910192915050565b601f82111561234c576000816000526020600020601f850160051c8101602086101561232d5750805b601f850160051c820191505b81811015610c3b57828155600101612339565b505050565b815167ffffffffffffffff81111561236b5761236b611d97565b61237f8161237984546122ae565b84612304565b602080601f8311600181146123b4576000841561239c5750858301515b600019600386901b1c1916600185901b178555610c3b565b600085815260208120601f198616915b828110156123e3578886015182559484019460019091019084016123c4565b50858210156124015787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6040815260006124246040830185611f47565b905060018060a01b03831660208301529392505050565b8181038181111561195157611951612139565b600061ffff80831681810361246557612465612139565b6001019392505050565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220da7e7eb2daf965868f5a25d602a7d262fd6b853b75fc1d918bf3e6f81438a87664736f6c63430008190033",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// BindingsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingsMetaData.Bin instead.
var BindingsBin = BindingsMetaData.Bin

// DeployBindings deploys a new Ethereum contract, binding an instance of Bindings to it.
func DeployBindings(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address) (common.Address, *types.Transaction, *Bindings, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingsBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_Bindings *BindingsCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_Bindings *BindingsSession) AvsDirectory() (common.Address, error) {
	return _Bindings.Contract.AvsDirectory(&_Bindings.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_Bindings *BindingsCallerSession) AvsDirectory() (common.Address, error) {
	return _Bindings.Contract.AvsDirectory(&_Bindings.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Bindings *BindingsCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Bindings *BindingsSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _Bindings.Contract.GetOperatorRestakedStrategies(&_Bindings.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Bindings *BindingsCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _Bindings.Contract.GetOperatorRestakedStrategies(&_Bindings.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Bindings *BindingsCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Bindings *BindingsSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _Bindings.Contract.GetRestakeableStrategies(&_Bindings.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Bindings *BindingsCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _Bindings.Contract.GetRestakeableStrategies(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Bindings *BindingsCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Bindings *BindingsSession) Paused(index uint8) (bool, error) {
	return _Bindings.Contract.Paused(&_Bindings.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Bindings *BindingsCallerSession) Paused(index uint8) (bool, error) {
	return _Bindings.Contract.Paused(&_Bindings.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Bindings *BindingsCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Bindings *BindingsSession) Paused0() (*big.Int, error) {
	return _Bindings.Contract.Paused0(&_Bindings.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Bindings *BindingsCallerSession) Paused0() (*big.Int, error) {
	return _Bindings.Contract.Paused0(&_Bindings.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Bindings *BindingsCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Bindings *BindingsSession) PauserRegistry() (common.Address, error) {
	return _Bindings.Contract.PauserRegistry(&_Bindings.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Bindings *BindingsCallerSession) PauserRegistry() (common.Address, error) {
	return _Bindings.Contract.PauserRegistry(&_Bindings.CallOpts)
}

// RetrieveMessage is a free data retrieval call binding the contract method 0x9ba8fd09.
//
// Solidity: function retrieveMessage(uint256 _messageId) view returns(string)
func (_Bindings *BindingsCaller) RetrieveMessage(opts *bind.CallOpts, _messageId *big.Int) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "retrieveMessage", _messageId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RetrieveMessage is a free data retrieval call binding the contract method 0x9ba8fd09.
//
// Solidity: function retrieveMessage(uint256 _messageId) view returns(string)
func (_Bindings *BindingsSession) RetrieveMessage(_messageId *big.Int) (string, error) {
	return _Bindings.Contract.RetrieveMessage(&_Bindings.CallOpts, _messageId)
}

// RetrieveMessage is a free data retrieval call binding the contract method 0x9ba8fd09.
//
// Solidity: function retrieveMessage(uint256 _messageId) view returns(string)
func (_Bindings *BindingsCallerSession) RetrieveMessage(_messageId *big.Int) (string, error) {
	return _Bindings.Contract.RetrieveMessage(&_Bindings.CallOpts, _messageId)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Bindings *BindingsTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Bindings *BindingsSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.DeregisterOperatorFromAVS(&_Bindings.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Bindings *BindingsTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.DeregisterOperatorFromAVS(&_Bindings.TransactOpts, operator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Bindings *BindingsTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Bindings *BindingsSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Bindings *BindingsTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Bindings *BindingsTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Bindings *BindingsSession) PauseAll() (*types.Transaction, error) {
	return _Bindings.Contract.PauseAll(&_Bindings.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Bindings *BindingsTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Bindings.Contract.PauseAll(&_Bindings.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Bindings *BindingsTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Bindings *BindingsSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterOperatorToAVS(&_Bindings.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Bindings *BindingsTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterOperatorToAVS(&_Bindings.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Bindings *BindingsTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Bindings *BindingsSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPauserRegistry(&_Bindings.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Bindings *BindingsTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPauserRegistry(&_Bindings.TransactOpts, newPauserRegistry)
}

// StoreValidatedMessage is a paid mutator transaction binding the contract method 0xf2cb613a.
//
// Solidity: function storeValidatedMessage(uint256 _messageId, string _message, bytes _signature) returns()
func (_Bindings *BindingsTransactor) StoreValidatedMessage(opts *bind.TransactOpts, _messageId *big.Int, _message string, _signature []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "storeValidatedMessage", _messageId, _message, _signature)
}

// StoreValidatedMessage is a paid mutator transaction binding the contract method 0xf2cb613a.
//
// Solidity: function storeValidatedMessage(uint256 _messageId, string _message, bytes _signature) returns()
func (_Bindings *BindingsSession) StoreValidatedMessage(_messageId *big.Int, _message string, _signature []byte) (*types.Transaction, error) {
	return _Bindings.Contract.StoreValidatedMessage(&_Bindings.TransactOpts, _messageId, _message, _signature)
}

// StoreValidatedMessage is a paid mutator transaction binding the contract method 0xf2cb613a.
//
// Solidity: function storeValidatedMessage(uint256 _messageId, string _message, bytes _signature) returns()
func (_Bindings *BindingsTransactorSession) StoreValidatedMessage(_messageId *big.Int, _message string, _signature []byte) (*types.Transaction, error) {
	return _Bindings.Contract.StoreValidatedMessage(&_Bindings.TransactOpts, _messageId, _message, _signature)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string _message) returns(uint256)
func (_Bindings *BindingsTransactor) SubmitMessage(opts *bind.TransactOpts, _message string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitMessage", _message)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string _message) returns(uint256)
func (_Bindings *BindingsSession) SubmitMessage(_message string) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitMessage(&_Bindings.TransactOpts, _message)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string _message) returns(uint256)
func (_Bindings *BindingsTransactorSession) SubmitMessage(_message string) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitMessage(&_Bindings.TransactOpts, _message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Bindings *BindingsTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Bindings *BindingsSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Unpause(&_Bindings.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Bindings *BindingsTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Unpause(&_Bindings.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Bindings *BindingsTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Bindings *BindingsSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateAVSMetadataURI(&_Bindings.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Bindings *BindingsTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateAVSMetadataURI(&_Bindings.TransactOpts, _metadataURI)
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMessageSubmittedIterator is returned from FilterMessageSubmitted and is used to iterate over the raw logs and unpacked data for MessageSubmitted events raised by the Bindings contract.
type BindingsMessageSubmittedIterator struct {
	Event *BindingsMessageSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsMessageSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMessageSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsMessageSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsMessageSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMessageSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMessageSubmitted represents a MessageSubmitted event raised by the Bindings contract.
type BindingsMessageSubmitted struct {
	MessageId *big.Int
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageSubmitted is a free log retrieval operation binding the contract event 0x880e27ad36e62ea827e957f86d06f355728f6630b504a86ca8c371962f929ff0.
//
// Solidity: event MessageSubmitted(uint256 indexed messageId, string message)
func (_Bindings *BindingsFilterer) FilterMessageSubmitted(opts *bind.FilterOpts, messageId []*big.Int) (*BindingsMessageSubmittedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MessageSubmitted", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMessageSubmittedIterator{contract: _Bindings.contract, event: "MessageSubmitted", logs: logs, sub: sub}, nil
}

// WatchMessageSubmitted is a free log subscription operation binding the contract event 0x880e27ad36e62ea827e957f86d06f355728f6630b504a86ca8c371962f929ff0.
//
// Solidity: event MessageSubmitted(uint256 indexed messageId, string message)
func (_Bindings *BindingsFilterer) WatchMessageSubmitted(opts *bind.WatchOpts, sink chan<- *BindingsMessageSubmitted, messageId []*big.Int) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MessageSubmitted", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMessageSubmitted)
				if err := _Bindings.contract.UnpackLog(event, "MessageSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageSubmitted is a log parse operation binding the contract event 0x880e27ad36e62ea827e957f86d06f355728f6630b504a86ca8c371962f929ff0.
//
// Solidity: event MessageSubmitted(uint256 indexed messageId, string message)
func (_Bindings *BindingsFilterer) ParseMessageSubmitted(log types.Log) (*BindingsMessageSubmitted, error) {
	event := new(BindingsMessageSubmitted)
	if err := _Bindings.contract.UnpackLog(event, "MessageSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMessageValidatedIterator is returned from FilterMessageValidated and is used to iterate over the raw logs and unpacked data for MessageValidated events raised by the Bindings contract.
type BindingsMessageValidatedIterator struct {
	Event *BindingsMessageValidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsMessageValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMessageValidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsMessageValidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsMessageValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMessageValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMessageValidated represents a MessageValidated event raised by the Bindings contract.
type BindingsMessageValidated struct {
	MessageId *big.Int
	Message   string
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageValidated is a free log retrieval operation binding the contract event 0x4357217b8431f5ad92bca34a3a90de7f3ce5e3eff6683ad82984c957e9c68b24.
//
// Solidity: event MessageValidated(uint256 indexed messageId, string message, address sender)
func (_Bindings *BindingsFilterer) FilterMessageValidated(opts *bind.FilterOpts, messageId []*big.Int) (*BindingsMessageValidatedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MessageValidated", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMessageValidatedIterator{contract: _Bindings.contract, event: "MessageValidated", logs: logs, sub: sub}, nil
}

// WatchMessageValidated is a free log subscription operation binding the contract event 0x4357217b8431f5ad92bca34a3a90de7f3ce5e3eff6683ad82984c957e9c68b24.
//
// Solidity: event MessageValidated(uint256 indexed messageId, string message, address sender)
func (_Bindings *BindingsFilterer) WatchMessageValidated(opts *bind.WatchOpts, sink chan<- *BindingsMessageValidated, messageId []*big.Int) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MessageValidated", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMessageValidated)
				if err := _Bindings.contract.UnpackLog(event, "MessageValidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageValidated is a log parse operation binding the contract event 0x4357217b8431f5ad92bca34a3a90de7f3ce5e3eff6683ad82984c957e9c68b24.
//
// Solidity: event MessageValidated(uint256 indexed messageId, string message, address sender)
func (_Bindings *BindingsFilterer) ParseMessageValidated(log types.Log) (*BindingsMessageValidated, error) {
	event := new(BindingsMessageValidated)
	if err := _Bindings.contract.UnpackLog(event, "MessageValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bindings contract.
type BindingsPausedIterator struct {
	Event *BindingsPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPaused represents a Paused event raised by the Bindings contract.
type BindingsPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Bindings *BindingsFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*BindingsPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPausedIterator{contract: _Bindings.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Bindings *BindingsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BindingsPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPaused)
				if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Bindings *BindingsFilterer) ParsePaused(log types.Log) (*BindingsPaused, error) {
	event := new(BindingsPaused)
	if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the Bindings contract.
type BindingsPauserRegistrySetIterator struct {
	Event *BindingsPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPauserRegistrySet represents a PauserRegistrySet event raised by the Bindings contract.
type BindingsPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Bindings *BindingsFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*BindingsPauserRegistrySetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &BindingsPauserRegistrySetIterator{contract: _Bindings.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Bindings *BindingsFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *BindingsPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPauserRegistrySet)
				if err := _Bindings.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Bindings *BindingsFilterer) ParsePauserRegistrySet(log types.Log) (*BindingsPauserRegistrySet, error) {
	event := new(BindingsPauserRegistrySet)
	if err := _Bindings.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bindings contract.
type BindingsUnpausedIterator struct {
	Event *BindingsUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUnpaused represents a Unpaused event raised by the Bindings contract.
type BindingsUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Bindings *BindingsFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*BindingsUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindingsUnpausedIterator{contract: _Bindings.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Bindings *BindingsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BindingsUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUnpaused)
				if err := _Bindings.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Bindings *BindingsFilterer) ParseUnpaused(log types.Log) (*BindingsUnpaused, error) {
	event := new(BindingsUnpaused)
	if err := _Bindings.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
