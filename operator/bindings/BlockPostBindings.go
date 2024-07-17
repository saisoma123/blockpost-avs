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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageSignatures\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageValidated\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messages\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retrieveMessage\",\"inputs\":[{\"name\":\"_messageId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeValidatedMessage\",\"inputs\":[{\"name\":\"_messageId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitMessage\",\"inputs\":[{\"name\":\"_message\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageStored\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSubmitted\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageValidated\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561001057600080fd5b506040516127a13803806127a183398101604081905261002f9161015b565b6001600160a01b0380841660c052808316608052811660a052828282610053610083565b505060ca80546001600160a01b0319166001600160a01b0393909316929092179091555050600160c955506101a8565b600054610100900460ff16156100ef5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015610141576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461015857600080fd5b50565b60008060006060848603121561017057600080fd5b835161017b81610143565b602085015190935061018c81610143565b604085015190925061019d81610143565b809150509250925092565b60805160a05160c0516125686102396000396000818161021a01528181610ccd01528181610f020152610f810152600081816107d001528181610922015281816109b901528181611081015281816111fb015261129a0152600081816105fb0152818161068a0152818161070a01528181610c7901528181610ea601528181610fbc015261115601526125686000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063886f1195116100b8578063a98fb3551161007c578063a98fb355146102ca578063e481af9d146102dd578063f257c822146102e5578063f2cb613a14610308578063f2fde38b1461031b578063fabc1cbc1461032e57600080fd5b8063886f11951461026d5780638da5cb5b146102805780639926ee7d146102915780639ba8fd09146102a4578063a364f4da146102b757600080fd5b8063595c6a671161010a578063595c6a67146101cb5780635ac86ab7146101d35780635c975abb146102065780636b3aa72e14610218578063708b34fe14610252578063715018a61461026557600080fd5b80630d80fefd1461014757806310d67a2f14610170578063136439dd1461018557806333cfb7b7146101985780633a6dc0cf146101b8575b600080fd5b61015a610155366004611da3565b610341565b6040516101679190611e0c565b60405180910390f35b61018361017e366004611e3b565b6103db565b005b610183610193366004611da3565b610497565b6101ab6101a6366004611e3b565b6105d6565b6040516101679190611e58565b61015a6101c6366004611da3565b610a86565b610183610a9f565b6101f66101e1366004611eb4565b609854600160ff9092169190911b9081161490565b6040519015158152602001610167565b6098545b604051908152602001610167565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610167565b61020a610260366004611f74565b610b66565b610183610c5a565b60975461023a906001600160a01b031681565b6033546001600160a01b031661023a565b61018361029f366004611fb1565b610c6e565b61015a6102b2366004611da3565b610d3a565b6101836102c5366004611e3b565b610e9b565b6101836102d8366004611f74565b610f62565b6101ab610fb6565b6101f66102f3366004611da3565b60cc6020526000908152604090205460ff1681565b61018361031636600461205d565b611363565b610183610329366004611e3b565b6115d0565b61018361033c366004611da3565b611646565b60cb602052600090815260409020805461035a906120ca565b80601f0160208091040260200160405190810160405280929190818152602001828054610386906120ca565b80156103d35780601f106103a8576101008083540402835291602001916103d3565b820191906000526020600020905b8154815290600101906020018083116103b657829003601f168201915b505050505081565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561042e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104529190612104565b6001600160a01b0316336001600160a01b03161461048b5760405162461bcd60e51b815260040161048290612121565b60405180910390fd5b610494816117a2565b50565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156104df573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610503919061216b565b61051f5760405162461bcd60e51b81526004016104829061218d565b609854818116146105985760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610482565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610642573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066691906121d5565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156106d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f591906121ee565b90506001600160c01b038116158061078f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610766573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078a9190612217565b60ff16155b156107ab57505060408051600081526020810190915292915050565b60006107bf826001600160c01b0316611899565b90506000805b825181101561088b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061080f5761080f612234565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610853573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087791906121d5565b6108819083612260565b91506001016107c5565b5060008167ffffffffffffffff8111156108a7576108a7611ed1565b6040519080825280602002602001820160405280156108d0578160200160208202803683370190505b5090506000805b8451811015610a795760008582815181106108f4576108f4612234565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610969573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098d91906121d5565b905060005b81811015610a6e576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610a07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2b9190612273565b60000151868681518110610a4157610a41612234565b6001600160a01b039092166020928302919091019091015284610a63816122e3565b955050600101610992565b5050506001016108d7565b5090979650505050505050565b60cd602052600090815260409020805461035a906120ca565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610ae7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0b919061216b565b610b275760405162461bcd60e51b81526004016104829061218d565b600019609881905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000609854600014610bba5760405162461bcd60e51b815260206004820152601c60248201527f5061757361626c653a20636f6e747261637420697320706175736564000000006044820152606401610482565b6000825111610c055760405162461bcd60e51b81526020600482015260176024820152764d6573736167652063616e6e6f7420626520656d70747960481b6044820152606401610482565b60c980549081906000610c17836122e3565b9190505550807f880e27ad36e62ea827e957f86d06f355728f6630b504a86ca8c371962f929ff084604051610c4c9190611e0c565b60405180910390a292915050565b610c6261195c565b610c6c60006119b6565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cb65760405162461bcd60e51b8152600401610482906122fc565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610d049085908590600401612374565b600060405180830381600087803b158015610d1e57600080fd5b505af1158015610d32573d6000803e3d6000fd5b505050505050565b600081815260cb6020526040812080546060929190610d58906120ca565b905011610da75760405162461bcd60e51b815260206004820152601960248201527f4d65737361676520494420646f6573206e6f74206578697374000000000000006044820152606401610482565b600082815260cc602052604090205460ff16610dfd5760405162461bcd60e51b815260206004820152601560248201527413595cdcd859d9481b9bdd081d985b1a59185d1959605a1b6044820152606401610482565b600082815260cb602052604090208054610e16906120ca565b80601f0160208091040260200160405190810160405280929190818152602001828054610e42906120ca565b8015610e8f5780601f10610e6457610100808354040283529160200191610e8f565b820191906000526020600020905b815481529060010190602001808311610e7257829003601f168201915b50505050509050919050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ee35760405162461bcd60e51b8152600401610482906122fc565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610f4757600080fd5b505af1158015610f5b573d6000803e3d6000fd5b5050505050565b610f6a61195c565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610f2d908490600401611e0c565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611018573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103c9190612217565b60ff1690508060000361105d57505060408051600081526020810190915290565b6000805b8281101561110857604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa1580156110d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f491906121d5565b6110fe9083612260565b9150600101611061565b5060008167ffffffffffffffff81111561112457611124611ed1565b60405190808252806020026020018201604052801561114d578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111d69190612217565b60ff1681101561135957604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa15801561124a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126e91906121d5565b905060005b8181101561134f576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156112e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130c9190612273565b6000015185858151811061132257611322612234565b6001600160a01b039092166020928302919091019091015283611344816122e3565b945050600101611273565b5050600101611154565b5090949350505050565b609854156113b35760405162461bcd60e51b815260206004820152601c60248201527f5061757361626c653a20636f6e747261637420697320706175736564000000006044820152606401610482565b60008251116113fe5760405162461bcd60e51b81526020600482015260176024820152764d6573736167652063616e6e6f7420626520656d70747960481b6044820152606401610482565b600083815260cc602052604090205460ff161561145d5760405162461bcd60e51b815260206004820152601960248201527f4d65737361676520616c72656164792076616c696461746564000000000000006044820152606401610482565b60008260405160200161147091906123bf565b60405160208183030381529060405280519060200120905060006114e1826040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905060006114ef8285611a08565b90506001600160a01b03811633146115495760405162461bcd60e51b815260206004820152601e60248201527f4d657373616765207369676e6572206973206e6f74206f70657261746f7200006044820152606401610482565b600086815260cb602052604090206115618682612428565b50600086815260cc60209081526040808320805460ff1916600117905560cd909152902061158f8582612428565b50857f94bb9cbeac3163d2f79f8b4d3ebd4287f647cf6d464b4beba6fbfdeb8c259058866040516115c09190611e0c565b60405180910390a2505050505050565b6115d861195c565b6001600160a01b03811661163d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610482565b610494816119b6565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611699573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116bd9190612104565b6001600160a01b0316336001600160a01b0316146116ed5760405162461bcd60e51b815260040161048290612121565b60985419811960985419161461176b5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610482565b609881905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016105cb565b6001600160a01b0381166118305760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610482565b609754604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1609780546001600160a01b0319166001600160a01b0392909216919091179055565b60606000806118a784611a2e565b61ffff1667ffffffffffffffff8111156118c3576118c3611ed1565b6040519080825280601f01601f1916602001820160405280156118ed576020820181803683370190505b5090506000805b825182108015611905575061010081105b15611359576001811b93508584161561194c578060f81b83838151811061192e5761192e612234565b60200101906001600160f81b031916908160001a9053508160010191505b611955816122e3565b90506118f4565b6033546001600160a01b03163314610c6c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610482565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000806000611a178585611a59565b91509150611a2481611ac7565b5090505b92915050565b6000805b8215611a2857611a436001846124e8565b9092169180611a51816124fb565b915050611a32565b6000808251604103611a8f5760208301516040840151606085015160001a611a8387828585611c7d565b94509450505050611ac0565b8251604003611ab85760208301516040840151611aad868383611d6a565b935093505050611ac0565b506000905060025b9250929050565b6000816004811115611adb57611adb61251c565b03611ae35750565b6001816004811115611af757611af761251c565b03611b445760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610482565b6002816004811115611b5857611b5861251c565b03611ba55760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610482565b6003816004811115611bb957611bb961251c565b03611c115760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610482565b6004816004811115611c2557611c2561251c565b036104945760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610482565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611cb45750600090506003611d61565b8460ff16601b14158015611ccc57508460ff16601c14155b15611cdd5750600090506004611d61565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611d31573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611d5a57600060019250925050611d61565b9150600090505b94509492505050565b6000806001600160ff1b03831681611d8760ff86901c601b612260565b9050611d9587828885611c7d565b935093505050935093915050565b600060208284031215611db557600080fd5b5035919050565b60005b83811015611dd7578181015183820152602001611dbf565b50506000910152565b60008151808452611df8816020860160208601611dbc565b601f01601f19169290920160200192915050565b602081526000611e1f6020830184611de0565b9392505050565b6001600160a01b038116811461049457600080fd5b600060208284031215611e4d57600080fd5b8135611e1f81611e26565b6020808252825182820181905260009190848201906040850190845b81811015611e995783516001600160a01b031683529284019291840191600101611e74565b50909695505050505050565b60ff8116811461049457600080fd5b600060208284031215611ec657600080fd5b8135611e1f81611ea5565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611ef857600080fd5b813567ffffffffffffffff80821115611f1357611f13611ed1565b604051601f8301601f19908116603f01168101908282118183101715611f3b57611f3b611ed1565b81604052838152866020858801011115611f5457600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215611f8657600080fd5b813567ffffffffffffffff811115611f9d57600080fd5b611fa984828501611ee7565b949350505050565b60008060408385031215611fc457600080fd5b8235611fcf81611e26565b9150602083013567ffffffffffffffff80821115611fec57600080fd5b908401906060828703121561200057600080fd5b60405160608101818110838211171561201b5761201b611ed1565b60405282358281111561202d57600080fd5b61203988828601611ee7565b82525060208301356020820152604083013560408201528093505050509250929050565b60008060006060848603121561207257600080fd5b83359250602084013567ffffffffffffffff8082111561209157600080fd5b61209d87838801611ee7565b935060408601359150808211156120b357600080fd5b506120c086828701611ee7565b9150509250925092565b600181811c908216806120de57607f821691505b6020821081036120fe57634e487b7160e01b600052602260045260246000fd5b50919050565b60006020828403121561211657600080fd5b8151611e1f81611e26565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561217d57600080fd5b81518015158114611e1f57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000602082840312156121e757600080fd5b5051919050565b60006020828403121561220057600080fd5b81516001600160c01b0381168114611e1f57600080fd5b60006020828403121561222957600080fd5b8151611e1f81611ea5565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115611a2857611a2861224a565b60006040828403121561228557600080fd5b6040516040810181811067ffffffffffffffff821117156122a8576122a8611ed1565b60405282516122b681611e26565b815260208301516bffffffffffffffffffffffff811681146122d757600080fd5b60208201529392505050565b6000600182016122f5576122f561224a565b5060010190565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261239e60a0840182611de0565b90506020840151606084015260408401516080840152809150509392505050565b600082516123d1818460208701611dbc565b9190910192915050565b601f821115612423576000816000526020600020601f850160051c810160208610156124045750805b601f850160051c820191505b81811015610d3257828155600101612410565b505050565b815167ffffffffffffffff81111561244257612442611ed1565b6124568161245084546120ca565b846123db565b602080601f83116001811461248b57600084156124735750858301515b600019600386901b1c1916600185901b178555610d32565b600085815260208120601f198616915b828110156124ba5788860151825594840194600190910190840161249b565b50858210156124d85787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b81810381811115611a2857611a2861224a565b600061ffff8083168181036125125761251261224a565b6001019392505050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212209c0666ebb2264c967a70605248bf7f423d75fa955862952058c16f55235ca4f064736f6c63430008190033",
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

// MessageSignatures is a free data retrieval call binding the contract method 0x3a6dc0cf.
//
// Solidity: function messageSignatures(uint256 ) view returns(bytes)
func (_Bindings *BindingsCaller) MessageSignatures(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "messageSignatures", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MessageSignatures is a free data retrieval call binding the contract method 0x3a6dc0cf.
//
// Solidity: function messageSignatures(uint256 ) view returns(bytes)
func (_Bindings *BindingsSession) MessageSignatures(arg0 *big.Int) ([]byte, error) {
	return _Bindings.Contract.MessageSignatures(&_Bindings.CallOpts, arg0)
}

// MessageSignatures is a free data retrieval call binding the contract method 0x3a6dc0cf.
//
// Solidity: function messageSignatures(uint256 ) view returns(bytes)
func (_Bindings *BindingsCallerSession) MessageSignatures(arg0 *big.Int) ([]byte, error) {
	return _Bindings.Contract.MessageSignatures(&_Bindings.CallOpts, arg0)
}

// MessageValidated is a free data retrieval call binding the contract method 0xf257c822.
//
// Solidity: function messageValidated(uint256 ) view returns(bool)
func (_Bindings *BindingsCaller) MessageValidated(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "messageValidated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageValidated is a free data retrieval call binding the contract method 0xf257c822.
//
// Solidity: function messageValidated(uint256 ) view returns(bool)
func (_Bindings *BindingsSession) MessageValidated(arg0 *big.Int) (bool, error) {
	return _Bindings.Contract.MessageValidated(&_Bindings.CallOpts, arg0)
}

// MessageValidated is a free data retrieval call binding the contract method 0xf257c822.
//
// Solidity: function messageValidated(uint256 ) view returns(bool)
func (_Bindings *BindingsCallerSession) MessageValidated(arg0 *big.Int) (bool, error) {
	return _Bindings.Contract.MessageValidated(&_Bindings.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string)
func (_Bindings *BindingsCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "messages", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string)
func (_Bindings *BindingsSession) Messages(arg0 *big.Int) (string, error) {
	return _Bindings.Contract.Messages(&_Bindings.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string)
func (_Bindings *BindingsCallerSession) Messages(arg0 *big.Int) (string, error) {
	return _Bindings.Contract.Messages(&_Bindings.CallOpts, arg0)
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

// BindingsMessageStoredIterator is returned from FilterMessageStored and is used to iterate over the raw logs and unpacked data for MessageStored events raised by the Bindings contract.
type BindingsMessageStoredIterator struct {
	Event *BindingsMessageStored // Event containing the contract specifics and raw log

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
func (it *BindingsMessageStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMessageStored)
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
		it.Event = new(BindingsMessageStored)
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
func (it *BindingsMessageStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMessageStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMessageStored represents a MessageStored event raised by the Bindings contract.
type BindingsMessageStored struct {
	MessageId *big.Int
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageStored is a free log retrieval operation binding the contract event 0x6a1c86538e4684602227977281b3fd4d12c6cc23e4db736ea995984616652f7c.
//
// Solidity: event MessageStored(uint256 indexed messageId, string message)
func (_Bindings *BindingsFilterer) FilterMessageStored(opts *bind.FilterOpts, messageId []*big.Int) (*BindingsMessageStoredIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MessageStored", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMessageStoredIterator{contract: _Bindings.contract, event: "MessageStored", logs: logs, sub: sub}, nil
}

// WatchMessageStored is a free log subscription operation binding the contract event 0x6a1c86538e4684602227977281b3fd4d12c6cc23e4db736ea995984616652f7c.
//
// Solidity: event MessageStored(uint256 indexed messageId, string message)
func (_Bindings *BindingsFilterer) WatchMessageStored(opts *bind.WatchOpts, sink chan<- *BindingsMessageStored, messageId []*big.Int) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MessageStored", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMessageStored)
				if err := _Bindings.contract.UnpackLog(event, "MessageStored", log); err != nil {
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

// ParseMessageStored is a log parse operation binding the contract event 0x6a1c86538e4684602227977281b3fd4d12c6cc23e4db736ea995984616652f7c.
//
// Solidity: event MessageStored(uint256 indexed messageId, string message)
func (_Bindings *BindingsFilterer) ParseMessageStored(log types.Log) (*BindingsMessageStored, error) {
	event := new(BindingsMessageStored)
	if err := _Bindings.contract.UnpackLog(event, "MessageStored", log); err != nil {
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
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageValidated is a free log retrieval operation binding the contract event 0x94bb9cbeac3163d2f79f8b4d3ebd4287f647cf6d464b4beba6fbfdeb8c259058.
//
// Solidity: event MessageValidated(uint256 indexed messageId, string message)
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

// WatchMessageValidated is a free log subscription operation binding the contract event 0x94bb9cbeac3163d2f79f8b4d3ebd4287f647cf6d464b4beba6fbfdeb8c259058.
//
// Solidity: event MessageValidated(uint256 indexed messageId, string message)
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

// ParseMessageValidated is a log parse operation binding the contract event 0x94bb9cbeac3163d2f79f8b4d3ebd4287f647cf6d464b4beba6fbfdeb8c259058.
//
// Solidity: event MessageValidated(uint256 indexed messageId, string message)
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
