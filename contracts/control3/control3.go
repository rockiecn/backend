// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package control3

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FsOut is an auto generated low-level Go binding around an user-defined struct.
type FsOut struct {
	Nonce    uint64
	SubNonce uint64
}

// GroupOut is an auto generated low-level Go binding around an user-defined struct.
type GroupOut struct {
	Size   uint64
	Sprice *big.Int
	Lost   *big.Int
}

// OrderIn is an auto generated low-level Go binding around an user-defined struct.
type OrderIn struct {
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint8
	Sprice *big.Int
}

// PWIn is an auto generated low-level Go binding around an user-defined struct.
type PWIn struct {
	PIndex uint64
	TIndex uint8
	Pay    *big.Int
	Lost   *big.Int
}

// RoleOut is an auto generated low-level Go binding around an user-defined struct.
type RoleOut struct {
	State     uint8
	RType     uint8
	Index     uint64
	GIndex    uint64
	Owner     common.Address
	Next      common.Address
	VerifyKey []byte
	Desc      []byte
}

// SettleOut is an auto generated low-level Go binding around an user-defined struct.
type SettleOut struct {
	Time    uint64
	Size    uint64
	Price   *big.Int
	MaxPay  *big.Int
	HasPaid *big.Int
	CanPay  *big.Int
	Lost    *big.Int
}

// StoreOut is an auto generated low-level Go binding around an user-defined struct.
type StoreOut struct {
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
}

// Control3ABI is the input ABI used to generate the binding from.
const Control3ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inst\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualMoney\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inst\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"quitRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pe\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPlePerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Control3FuncSigs maps the 4-byte function signature to its string representation.
var Control3FuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"de9375f2": "auth()",
	"bd6b2222": "inst()",
	"022914a7": "owners(address)",
	"d6985699": "quitRole(address,uint64)",
	"713d3b95": "setLevel(uint64,uint8,bytes[5])",
	"bd5d54b7": "setPeriod(uint64,bytes[5])",
	"130b89b1": "setPlePerB(uint64,uint256,uint256,bytes[5])",
	"54fd4d50": "version()",
	"9b4c757a": "withdraw(address,uint64,uint8,uint256)",
}

// Control3Bin is the compiled bytecode used for deploying new contracts.
var Control3Bin = "0x60806040526001805461ffff60a01b1916600160a11b1790553480156200002557600080fd5b50604051620020cd380380620020cd833981016040819052620000489162000097565b600180546001600160a01b039384166001600160a01b03199182161790915560028054929093169116179055620000cf565b80516001600160a01b03811681146200009257600080fd5b919050565b60008060408385031215620000ab57600080fd5b620000b6836200007a565b9150620000c6602084016200007a565b90509250929050565b611fee80620000df6000396000f3fe6080604052600436106100955760003560e01c80639b4c757a116100595780639b4c757a1461017d578063bd5d54b71461019d578063bd6b2222146101bd578063d6985699146101f5578063de9375f21461021557600080fd5b8063022914a7146100a1578063130b89b1146100e65780634bf1b1341461010857806354fd4d5014610128578063713d3b951461015d57600080fd5b3661009c57005b600080fd5b3480156100ad57600080fd5b506100d16100bc366004611961565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156100f257600080fd5b50610106610101366004611ad6565b610235565b005b34801561011457600080fd5b50610106610123366004611b46565b6103ed565b34801561013457600080fd5b5060015461014a90600160a01b900461ffff1681565b60405161ffff90911681526020016100dd565b34801561016957600080fd5b50610106610178366004611bb6565b610566565b34801561018957600080fd5b50610106610198366004611be6565b61075a565b3480156101a957600080fd5b506101066101b8366004611c37565b61125a565b3480156101c957600080fd5b506002546101dd906001600160a01b031681565b6040516001600160a01b0390911681526020016100dd565b34801561020157600080fd5b50610106610210366004611c86565b6113f0565b34801561022157600080fd5b506001546101dd906001600160a01b031681565b6040516001600160601b03193060601b1660208201526c39b2ba283632b233b2a832b92160991b60348201526001600160c01b031960c086901b166041820152604981018490526069810183905260009060890160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906102d29084908690600401611cbf565b600060405180830381600087803b1580156102ec57600080fd5b505af1158015610300573d6000803e3d6000fd5b5050600254604051633ec7d5b960e01b8152600b60048201526001600160a01b039091169250829150633ec7d5b990602401602060405180830381865afa15801561034f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103739190611d46565b60405163cfb9854f60e01b81526001600160401b038816600482015260248101879052604481018690526001600160a01b03919091169063cfb9854f90606401600060405180830381600087803b1580156103cd57600080fd5b505af11580156103e1573d6000803e3d6000fd5b50505050505050505050565b823b600081900361043a5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064015b60405180910390fd5b6040516001600160601b031930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906104c29084908790600401611cbf565b600060405180830381600087803b1580156104dc57600080fd5b505af11580156104f0573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b60008260ff16116105a75760405162461bcd60e51b815260206004820152600b60248201526a1ddc9bdb99c81b195d995b60aa1b6044820152606401610431565b6040516001600160601b03193060601b166020820152671cd95d13195d995b60c21b60348201526001600160c01b031960c085901b16603c8201526001600160f81b031960f884901b16604482015260009060450160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906106459084908690600401611cbf565b600060405180830381600087803b15801561065f57600080fd5b505af1158015610673573d6000803e3d6000fd5b5050600254604051633ec7d5b960e01b8152600b60048201526001600160a01b039091169250829150633ec7d5b990602401602060405180830381865afa1580156106c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e69190611d46565b60405163936945bd60e01b81526001600160401b038716600482015260ff861660248201526001600160a01b03919091169063936945bd90604401600060405180830381600087803b15801561073b57600080fd5b505af115801561074f573d6000803e3d6000fd5b505050505050505050565b3360009081526020819052604090205460ff166107a55760405162461bcd60e51b81526020600482015260096024820152683737ba1037bbb732b960b91b6044820152606401610431565b600254604051633ec7d5b960e01b8152600760048201526001600160a01b03909116906000908290633ec7d5b990602401602060405180830381865afa1580156107f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108179190611d46565b6040516340d5dc7f60e11b815260ff861660048201526001600160a01b0391909116906381abb8fe906024016040805180830381865afa15801561085f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108839190611d63565b50604051633ec7d5b960e01b8152600660048201529091506000908190819081906001600160a01b03871690633ec7d5b990602401602060405180830381865afa1580156108d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f99190611d46565b604051637738515f60e01b81526001600160401b038b166004820152600060248201526001600160a01b039190911690637738515f9060440160a060405180830381865afa15801561094f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109739190611d92565b945094509450945050836001600160a01b03168a6001600160a01b0316146109ce5760405162461bcd60e51b815260206004820152600e60248201526d34b63632b3b0b61031b0b63632b960911b6044820152606401610431565b604051633ec7d5b960e01b8152600b60048201526000906001600160a01b03881690633ec7d5b990602401602060405180830381865afa158015610a16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3a9190611d46565b90506001600160401b03841615610afd57604051633d81c0dd60e21b81526001600160401b03851660048201526001600160a01b0382169063f607037490602401602060405180830381865afa158015610a98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abc9190611e07565b60ff16600103610afd5760405162461bcd60e51b815260206004820152600c60248201526b19dc9bdd5c0818985b9b995960a21b6044820152606401610431565b60008260ff16600303610da8576040516360b5903d60e01b81526001600160401b038616600482015260009081906001600160a01b038516906360b5903d906024016040805180830381865afa158015610b5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7f9190611e24565b915091508560ff16600203610c9c57604051633ec7d5b960e01b8152600a600482015281906001600160a01b038c1690633ec7d5b990602401602060405180830381865afa158015610bd5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bf99190611d46565b6001600160a01b03166364c64a4d8f8f6040518363ffffffff1660e01b8152600401610c3d9291906001600160401b0392909216825260ff16602082015260400190565b60e060405180830381865afa158015610c5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7e9190611e48565b602001516001600160401b0316610c959190611ef3565b9250610da5565b8560ff16600303610da557604051633ec7d5b960e01b8152600a600482015282906001600160a01b038c1690633ec7d5b990602401602060405180830381865afa158015610cee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d129190611d46565b604051635475cf8d60e01b81526001600160401b038a16600482015260ff8f1660248201526001600160a01b039190911690635475cf8d90604401606060405180830381865afa158015610d6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8e9190611f12565b51610da291906001600160401b0316611ef3565b92505b50505b604051633ec7d5b960e01b8152600860048201526000906001600160a01b038a1690633ec7d5b990602401602060405180830381865afa158015610df0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e149190611d46565b60405163fc3ba0ad60e01b81526001600160401b038e16600482015260ff8d1660248201526001600160a01b03919091169063fc3ba0ad90604401602060405180830381865afa158015610e6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e909190611f73565b9050808211610ea25760009150610eaf565b610eac8183611f8c565b91505b604051633ec7d5b960e01b8152600a60048201526000906001600160a01b038b1690633ec7d5b990602401602060405180830381865afa158015610ef7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1b9190611d46565b604051635affa0f360e01b81526001600160401b038f16600482015260ff8e166024820152604481018d9052606481018590526001600160a01b039190911690635affa0f3906084016020604051808303816000875af1158015610f83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa79190611f73565b90508a1580610fb55750808b115b8015610fca57506000876001600160401b0316115b8015610fd957508560ff166003145b1561111057808b1115610ff757610ff0818c611f8c565b9250610ffc565b600092505b604051633ec7d5b960e01b8152600d60048201526001600160a01b038b1690633ec7d5b990602401602060405180830381865afa158015611041573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110659190611d46565b6001600160a01b0316639cad43788e8e8a876040518563ffffffff1660e01b81526004016110be94939291906001600160401b03948516815260ff93909316602084015292166040820152606081019190915260800190565b6020604051808303816000875af11580156110dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111019190611f73565b925061110d8382611fa5565b90505b604051633ec7d5b960e01b8152600c60048201526001600160a01b038b1690633ec7d5b990602401602060405180830381865afa158015611155573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111799190611d46565b60405163530345bb60e01b81526001600160a01b038b811660048301528a8116602483015260448201849052919091169063530345bb90606401600060405180830381600087803b1580156111cd57600080fd5b505af11580156111e1573d6000803e3d6000fd5b505050508c6001600160401b03168e6001600160a01b03167f8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c08e8e856040516112429392919060ff9390931683526020830191909152604082015260600190565b60405180910390a35050505050505050505050505050565b6040516001600160601b03193060601b166020820152681cd95d14195c9a5bd960ba1b60348201526001600160c01b031960c084901b16603d82015260009060450160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906112e59084908690600401611cbf565b600060405180830381600087803b1580156112ff57600080fd5b505af1158015611313573d6000803e3d6000fd5b5050600254604051633ec7d5b960e01b8152600d60048201526001600160a01b039091169250829150633ec7d5b990602401602060405180830381865afa158015611362573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113869190611d46565b60405163184e5af560e11b81526001600160401b03861660048201526001600160a01b03919091169063309cb5ea90602401600060405180830381600087803b1580156113d257600080fd5b505af11580156113e6573d6000803e3d6000fd5b5050505050505050565b3360009081526020819052604090205460ff1661143b5760405162461bcd60e51b81526020600482015260096024820152683737ba1037bbb732b960b91b6044820152606401610431565b600254604051633ec7d5b960e01b8152600660048201526001600160a01b03909116906000908190819081908590633ec7d5b990602401602060405180830381865afa15801561148f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b39190611d46565b604051637738515f60e01b81526001600160401b0388166004820152600060248201526001600160a01b039190911690637738515f9060440160a060405180830381865afa158015611509573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061152d9190611d92565b945094509450945050836001600160a01b0316876001600160a01b0316146115845760405162461bcd60e51b815260206004820152600a6024820152696e65656420706179656560b01b6044820152606401610431565b8060ff166003146115c75760405162461bcd60e51b815260206004820152600d60248201526c73686f756c642061637469766560981b6044820152606401610431565b604051633ec7d5b960e01b8152600760048201526000906001600160a01b03871690633ec7d5b990602401602060405180830381865afa15801561160f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116339190611d46565b6001600160a01b0316637600b86a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611670573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116949190611e07565b604051633ec7d5b960e01b8152600b600482015290915060009081906001600160a01b03891690633ec7d5b990602401602060405180830381865afa1580156116e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117059190611d46565b6040516360b5903d60e01b81526001600160401b03881660048201526001600160a01b0391909116906360b5903d906024016040805180830381865afa158015611753573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117779190611e24565b604051633ec7d5b960e01b8152600a600482015291935091506001600160a01b03891690633ec7d5b990602401602060405180830381865afa1580156117c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e59190611d46565b604051630489342d60e31b81526001600160401b0380891660048301528b16602482015260ff8086166044830152871660648201526084810184905260a481018390526001600160a01b039190911690632449a1689060c401600060405180830381600087803b15801561185857600080fd5b505af115801561186c573d6000803e3d6000fd5b5050604051633ec7d5b960e01b8152600660048201526001600160a01b038b169250633ec7d5b99150602401602060405180830381865afa1580156118b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118d99190611d46565b6040516356d6b69d60e11b81526001600160401b038b1660048201526001600160a01b03919091169063adad6d3a90602401600060405180830381600087803b15801561192557600080fd5b505af1158015611939573d6000803e3d6000fd5b5050505050505050505050505050565b6001600160a01b038116811461195e57600080fd5b50565b60006020828403121561197357600080fd5b813561197e81611949565b9392505050565b6001600160401b038116811461195e57600080fd5b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156119d2576119d261199a565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611a0057611a0061199a565b604052919050565b6000601f8381840112611a1a57600080fd5b611a226119b0565b8060a0850186811115611a3457600080fd5b855b81811015611aca5780356001600160401b0380821115611a565760008081fd5b81890191508987830112611a6a5760008081fd5b8135602082821115611a7e57611a7e61199a565b611a8f828a01601f191682016119d8565b92508183528b81838601011115611aa65760008081fd5b81818501828501376000918301810191909152908652949094019350602001611a36565b50909695505050505050565b60008060008060808587031215611aec57600080fd5b8435611af781611985565b9350602085013592506040850135915060608501356001600160401b03811115611b2057600080fd5b611b2c87828801611a08565b91505092959194509250565b801515811461195e57600080fd5b600080600060608486031215611b5b57600080fd5b8335611b6681611949565b92506020840135611b7681611b38565b915060408401356001600160401b03811115611b9157600080fd5b611b9d86828701611a08565b9150509250925092565b60ff8116811461195e57600080fd5b600080600060608486031215611bcb57600080fd5b8335611bd681611985565b92506020840135611b7681611ba7565b60008060008060808587031215611bfc57600080fd5b8435611c0781611949565b93506020850135611c1781611985565b92506040850135611c2781611ba7565b9396929550929360600135925050565b60008060408385031215611c4a57600080fd5b8235611c5581611985565b915060208301356001600160401b03811115611c7057600080fd5b611c7c85828601611a08565b9150509250929050565b60008060408385031215611c9957600080fd5b8235611ca481611949565b91506020830135611cb481611985565b809150509250929050565b8281526040602080830182905260009160e08401919084018584805b6005811015611d3857878603603f1901845282518051808852835b81811015611d11578281018801518982018901528701611cf6565b508781018701849052601f01601f1916909601850195509284019291840191600101611cdb565b509398975050505050505050565b600060208284031215611d5857600080fd5b815161197e81611949565b60008060408385031215611d7657600080fd5b8251611d8181611949565b6020840151909250611cb481611b38565b600080600080600060a08688031215611daa57600080fd5b8551611db581611949565b6020870151909550611dc681611949565b6040870151909450611dd781611985565b6060870151909350611de881611ba7565b6080870151909250611df981611ba7565b809150509295509295909350565b600060208284031215611e1957600080fd5b815161197e81611ba7565b60008060408385031215611e3757600080fd5b505080516020909101519092909150565b600060e08284031215611e5a57600080fd5b60405160e081018181106001600160401b0382111715611e7c57611e7c61199a565b6040528251611e8a81611985565b81526020830151611e9a81611985565b8060208301525060408301516040820152606083015160608201526080830151608082015260a083015160a082015260c083015160c08201528091505092915050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615611f0d57611f0d611edd565b500290565b600060608284031215611f2457600080fd5b604051606081018181106001600160401b0382111715611f4657611f4661199a565b6040528251611f5481611985565b8152602083810151908201526040928301519281019290925250919050565b600060208284031215611f8557600080fd5b5051919050565b81810381811115611f9f57611f9f611edd565b92915050565b80820180821115611f9f57611f9f611edd56fea2646970667358221220bbbfedee2c7b0f7fdc5951ea8b537080a1ba2d6bae5637c58fa851c080ad60e764736f6c63430008100033"

// DeployControl3 deploys a new Ethereum contract, binding an instance of Control3 to it.
func DeployControl3(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, _inst common.Address) (common.Address, *types.Transaction, *Control3, error) {
	parsed, err := abi.JSON(strings.NewReader(Control3ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Control3Bin), backend, _a, _inst)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Control3{Control3Caller: Control3Caller{contract: contract}, Control3Transactor: Control3Transactor{contract: contract}, Control3Filterer: Control3Filterer{contract: contract}}, nil
}

// Control3 is an auto generated Go binding around an Ethereum contract.
type Control3 struct {
	Control3Caller     // Read-only binding to the contract
	Control3Transactor // Write-only binding to the contract
	Control3Filterer   // Log filterer for contract events
}

// Control3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Control3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Control3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Control3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Control3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Control3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Control3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Control3Session struct {
	Contract     *Control3         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Control3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Control3CallerSession struct {
	Contract *Control3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Control3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Control3TransactorSession struct {
	Contract     *Control3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Control3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Control3Raw struct {
	Contract *Control3 // Generic contract binding to access the raw methods on
}

// Control3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Control3CallerRaw struct {
	Contract *Control3Caller // Generic read-only contract binding to access the raw methods on
}

// Control3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Control3TransactorRaw struct {
	Contract *Control3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewControl3 creates a new instance of Control3, bound to a specific deployed contract.
func NewControl3(address common.Address, backend bind.ContractBackend) (*Control3, error) {
	contract, err := bindControl3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Control3{Control3Caller: Control3Caller{contract: contract}, Control3Transactor: Control3Transactor{contract: contract}, Control3Filterer: Control3Filterer{contract: contract}}, nil
}

// NewControl3Caller creates a new read-only instance of Control3, bound to a specific deployed contract.
func NewControl3Caller(address common.Address, caller bind.ContractCaller) (*Control3Caller, error) {
	contract, err := bindControl3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Control3Caller{contract: contract}, nil
}

// NewControl3Transactor creates a new write-only instance of Control3, bound to a specific deployed contract.
func NewControl3Transactor(address common.Address, transactor bind.ContractTransactor) (*Control3Transactor, error) {
	contract, err := bindControl3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Control3Transactor{contract: contract}, nil
}

// NewControl3Filterer creates a new log filterer instance of Control3, bound to a specific deployed contract.
func NewControl3Filterer(address common.Address, filterer bind.ContractFilterer) (*Control3Filterer, error) {
	contract, err := bindControl3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Control3Filterer{contract: contract}, nil
}

// bindControl3 binds a generic wrapper to an already deployed contract.
func bindControl3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Control3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control3 *Control3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Control3.Contract.Control3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control3 *Control3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Control3.Contract.Control3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control3 *Control3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Control3.Contract.Control3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control3 *Control3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Control3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control3 *Control3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Control3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control3 *Control3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Control3.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Control3 *Control3Caller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Control3.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Control3 *Control3Session) Auth() (common.Address, error) {
	return _Control3.Contract.Auth(&_Control3.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Control3 *Control3CallerSession) Auth() (common.Address, error) {
	return _Control3.Contract.Auth(&_Control3.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Control3 *Control3Caller) Inst(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Control3.contract.Call(opts, &out, "inst")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Control3 *Control3Session) Inst() (common.Address, error) {
	return _Control3.Contract.Inst(&_Control3.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Control3 *Control3CallerSession) Inst() (common.Address, error) {
	return _Control3.Contract.Inst(&_Control3.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Control3 *Control3Caller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Control3.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Control3 *Control3Session) Owners(arg0 common.Address) (bool, error) {
	return _Control3.Contract.Owners(&_Control3.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Control3 *Control3CallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Control3.Contract.Owners(&_Control3.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Control3 *Control3Caller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Control3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Control3 *Control3Session) Version() (uint16, error) {
	return _Control3.Contract.Version(&_Control3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Control3 *Control3CallerSession) Version() (uint16, error) {
	return _Control3.Contract.Version(&_Control3.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Control3 *Control3Transactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Control3 *Control3Session) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.Add(&_Control3.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Control3 *Control3TransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.Add(&_Control3.TransactOpts, _a, isSet, signs)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_Control3 *Control3Transactor) QuitRole(opts *bind.TransactOpts, _a common.Address, _i uint64) (*types.Transaction, error) {
	return _Control3.contract.Transact(opts, "quitRole", _a, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_Control3 *Control3Session) QuitRole(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _Control3.Contract.QuitRole(&_Control3.TransactOpts, _a, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_Control3 *Control3TransactorSession) QuitRole(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _Control3.Contract.QuitRole(&_Control3.TransactOpts, _a, _i)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_Control3 *Control3Transactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.contract.Transact(opts, "setLevel", _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_Control3 *Control3Session) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.SetLevel(&_Control3.TransactOpts, _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_Control3 *Control3TransactorSession) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.SetLevel(&_Control3.TransactOpts, _gi, _le, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_Control3 *Control3Transactor) SetPeriod(opts *bind.TransactOpts, _pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.contract.Transact(opts, "setPeriod", _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_Control3 *Control3Session) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.SetPeriod(&_Control3.TransactOpts, _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_Control3 *Control3TransactorSession) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.SetPeriod(&_Control3.TransactOpts, _pe, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_Control3 *Control3Transactor) SetPlePerB(opts *bind.TransactOpts, _gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.contract.Transact(opts, "setPlePerB", _gi, _kpPerB, _ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_Control3 *Control3Session) SetPlePerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.SetPlePerB(&_Control3.TransactOpts, _gi, _kpPerB, _ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_Control3 *Control3TransactorSession) SetPlePerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Control3.Contract.SetPlePerB(&_Control3.TransactOpts, _gi, _kpPerB, _ppPerB, signs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 money) returns()
func (_Control3 *Control3Transactor) Withdraw(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _Control3.contract.Transact(opts, "withdraw", _a, _i, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 money) returns()
func (_Control3 *Control3Session) Withdraw(_a common.Address, _i uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _Control3.Contract.Withdraw(&_Control3.TransactOpts, _a, _i, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 money) returns()
func (_Control3 *Control3TransactorSession) Withdraw(_a common.Address, _i uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _Control3.Contract.Withdraw(&_Control3.TransactOpts, _a, _i, _ti, money)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Control3 *Control3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Control3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Control3 *Control3Session) Receive() (*types.Transaction, error) {
	return _Control3.Contract.Receive(&_Control3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Control3 *Control3TransactorSession) Receive() (*types.Transaction, error) {
	return _Control3.Contract.Receive(&_Control3.TransactOpts)
}

// Control3AddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Control3 contract.
type Control3AddOwnerIterator struct {
	Event *Control3AddOwner // Event containing the contract specifics and raw log

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
func (it *Control3AddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Control3AddOwner)
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
		it.Event = new(Control3AddOwner)
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
func (it *Control3AddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Control3AddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Control3AddOwner represents a AddOwner event raised by the Control3 contract.
type Control3AddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Control3 *Control3Filterer) FilterAddOwner(opts *bind.FilterOpts) (*Control3AddOwnerIterator, error) {

	logs, sub, err := _Control3.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &Control3AddOwnerIterator{contract: _Control3.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Control3 *Control3Filterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *Control3AddOwner) (event.Subscription, error) {

	logs, sub, err := _Control3.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Control3AddOwner)
				if err := _Control3.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Control3 *Control3Filterer) ParseAddOwner(log types.Log) (*Control3AddOwner, error) {
	event := new(Control3AddOwner)
	if err := _Control3.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Control3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Control3 contract.
type Control3WithdrawIterator struct {
	Event *Control3Withdraw // Event containing the contract specifics and raw log

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
func (it *Control3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Control3Withdraw)
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
		it.Event = new(Control3Withdraw)
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
func (it *Control3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Control3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Control3Withdraw represents a Withdraw event raised by the Control3 contract.
type Control3Withdraw struct {
	Addr        common.Address
	I           uint64
	Ti          uint8
	Money       *big.Int
	ActualMoney *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_Control3 *Control3Filterer) FilterWithdraw(opts *bind.FilterOpts, addr []common.Address, i []uint64) (*Control3WithdrawIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _Control3.contract.FilterLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return &Control3WithdrawIterator{contract: _Control3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_Control3 *Control3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Control3Withdraw, addr []common.Address, i []uint64) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _Control3.contract.WatchLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Control3Withdraw)
				if err := _Control3.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_Control3 *Control3Filterer) ParseWithdraw(log types.Log) (*Control3Withdraw, error) {
	event := new(Control3Withdraw)
	if err := _Control3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuthABI is the input ABI used to generate the binding from.
const IAuthABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"perm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAuthFuncSigs maps the 4-byte function signature to its string representation.
var IAuthFuncSigs = map[string]string{
	"a96bba9d": "perm(bytes32,bytes[5])",
}

// IAuth is an auto generated Go binding around an Ethereum contract.
type IAuth struct {
	IAuthCaller     // Read-only binding to the contract
	IAuthTransactor // Write-only binding to the contract
	IAuthFilterer   // Log filterer for contract events
}

// IAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthSession struct {
	Contract     *IAuth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthCallerSession struct {
	Contract *IAuthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthTransactorSession struct {
	Contract     *IAuthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthRaw struct {
	Contract *IAuth // Generic contract binding to access the raw methods on
}

// IAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthCallerRaw struct {
	Contract *IAuthCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthTransactorRaw struct {
	Contract *IAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuth creates a new instance of IAuth, bound to a specific deployed contract.
func NewIAuth(address common.Address, backend bind.ContractBackend) (*IAuth, error) {
	contract, err := bindIAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuth{IAuthCaller: IAuthCaller{contract: contract}, IAuthTransactor: IAuthTransactor{contract: contract}, IAuthFilterer: IAuthFilterer{contract: contract}}, nil
}

// NewIAuthCaller creates a new read-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthCaller(address common.Address, caller bind.ContractCaller) (*IAuthCaller, error) {
	contract, err := bindIAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthCaller{contract: contract}, nil
}

// NewIAuthTransactor creates a new write-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthTransactor, error) {
	contract, err := bindIAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthTransactor{contract: contract}, nil
}

// NewIAuthFilterer creates a new log filterer instance of IAuth, bound to a specific deployed contract.
func NewIAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthFilterer, error) {
	contract, err := bindIAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthFilterer{contract: contract}, nil
}

// bindIAuth binds a generic wrapper to an already deployed contract.
func bindIAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.IAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transact(opts, method, params...)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactor) Perm(opts *bind.TransactOpts, _h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.contract.Transact(opts, "perm", _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactorSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// IControl3ABI is the input ABI used to generate the binding from.
const IControl3ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualMoney\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"quitRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pe\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPlePerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControl3FuncSigs maps the 4-byte function signature to its string representation.
var IControl3FuncSigs = map[string]string{
	"d6985699": "quitRole(address,uint64)",
	"713d3b95": "setLevel(uint64,uint8,bytes[5])",
	"bd5d54b7": "setPeriod(uint64,bytes[5])",
	"130b89b1": "setPlePerB(uint64,uint256,uint256,bytes[5])",
	"9b4c757a": "withdraw(address,uint64,uint8,uint256)",
}

// IControl3 is an auto generated Go binding around an Ethereum contract.
type IControl3 struct {
	IControl3Caller     // Read-only binding to the contract
	IControl3Transactor // Write-only binding to the contract
	IControl3Filterer   // Log filterer for contract events
}

// IControl3Caller is an auto generated read-only Go binding around an Ethereum contract.
type IControl3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IControl3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControl3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControl3Session struct {
	Contract     *IControl3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControl3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControl3CallerSession struct {
	Contract *IControl3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IControl3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControl3TransactorSession struct {
	Contract     *IControl3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IControl3Raw is an auto generated low-level Go binding around an Ethereum contract.
type IControl3Raw struct {
	Contract *IControl3 // Generic contract binding to access the raw methods on
}

// IControl3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControl3CallerRaw struct {
	Contract *IControl3Caller // Generic read-only contract binding to access the raw methods on
}

// IControl3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControl3TransactorRaw struct {
	Contract *IControl3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl3 creates a new instance of IControl3, bound to a specific deployed contract.
func NewIControl3(address common.Address, backend bind.ContractBackend) (*IControl3, error) {
	contract, err := bindIControl3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl3{IControl3Caller: IControl3Caller{contract: contract}, IControl3Transactor: IControl3Transactor{contract: contract}, IControl3Filterer: IControl3Filterer{contract: contract}}, nil
}

// NewIControl3Caller creates a new read-only instance of IControl3, bound to a specific deployed contract.
func NewIControl3Caller(address common.Address, caller bind.ContractCaller) (*IControl3Caller, error) {
	contract, err := bindIControl3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControl3Caller{contract: contract}, nil
}

// NewIControl3Transactor creates a new write-only instance of IControl3, bound to a specific deployed contract.
func NewIControl3Transactor(address common.Address, transactor bind.ContractTransactor) (*IControl3Transactor, error) {
	contract, err := bindIControl3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControl3Transactor{contract: contract}, nil
}

// NewIControl3Filterer creates a new log filterer instance of IControl3, bound to a specific deployed contract.
func NewIControl3Filterer(address common.Address, filterer bind.ContractFilterer) (*IControl3Filterer, error) {
	contract, err := bindIControl3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControl3Filterer{contract: contract}, nil
}

// bindIControl3 binds a generic wrapper to an already deployed contract.
func bindIControl3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControl3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl3 *IControl3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl3.Contract.IControl3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl3 *IControl3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl3.Contract.IControl3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl3 *IControl3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl3.Contract.IControl3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl3 *IControl3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl3 *IControl3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl3 *IControl3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl3.Contract.contract.Transact(opts, method, params...)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_IControl3 *IControl3Transactor) QuitRole(opts *bind.TransactOpts, _a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "quitRole", _a, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_IControl3 *IControl3Session) QuitRole(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl3.Contract.QuitRole(&_IControl3.TransactOpts, _a, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_IControl3 *IControl3TransactorSession) QuitRole(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl3.Contract.QuitRole(&_IControl3.TransactOpts, _a, _i)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IControl3 *IControl3Transactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "setLevel", _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IControl3 *IControl3Session) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetLevel(&_IControl3.TransactOpts, _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IControl3 *IControl3TransactorSession) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetLevel(&_IControl3.TransactOpts, _gi, _le, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IControl3 *IControl3Transactor) SetPeriod(opts *bind.TransactOpts, _pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "setPeriod", _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IControl3 *IControl3Session) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPeriod(&_IControl3.TransactOpts, _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IControl3 *IControl3TransactorSession) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPeriod(&_IControl3.TransactOpts, _pe, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_IControl3 *IControl3Transactor) SetPlePerB(opts *bind.TransactOpts, _gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "setPlePerB", _gi, _kpPerB, _ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_IControl3 *IControl3Session) SetPlePerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPlePerB(&_IControl3.TransactOpts, _gi, _kpPerB, _ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_IControl3 *IControl3TransactorSession) SetPlePerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPlePerB(&_IControl3.TransactOpts, _gi, _kpPerB, _ppPerB, signs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl3 *IControl3Transactor) Withdraw(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "withdraw", _a, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl3 *IControl3Session) Withdraw(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl3.Contract.Withdraw(&_IControl3.TransactOpts, _a, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl3 *IControl3TransactorSession) Withdraw(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl3.Contract.Withdraw(&_IControl3.TransactOpts, _a, _i, _ti, _money)
}

// IControl3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IControl3 contract.
type IControl3WithdrawIterator struct {
	Event *IControl3Withdraw // Event containing the contract specifics and raw log

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
func (it *IControl3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl3Withdraw)
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
		it.Event = new(IControl3Withdraw)
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
func (it *IControl3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl3Withdraw represents a Withdraw event raised by the IControl3 contract.
type IControl3Withdraw struct {
	Addr        common.Address
	I           uint64
	Ti          uint8
	Money       *big.Int
	ActualMoney *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl3 *IControl3Filterer) FilterWithdraw(opts *bind.FilterOpts, addr []common.Address, i []uint64) (*IControl3WithdrawIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl3.contract.FilterLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return &IControl3WithdrawIterator{contract: _IControl3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl3 *IControl3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IControl3Withdraw, addr []common.Address, i []uint64) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl3.contract.WatchLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl3Withdraw)
				if err := _IControl3.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl3 *IControl3Filterer) ParseWithdraw(log types.Log) (*IControl3Withdraw, error) {
	event := new(IControl3Withdraw)
	if err := _IControl3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysGetterABI is the input ABI used to generate the binding from.
const IFileSysGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"}],\"name\":\"getFsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structFsOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getGInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structGroupOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getSettleInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hasPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"canPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structSettleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getStoreInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structStoreOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IFileSysGetterFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysGetterFuncSigs = map[string]string{
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"59f7b572": "getFsInfo(uint64,uint64)",
	"5475cf8d": "getGInfo(uint64,uint8)",
	"64c64a4d": "getSettleInfo(uint64,uint8)",
	"8f328ecd": "getStoreInfo(uint64,uint64,uint8)",
}

// IFileSysGetter is an auto generated Go binding around an Ethereum contract.
type IFileSysGetter struct {
	IFileSysGetterCaller     // Read-only binding to the contract
	IFileSysGetterTransactor // Write-only binding to the contract
	IFileSysGetterFilterer   // Log filterer for contract events
}

// IFileSysGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysGetterSession struct {
	Contract     *IFileSysGetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysGetterCallerSession struct {
	Contract *IFileSysGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFileSysGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysGetterTransactorSession struct {
	Contract     *IFileSysGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFileSysGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysGetterRaw struct {
	Contract *IFileSysGetter // Generic contract binding to access the raw methods on
}

// IFileSysGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysGetterCallerRaw struct {
	Contract *IFileSysGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysGetterTransactorRaw struct {
	Contract *IFileSysGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSysGetter creates a new instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetter(address common.Address, backend bind.ContractBackend) (*IFileSysGetter, error) {
	contract, err := bindIFileSysGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetter{IFileSysGetterCaller: IFileSysGetterCaller{contract: contract}, IFileSysGetterTransactor: IFileSysGetterTransactor{contract: contract}, IFileSysGetterFilterer: IFileSysGetterFilterer{contract: contract}}, nil
}

// NewIFileSysGetterCaller creates a new read-only instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetterCaller(address common.Address, caller bind.ContractCaller) (*IFileSysGetterCaller, error) {
	contract, err := bindIFileSysGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetterCaller{contract: contract}, nil
}

// NewIFileSysGetterTransactor creates a new write-only instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysGetterTransactor, error) {
	contract, err := bindIFileSysGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetterTransactor{contract: contract}, nil
}

// NewIFileSysGetterFilterer creates a new log filterer instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysGetterFilterer, error) {
	contract, err := bindIFileSysGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetterFilterer{contract: contract}, nil
}

// bindIFileSysGetter binds a generic wrapper to an already deployed contract.
func bindIFileSysGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysGetter *IFileSysGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysGetter.Contract.IFileSysGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysGetter *IFileSysGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.IFileSysGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysGetter *IFileSysGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.IFileSysGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysGetter *IFileSysGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysGetter *IFileSysGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysGetter *IFileSysGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSysGetter *IFileSysGetterCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSysGetter *IFileSysGetterSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _IFileSysGetter.Contract.BalanceOf(&_IFileSysGetter.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSysGetter *IFileSysGetterCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _IFileSysGetter.Contract.BalanceOf(&_IFileSysGetter.CallOpts, _i, _ti)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSysGetter *IFileSysGetterCaller) GetFsInfo(opts *bind.CallOpts, _ui uint64, _pi uint64) (FsOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getFsInfo", _ui, _pi)

	if err != nil {
		return *new(FsOut), err
	}

	out0 := *abi.ConvertType(out[0], new(FsOut)).(*FsOut)

	return out0, err

}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSysGetter *IFileSysGetterSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _IFileSysGetter.Contract.GetFsInfo(&_IFileSysGetter.CallOpts, _ui, _pi)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _IFileSysGetter.Contract.GetFsInfo(&_IFileSysGetter.CallOpts, _ui, _pi)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCaller) GetGInfo(opts *bind.CallOpts, _gi uint64, _ti uint8) (GroupOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getGInfo", _gi, _ti)

	if err != nil {
		return *new(GroupOut), err
	}

	out0 := *abi.ConvertType(out[0], new(GroupOut)).(*GroupOut)

	return out0, err

}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _IFileSysGetter.Contract.GetGInfo(&_IFileSysGetter.CallOpts, _gi, _ti)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _IFileSysGetter.Contract.GetGInfo(&_IFileSysGetter.CallOpts, _gi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCaller) GetSettleInfo(opts *bind.CallOpts, _pi uint64, _ti uint8) (SettleOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getSettleInfo", _pi, _ti)

	if err != nil {
		return *new(SettleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(SettleOut)).(*SettleOut)

	return out0, err

}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _IFileSysGetter.Contract.GetSettleInfo(&_IFileSysGetter.CallOpts, _pi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _IFileSysGetter.Contract.GetSettleInfo(&_IFileSysGetter.CallOpts, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSysGetter *IFileSysGetterCaller) GetStoreInfo(opts *bind.CallOpts, _ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getStoreInfo", _ui, _pi, _ti)

	if err != nil {
		return *new(StoreOut), err
	}

	out0 := *abi.ConvertType(out[0], new(StoreOut)).(*StoreOut)

	return out0, err

}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSysGetter *IFileSysGetterSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _IFileSysGetter.Contract.GetStoreInfo(&_IFileSysGetter.CallOpts, _ui, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _IFileSysGetter.Contract.GetStoreInfo(&_IFileSysGetter.CallOpts, _ui, _pi, _ti)
}

// IFileSysSetterABI is the input ABI used to generate the binding from.
const IFileSysSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddRepair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"SubOrder\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kPB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPB\",\"type\":\"uint256\"}],\"name\":\"addPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"proWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IFileSysSetterFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysSetterFuncSigs = map[string]string{
	"9058d8e7": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint8,uint8,uint64)",
	"2449a168": "addPenalty(uint64,uint64,uint8,uint8,uint256,uint256)",
	"80faaf88": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"a4703e16": "proWithdraw((uint64,uint8,uint256,uint256),uint64)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"248d02a0": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IFileSysSetter is an auto generated Go binding around an Ethereum contract.
type IFileSysSetter struct {
	IFileSysSetterCaller     // Read-only binding to the contract
	IFileSysSetterTransactor // Write-only binding to the contract
	IFileSysSetterFilterer   // Log filterer for contract events
}

// IFileSysSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysSetterSession struct {
	Contract     *IFileSysSetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysSetterCallerSession struct {
	Contract *IFileSysSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFileSysSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysSetterTransactorSession struct {
	Contract     *IFileSysSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFileSysSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysSetterRaw struct {
	Contract *IFileSysSetter // Generic contract binding to access the raw methods on
}

// IFileSysSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysSetterCallerRaw struct {
	Contract *IFileSysSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactorRaw struct {
	Contract *IFileSysSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSysSetter creates a new instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetter(address common.Address, backend bind.ContractBackend) (*IFileSysSetter, error) {
	contract, err := bindIFileSysSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetter{IFileSysSetterCaller: IFileSysSetterCaller{contract: contract}, IFileSysSetterTransactor: IFileSysSetterTransactor{contract: contract}, IFileSysSetterFilterer: IFileSysSetterFilterer{contract: contract}}, nil
}

// NewIFileSysSetterCaller creates a new read-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterCaller(address common.Address, caller bind.ContractCaller) (*IFileSysSetterCaller, error) {
	contract, err := bindIFileSysSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterCaller{contract: contract}, nil
}

// NewIFileSysSetterTransactor creates a new write-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysSetterTransactor, error) {
	contract, err := bindIFileSysSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterTransactor{contract: contract}, nil
}

// NewIFileSysSetterFilterer creates a new log filterer instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysSetterFilterer, error) {
	contract, err := bindIFileSysSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterFilterer{contract: contract}, nil
}

// bindIFileSysSetter binds a generic wrapper to an already deployed contract.
func bindIFileSysSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.IFileSysSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddOrder(opts *bind.TransactOpts, ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addOrder", ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _tr, _gi)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddPenalty(opts *bind.TransactOpts, _gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addPenalty", _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddPenalty(&_IFileSysSetter.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddPenalty(&_IFileSysSetter.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddRepair(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addRepair", ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "proWithdraw", ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "recharge", _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) SubOrder(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "subOrder", ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "withdraw", _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money, _lock)
}

// IFileSysSetterAddOrderIterator is returned from FilterAddOrder and is used to iterate over the raw logs and unpacked data for AddOrder events raised by the IFileSysSetter contract.
type IFileSysSetterAddOrderIterator struct {
	Event *IFileSysSetterAddOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddOrder)
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
		it.Event = new(IFileSysSetterAddOrder)
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
func (it *IFileSysSetterAddOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddOrder represents a AddOrder event raised by the IFileSysSetter contract.
type IFileSysSetterAddOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddOrder is a free log retrieval operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddOrderIterator{contract: _IFileSysSetter.contract, event: "AddOrder", logs: logs, sub: sub}, nil
}

// WatchAddOrder is a free log subscription operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
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

// ParseAddOrder is a log parse operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddOrder(log types.Log) (*IFileSysSetterAddOrder, error) {
	event := new(IFileSysSetterAddOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterAddRepairIterator is returned from FilterAddRepair and is used to iterate over the raw logs and unpacked data for AddRepair events raised by the IFileSysSetter contract.
type IFileSysSetterAddRepairIterator struct {
	Event *IFileSysSetterAddRepair // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddRepairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddRepair)
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
		it.Event = new(IFileSysSetterAddRepair)
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
func (it *IFileSysSetterAddRepairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddRepairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddRepair represents a AddRepair event raised by the IFileSysSetter contract.
type IFileSysSetterAddRepair struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddRepair is a free log retrieval operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddRepair(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddRepairIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddRepairIterator{contract: _IFileSysSetter.contract, event: "AddRepair", logs: logs, sub: sub}, nil
}

// WatchAddRepair is a free log subscription operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddRepair(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddRepair, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddRepair)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
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

// ParseAddRepair is a log parse operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddRepair(log types.Log) (*IFileSysSetterAddRepair, error) {
	event := new(IFileSysSetterAddRepair)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterSubOrderIterator is returned from FilterSubOrder and is used to iterate over the raw logs and unpacked data for SubOrder events raised by the IFileSysSetter contract.
type IFileSysSetterSubOrderIterator struct {
	Event *IFileSysSetterSubOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterSubOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterSubOrder)
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
		it.Event = new(IFileSysSetterSubOrder)
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
func (it *IFileSysSetterSubOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterSubOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterSubOrder represents a SubOrder event raised by the IFileSysSetter contract.
type IFileSysSetterSubOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubOrder is a free log retrieval operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterSubOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterSubOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterSubOrderIterator{contract: _IFileSysSetter.contract, event: "SubOrder", logs: logs, sub: sub}, nil
}

// WatchSubOrder is a free log subscription operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchSubOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterSubOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterSubOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
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

// ParseSubOrder is a log parse operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseSubOrder(log types.Log) (*IFileSysSetterSubOrder, error) {
	event := new(IFileSysSetterSubOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGroupGetterABI is the input ABI used to generate the binding from.
const IGroupGetterABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_pm\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getDesc\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getPInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getPlePerB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IGroupGetterFuncSigs maps the 4-byte function signature to its string representation.
var IGroupGetterFuncSigs = map[string]string{
	"4acc2468": "add(uint8,uint64,uint256)",
	"2a2722ad": "getDesc(uint64)",
	"059e783d": "getGCnt()",
	"f6070374": "getGS(uint64)",
	"770609a8": "getLevel(uint64)",
	"50aa2329": "getPInfo(uint64)",
	"60b5903d": "getPlePerB(uint64)",
	"e795ea8a": "getRate(uint64)",
}

// IGroupGetter is an auto generated Go binding around an Ethereum contract.
type IGroupGetter struct {
	IGroupGetterCaller     // Read-only binding to the contract
	IGroupGetterTransactor // Write-only binding to the contract
	IGroupGetterFilterer   // Log filterer for contract events
}

// IGroupGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGroupGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGroupGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGroupGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGroupGetterSession struct {
	Contract     *IGroupGetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGroupGetterCallerSession struct {
	Contract *IGroupGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGroupGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGroupGetterTransactorSession struct {
	Contract     *IGroupGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGroupGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGroupGetterRaw struct {
	Contract *IGroupGetter // Generic contract binding to access the raw methods on
}

// IGroupGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGroupGetterCallerRaw struct {
	Contract *IGroupGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IGroupGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGroupGetterTransactorRaw struct {
	Contract *IGroupGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGroupGetter creates a new instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetter(address common.Address, backend bind.ContractBackend) (*IGroupGetter, error) {
	contract, err := bindIGroupGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGroupGetter{IGroupGetterCaller: IGroupGetterCaller{contract: contract}, IGroupGetterTransactor: IGroupGetterTransactor{contract: contract}, IGroupGetterFilterer: IGroupGetterFilterer{contract: contract}}, nil
}

// NewIGroupGetterCaller creates a new read-only instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetterCaller(address common.Address, caller bind.ContractCaller) (*IGroupGetterCaller, error) {
	contract, err := bindIGroupGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupGetterCaller{contract: contract}, nil
}

// NewIGroupGetterTransactor creates a new write-only instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IGroupGetterTransactor, error) {
	contract, err := bindIGroupGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupGetterTransactor{contract: contract}, nil
}

// NewIGroupGetterFilterer creates a new log filterer instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IGroupGetterFilterer, error) {
	contract, err := bindIGroupGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGroupGetterFilterer{contract: contract}, nil
}

// bindIGroupGetter binds a generic wrapper to an already deployed contract.
func bindIGroupGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGroupGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupGetter *IGroupGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupGetter.Contract.IGroupGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupGetter *IGroupGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupGetter.Contract.IGroupGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupGetter *IGroupGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupGetter.Contract.IGroupGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupGetter *IGroupGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupGetter *IGroupGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupGetter *IGroupGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupGetter.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroupGetter *IGroupGetterCaller) Add(opts *bind.CallOpts, _rType uint8, _gi uint64, _pm *big.Int) error {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "add", _rType, _gi, _pm)

	if err != nil {
		return err
	}

	return err

}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroupGetter *IGroupGetterSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroupGetter.Contract.Add(&_IGroupGetter.CallOpts, _rType, _gi, _pm)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroupGetter *IGroupGetterCallerSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroupGetter.Contract.Add(&_IGroupGetter.CallOpts, _rType, _gi, _pm)
}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_IGroupGetter *IGroupGetterCaller) GetDesc(opts *bind.CallOpts, i uint64) ([]byte, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getDesc", i)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_IGroupGetter *IGroupGetterSession) GetDesc(i uint64) ([]byte, error) {
	return _IGroupGetter.Contract.GetDesc(&_IGroupGetter.CallOpts, i)
}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_IGroupGetter *IGroupGetterCallerSession) GetDesc(i uint64) ([]byte, error) {
	return _IGroupGetter.Contract.GetDesc(&_IGroupGetter.CallOpts, i)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroupGetter *IGroupGetterCaller) GetGCnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getGCnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroupGetter *IGroupGetterSession) GetGCnt() (uint64, error) {
	return _IGroupGetter.Contract.GetGCnt(&_IGroupGetter.CallOpts)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroupGetter *IGroupGetterCallerSession) GetGCnt() (uint64, error) {
	return _IGroupGetter.Contract.GetGCnt(&_IGroupGetter.CallOpts)
}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterCaller) GetGS(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getGS", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterSession) GetGS(i uint64) (uint8, error) {
	return _IGroupGetter.Contract.GetGS(&_IGroupGetter.CallOpts, i)
}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterCallerSession) GetGS(i uint64) (uint8, error) {
	return _IGroupGetter.Contract.GetGS(&_IGroupGetter.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterCaller) GetLevel(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getLevel", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterSession) GetLevel(i uint64) (uint8, error) {
	return _IGroupGetter.Contract.GetLevel(&_IGroupGetter.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterCallerSession) GetLevel(i uint64) (uint8, error) {
	return _IGroupGetter.Contract.GetLevel(&_IGroupGetter.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterCaller) GetPInfo(opts *bind.CallOpts, _i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getPInfo", _i)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroupGetter.Contract.GetPInfo(&_IGroupGetter.CallOpts, _i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterCallerSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroupGetter.Contract.GetPInfo(&_IGroupGetter.CallOpts, _i)
}

// GetPlePerB is a free data retrieval call binding the contract method 0x60b5903d.
//
// Solidity: function getPlePerB(uint64 i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterCaller) GetPlePerB(opts *bind.CallOpts, i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getPlePerB", i)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPlePerB is a free data retrieval call binding the contract method 0x60b5903d.
//
// Solidity: function getPlePerB(uint64 i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterSession) GetPlePerB(i uint64) (*big.Int, *big.Int, error) {
	return _IGroupGetter.Contract.GetPlePerB(&_IGroupGetter.CallOpts, i)
}

// GetPlePerB is a free data retrieval call binding the contract method 0x60b5903d.
//
// Solidity: function getPlePerB(uint64 i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterCallerSession) GetPlePerB(i uint64) (*big.Int, *big.Int, error) {
	return _IGroupGetter.Contract.GetPlePerB(&_IGroupGetter.CallOpts, i)
}

// GetRate is a free data retrieval call binding the contract method 0xe795ea8a.
//
// Solidity: function getRate(uint64 i) view returns(uint8, uint8)
func (_IGroupGetter *IGroupGetterCaller) GetRate(opts *bind.CallOpts, i uint64) (uint8, uint8, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getRate", i)

	if err != nil {
		return *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetRate is a free data retrieval call binding the contract method 0xe795ea8a.
//
// Solidity: function getRate(uint64 i) view returns(uint8, uint8)
func (_IGroupGetter *IGroupGetterSession) GetRate(i uint64) (uint8, uint8, error) {
	return _IGroupGetter.Contract.GetRate(&_IGroupGetter.CallOpts, i)
}

// GetRate is a free data retrieval call binding the contract method 0xe795ea8a.
//
// Solidity: function getRate(uint64 i) view returns(uint8, uint8)
func (_IGroupGetter *IGroupGetterCallerSession) GetRate(i uint64) (uint8, uint8, error) {
	return _IGroupGetter.Contract.GetRate(&_IGroupGetter.CallOpts, i)
}

// IGroupSetterABI is the input ABI used to generate the binding from.
const IGroupSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"Create\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_kc\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_kPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"setDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pp\",\"type\":\"uint256\"}],\"name\":\"setP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"}],\"name\":\"setPerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"kcnt\",\"type\":\"uint8\"}],\"name\":\"setS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGroupSetterFuncSigs maps the 4-byte function signature to its string representation.
var IGroupSetterFuncSigs = map[string]string{
	"ad578aeb": "activate(uint64,uint8)",
	"ae38ca77": "create(uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes)",
	"773854ed": "setDesc(uint64,bytes)",
	"936945bd": "setLevel(uint64,uint8)",
	"2a0212e5": "setP(uint64,uint256,uint256)",
	"cfb9854f": "setPerB(uint64,uint256,uint256)",
	"1e752182": "setRate(uint64,uint8,uint8)",
	"5f9dbbf9": "setS(uint64,uint8,uint8)",
}

// IGroupSetter is an auto generated Go binding around an Ethereum contract.
type IGroupSetter struct {
	IGroupSetterCaller     // Read-only binding to the contract
	IGroupSetterTransactor // Write-only binding to the contract
	IGroupSetterFilterer   // Log filterer for contract events
}

// IGroupSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGroupSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGroupSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGroupSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGroupSetterSession struct {
	Contract     *IGroupSetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGroupSetterCallerSession struct {
	Contract *IGroupSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGroupSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGroupSetterTransactorSession struct {
	Contract     *IGroupSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGroupSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGroupSetterRaw struct {
	Contract *IGroupSetter // Generic contract binding to access the raw methods on
}

// IGroupSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGroupSetterCallerRaw struct {
	Contract *IGroupSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IGroupSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGroupSetterTransactorRaw struct {
	Contract *IGroupSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGroupSetter creates a new instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetter(address common.Address, backend bind.ContractBackend) (*IGroupSetter, error) {
	contract, err := bindIGroupSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGroupSetter{IGroupSetterCaller: IGroupSetterCaller{contract: contract}, IGroupSetterTransactor: IGroupSetterTransactor{contract: contract}, IGroupSetterFilterer: IGroupSetterFilterer{contract: contract}}, nil
}

// NewIGroupSetterCaller creates a new read-only instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetterCaller(address common.Address, caller bind.ContractCaller) (*IGroupSetterCaller, error) {
	contract, err := bindIGroupSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupSetterCaller{contract: contract}, nil
}

// NewIGroupSetterTransactor creates a new write-only instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IGroupSetterTransactor, error) {
	contract, err := bindIGroupSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupSetterTransactor{contract: contract}, nil
}

// NewIGroupSetterFilterer creates a new log filterer instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IGroupSetterFilterer, error) {
	contract, err := bindIGroupSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGroupSetterFilterer{contract: contract}, nil
}

// bindIGroupSetter binds a generic wrapper to an already deployed contract.
func bindIGroupSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGroupSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupSetter *IGroupSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupSetter.Contract.IGroupSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupSetter *IGroupSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupSetter.Contract.IGroupSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupSetter *IGroupSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupSetter.Contract.IGroupSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupSetter *IGroupSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupSetter *IGroupSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupSetter *IGroupSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupSetter.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroupSetter *IGroupSetterTransactor) Activate(opts *bind.TransactOpts, _gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "activate", _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroupSetter *IGroupSetterSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Activate(&_IGroupSetter.TransactOpts, _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Activate(&_IGroupSetter.TransactOpts, _gi, _kc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kPerB, uint256 _pPerB, bytes _desc) returns()
func (_IGroupSetter *IGroupSetterTransactor) Create(opts *bind.TransactOpts, _level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kPerB *big.Int, _pPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "create", _level, _mr, _tr, _kr, _pr, _kPerB, _pPerB, _desc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kPerB, uint256 _pPerB, bytes _desc) returns()
func (_IGroupSetter *IGroupSetterSession) Create(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kPerB *big.Int, _pPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Create(&_IGroupSetter.TransactOpts, _level, _mr, _tr, _kr, _pr, _kPerB, _pPerB, _desc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kPerB, uint256 _pPerB, bytes _desc) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) Create(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kPerB *big.Int, _pPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Create(&_IGroupSetter.TransactOpts, _level, _mr, _tr, _kr, _pr, _kPerB, _pPerB, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_IGroupSetter *IGroupSetterTransactor) SetDesc(opts *bind.TransactOpts, _gi uint64, _desc []byte) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "setDesc", _gi, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_IGroupSetter *IGroupSetterSession) SetDesc(_gi uint64, _desc []byte) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetDesc(&_IGroupSetter.TransactOpts, _gi, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) SetDesc(_gi uint64, _desc []byte) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetDesc(&_IGroupSetter.TransactOpts, _gi, _desc)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_IGroupSetter *IGroupSetterTransactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "setLevel", _gi, _le)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_IGroupSetter *IGroupSetterSession) SetLevel(_gi uint64, _le uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetLevel(&_IGroupSetter.TransactOpts, _gi, _le)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) SetLevel(_gi uint64, _le uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetLevel(&_IGroupSetter.TransactOpts, _gi, _le)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_IGroupSetter *IGroupSetterTransactor) SetP(opts *bind.TransactOpts, _gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "setP", _gi, _kp, _pp)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_IGroupSetter *IGroupSetterSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetP(&_IGroupSetter.TransactOpts, _gi, _kp, _pp)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetP(&_IGroupSetter.TransactOpts, _gi, _kp, _pp)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_IGroupSetter *IGroupSetterTransactor) SetPerB(opts *bind.TransactOpts, _gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "setPerB", _gi, _kpPerB, _ppPerB)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_IGroupSetter *IGroupSetterSession) SetPerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetPerB(&_IGroupSetter.TransactOpts, _gi, _kpPerB, _ppPerB)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) SetPerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetPerB(&_IGroupSetter.TransactOpts, _gi, _kpPerB, _ppPerB)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_IGroupSetter *IGroupSetterTransactor) SetRate(opts *bind.TransactOpts, _gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "setRate", _gi, _mr, _tr)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_IGroupSetter *IGroupSetterSession) SetRate(_gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetRate(&_IGroupSetter.TransactOpts, _gi, _mr, _tr)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) SetRate(_gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetRate(&_IGroupSetter.TransactOpts, _gi, _mr, _tr)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_IGroupSetter *IGroupSetterTransactor) SetS(opts *bind.TransactOpts, _gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "setS", _gi, _s, kcnt)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_IGroupSetter *IGroupSetterSession) SetS(_gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetS(&_IGroupSetter.TransactOpts, _gi, _s, kcnt)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) SetS(_gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.SetS(&_IGroupSetter.TransactOpts, _gi, _s, kcnt)
}

// IGroupSetterCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the IGroupSetter contract.
type IGroupSetterCreateIterator struct {
	Event *IGroupSetterCreate // Event containing the contract specifics and raw log

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
func (it *IGroupSetterCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGroupSetterCreate)
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
		it.Event = new(IGroupSetterCreate)
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
func (it *IGroupSetterCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGroupSetterCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGroupSetterCreate represents a Create event raised by the IGroupSetter contract.
type IGroupSetterCreate struct {
	GIndex uint64
	Level  uint8
	Mr     uint8
	Tr     uint8
	Kr     *big.Int
	Pr     *big.Int
	KPerB  *big.Int
	PPerB  *big.Int
	Desc   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x4b14fb8b0a4d249a53bfa9dadcdeab4977149168de6dd5efa8a37e41107b3fda.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 tr, uint256 kr, uint256 pr, uint256 kPerB, uint256 pPerB, bytes _desc)
func (_IGroupSetter *IGroupSetterFilterer) FilterCreate(opts *bind.FilterOpts) (*IGroupSetterCreateIterator, error) {

	logs, sub, err := _IGroupSetter.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &IGroupSetterCreateIterator{contract: _IGroupSetter.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x4b14fb8b0a4d249a53bfa9dadcdeab4977149168de6dd5efa8a37e41107b3fda.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 tr, uint256 kr, uint256 pr, uint256 kPerB, uint256 pPerB, bytes _desc)
func (_IGroupSetter *IGroupSetterFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *IGroupSetterCreate) (event.Subscription, error) {

	logs, sub, err := _IGroupSetter.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGroupSetterCreate)
				if err := _IGroupSetter.contract.UnpackLog(event, "Create", log); err != nil {
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

// ParseCreate is a log parse operation binding the contract event 0x4b14fb8b0a4d249a53bfa9dadcdeab4977149168de6dd5efa8a37e41107b3fda.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 tr, uint256 kr, uint256 pr, uint256 kPerB, uint256 pPerB, bytes _desc)
func (_IGroupSetter *IGroupSetterFilterer) ParseCreate(log types.Log) (*IGroupSetterCreate, error) {
	event := new(IGroupSetterCreate)
	if err := _IGroupSetter.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInstanceABI is the input ABI used to generate the binding from.
const IInstanceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Alter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"instances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInstanceFuncSigs maps the 4-byte function signature to its string representation.
var IInstanceFuncSigs = map[string]string{
	"3ec7d5b9": "instances(uint8)",
}

// IInstance is an auto generated Go binding around an Ethereum contract.
type IInstance struct {
	IInstanceCaller     // Read-only binding to the contract
	IInstanceTransactor // Write-only binding to the contract
	IInstanceFilterer   // Log filterer for contract events
}

// IInstanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInstanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInstanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInstanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInstanceSession struct {
	Contract     *IInstance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInstanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInstanceCallerSession struct {
	Contract *IInstanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IInstanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInstanceTransactorSession struct {
	Contract     *IInstanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInstanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInstanceRaw struct {
	Contract *IInstance // Generic contract binding to access the raw methods on
}

// IInstanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInstanceCallerRaw struct {
	Contract *IInstanceCaller // Generic read-only contract binding to access the raw methods on
}

// IInstanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInstanceTransactorRaw struct {
	Contract *IInstanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInstance creates a new instance of IInstance, bound to a specific deployed contract.
func NewIInstance(address common.Address, backend bind.ContractBackend) (*IInstance, error) {
	contract, err := bindIInstance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInstance{IInstanceCaller: IInstanceCaller{contract: contract}, IInstanceTransactor: IInstanceTransactor{contract: contract}, IInstanceFilterer: IInstanceFilterer{contract: contract}}, nil
}

// NewIInstanceCaller creates a new read-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceCaller(address common.Address, caller bind.ContractCaller) (*IInstanceCaller, error) {
	contract, err := bindIInstance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceCaller{contract: contract}, nil
}

// NewIInstanceTransactor creates a new write-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IInstanceTransactor, error) {
	contract, err := bindIInstance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceTransactor{contract: contract}, nil
}

// NewIInstanceFilterer creates a new log filterer instance of IInstance, bound to a specific deployed contract.
func NewIInstanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IInstanceFilterer, error) {
	contract, err := bindIInstance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInstanceFilterer{contract: contract}, nil
}

// bindIInstance binds a generic wrapper to an already deployed contract.
func bindIInstance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInstanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.IInstanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transact(opts, method, params...)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCaller) Instances(opts *bind.CallOpts, _type uint8) (common.Address, error) {
	var out []interface{}
	err := _IInstance.contract.Call(opts, &out, "instances", _type)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCallerSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// IInstanceAlterIterator is returned from FilterAlter and is used to iterate over the raw logs and unpacked data for Alter events raised by the IInstance contract.
type IInstanceAlterIterator struct {
	Event *IInstanceAlter // Event containing the contract specifics and raw log

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
func (it *IInstanceAlterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInstanceAlter)
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
		it.Event = new(IInstanceAlter)
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
func (it *IInstanceAlterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInstanceAlterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInstanceAlter represents a Alter event raised by the IInstance contract.
type IInstanceAlter struct {
	Type uint8
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlter is a free log retrieval operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) FilterAlter(opts *bind.FilterOpts) (*IInstanceAlterIterator, error) {

	logs, sub, err := _IInstance.contract.FilterLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return &IInstanceAlterIterator{contract: _IInstance.contract, event: "Alter", logs: logs, sub: sub}, nil
}

// WatchAlter is a free log subscription operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) WatchAlter(opts *bind.WatchOpts, sink chan<- *IInstanceAlter) (event.Subscription, error) {

	logs, sub, err := _IInstance.contract.WatchLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInstanceAlter)
				if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
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

// ParseAlter is a log parse operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) ParseAlter(log types.Log) (*IInstanceAlter, error) {
	event := new(IInstanceAlter)
	if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageABI is the input ABI used to generate the binding from.
const IKmanageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pe\",\"type\":\"uint64\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKmanageFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageFuncSigs = map[string]string{
	"49d2a1b4": "addCnt(uint64,uint64,uint64)",
	"a666bc94": "addKeeper(uint64,uint64)",
	"2d1664e9": "addProfit(uint8,uint64,uint256)",
	"f57d6e17": "balanceOf(uint64,uint64,uint8)",
	"da41f894": "getK(uint64,uint64)",
	"4cfd1e2d": "getKCnt(uint64)",
	"681e4819": "getPf(uint8,uint64)",
	"309cb5ea": "setPeriod(uint64)",
	"9cad4378": "withdraw(uint64,uint8,uint64,uint256)",
}

// IKmanage is an auto generated Go binding around an Ethereum contract.
type IKmanage struct {
	IKmanageCaller     // Read-only binding to the contract
	IKmanageTransactor // Write-only binding to the contract
	IKmanageFilterer   // Log filterer for contract events
}

// IKmanageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageSession struct {
	Contract     *IKmanage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageCallerSession struct {
	Contract *IKmanageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IKmanageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageTransactorSession struct {
	Contract     *IKmanageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IKmanageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageRaw struct {
	Contract *IKmanage // Generic contract binding to access the raw methods on
}

// IKmanageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageCallerRaw struct {
	Contract *IKmanageCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageTransactorRaw struct {
	Contract *IKmanageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanage creates a new instance of IKmanage, bound to a specific deployed contract.
func NewIKmanage(address common.Address, backend bind.ContractBackend) (*IKmanage, error) {
	contract, err := bindIKmanage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanage{IKmanageCaller: IKmanageCaller{contract: contract}, IKmanageTransactor: IKmanageTransactor{contract: contract}, IKmanageFilterer: IKmanageFilterer{contract: contract}}, nil
}

// NewIKmanageCaller creates a new read-only instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageCaller(address common.Address, caller bind.ContractCaller) (*IKmanageCaller, error) {
	contract, err := bindIKmanage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageCaller{contract: contract}, nil
}

// NewIKmanageTransactor creates a new write-only instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageTransactor, error) {
	contract, err := bindIKmanage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageTransactor{contract: contract}, nil
}

// NewIKmanageFilterer creates a new log filterer instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageFilterer, error) {
	contract, err := bindIKmanage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageFilterer{contract: contract}, nil
}

// bindIKmanage binds a generic wrapper to an already deployed contract.
func bindIKmanage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanage *IKmanageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanage.Contract.IKmanageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanage *IKmanageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanage.Contract.IKmanageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanage *IKmanageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanage.Contract.IKmanageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanage *IKmanageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanage *IKmanageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanage *IKmanageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanage.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf57d6e17.
//
// Solidity: function balanceOf(uint64 _gi, uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageCaller) BalanceOf(opts *bind.CallOpts, _gi uint64, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "balanceOf", _gi, _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf57d6e17.
//
// Solidity: function balanceOf(uint64 _gi, uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageSession) BalanceOf(_gi uint64, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanage.Contract.BalanceOf(&_IKmanage.CallOpts, _gi, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf57d6e17.
//
// Solidity: function balanceOf(uint64 _gi, uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageCallerSession) BalanceOf(_gi uint64, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanage.Contract.BalanceOf(&_IKmanage.CallOpts, _gi, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xda41f894.
//
// Solidity: function getK(uint64 _gi, uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageCaller) GetK(opts *bind.CallOpts, _gi uint64, _i uint64) (uint64, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getK", _gi, _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xda41f894.
//
// Solidity: function getK(uint64 _gi, uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageSession) GetK(_gi uint64, _i uint64) (uint64, error) {
	return _IKmanage.Contract.GetK(&_IKmanage.CallOpts, _gi, _i)
}

// GetK is a free data retrieval call binding the contract method 0xda41f894.
//
// Solidity: function getK(uint64 _gi, uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageCallerSession) GetK(_gi uint64, _i uint64) (uint64, error) {
	return _IKmanage.Contract.GetK(&_IKmanage.CallOpts, _gi, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x4cfd1e2d.
//
// Solidity: function getKCnt(uint64 _gi) view returns(uint8)
func (_IKmanage *IKmanageCaller) GetKCnt(opts *bind.CallOpts, _gi uint64) (uint8, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getKCnt", _gi)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x4cfd1e2d.
//
// Solidity: function getKCnt(uint64 _gi) view returns(uint8)
func (_IKmanage *IKmanageSession) GetKCnt(_gi uint64) (uint8, error) {
	return _IKmanage.Contract.GetKCnt(&_IKmanage.CallOpts, _gi)
}

// GetKCnt is a free data retrieval call binding the contract method 0x4cfd1e2d.
//
// Solidity: function getKCnt(uint64 _gi) view returns(uint8)
func (_IKmanage *IKmanageCallerSession) GetKCnt(_gi uint64) (uint8, error) {
	return _IKmanage.Contract.GetKCnt(&_IKmanage.CallOpts, _gi)
}

// GetPf is a free data retrieval call binding the contract method 0x681e4819.
//
// Solidity: function getPf(uint8 _ti, uint64 _gi) view returns(uint64, uint256)
func (_IKmanage *IKmanageCaller) GetPf(opts *bind.CallOpts, _ti uint8, _gi uint64) (uint64, *big.Int, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getPf", _ti, _gi)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x681e4819.
//
// Solidity: function getPf(uint8 _ti, uint64 _gi) view returns(uint64, uint256)
func (_IKmanage *IKmanageSession) GetPf(_ti uint8, _gi uint64) (uint64, *big.Int, error) {
	return _IKmanage.Contract.GetPf(&_IKmanage.CallOpts, _ti, _gi)
}

// GetPf is a free data retrieval call binding the contract method 0x681e4819.
//
// Solidity: function getPf(uint8 _ti, uint64 _gi) view returns(uint64, uint256)
func (_IKmanage *IKmanageCallerSession) GetPf(_ti uint8, _gi uint64) (uint64, *big.Int, error) {
	return _IKmanage.Contract.GetPf(&_IKmanage.CallOpts, _ti, _gi)
}

// AddCnt is a paid mutator transaction binding the contract method 0x49d2a1b4.
//
// Solidity: function addCnt(uint64 _gi, uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageTransactor) AddCnt(opts *bind.TransactOpts, _gi uint64, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addCnt", _gi, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x49d2a1b4.
//
// Solidity: function addCnt(uint64 _gi, uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageSession) AddCnt(_gi uint64, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddCnt(&_IKmanage.TransactOpts, _gi, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x49d2a1b4.
//
// Solidity: function addCnt(uint64 _gi, uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageTransactorSession) AddCnt(_gi uint64, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddCnt(&_IKmanage.TransactOpts, _gi, _ki, _cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0xa666bc94.
//
// Solidity: function addKeeper(uint64 _gi, uint64 _ki) returns()
func (_IKmanage *IKmanageTransactor) AddKeeper(opts *bind.TransactOpts, _gi uint64, _ki uint64) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addKeeper", _gi, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0xa666bc94.
//
// Solidity: function addKeeper(uint64 _gi, uint64 _ki) returns()
func (_IKmanage *IKmanageSession) AddKeeper(_gi uint64, _ki uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddKeeper(&_IKmanage.TransactOpts, _gi, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0xa666bc94.
//
// Solidity: function addKeeper(uint64 _gi, uint64 _ki) returns()
func (_IKmanage *IKmanageTransactorSession) AddKeeper(_gi uint64, _ki uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddKeeper(&_IKmanage.TransactOpts, _gi, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x2d1664e9.
//
// Solidity: function addProfit(uint8 _ti, uint64 _gi, uint256 _money) returns()
func (_IKmanage *IKmanageTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _gi uint64, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addProfit", _ti, _gi, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x2d1664e9.
//
// Solidity: function addProfit(uint8 _ti, uint64 _gi, uint256 _money) returns()
func (_IKmanage *IKmanageSession) AddProfit(_ti uint8, _gi uint64, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.AddProfit(&_IKmanage.TransactOpts, _ti, _gi, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x2d1664e9.
//
// Solidity: function addProfit(uint8 _ti, uint64 _gi, uint256 _money) returns()
func (_IKmanage *IKmanageTransactorSession) AddProfit(_ti uint8, _gi uint64, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.AddProfit(&_IKmanage.TransactOpts, _ti, _gi, _money)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x309cb5ea.
//
// Solidity: function setPeriod(uint64 _pe) returns()
func (_IKmanage *IKmanageTransactor) SetPeriod(opts *bind.TransactOpts, _pe uint64) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "setPeriod", _pe)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x309cb5ea.
//
// Solidity: function setPeriod(uint64 _pe) returns()
func (_IKmanage *IKmanageSession) SetPeriod(_pe uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.SetPeriod(&_IKmanage.TransactOpts, _pe)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x309cb5ea.
//
// Solidity: function setPeriod(uint64 _pe) returns()
func (_IKmanage *IKmanageTransactorSession) SetPeriod(_pe uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.SetPeriod(&_IKmanage.TransactOpts, _pe)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9cad4378.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint64 _gi, uint256 money) returns(uint256)
func (_IKmanage *IKmanageTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, _gi uint64, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "withdraw", _ki, _ti, _gi, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9cad4378.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint64 _gi, uint256 money) returns(uint256)
func (_IKmanage *IKmanageSession) Withdraw(_ki uint64, _ti uint8, _gi uint64, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.Withdraw(&_IKmanage.TransactOpts, _ki, _ti, _gi, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9cad4378.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint64 _gi, uint256 money) returns(uint256)
func (_IKmanage *IKmanageTransactorSession) Withdraw(_ki uint64, _ti uint8, _gi uint64, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.Withdraw(&_IKmanage.TransactOpts, _ki, _ti, _gi, money)
}

// IKmanageAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the IKmanage contract.
type IKmanageAddCntIterator struct {
	Event *IKmanageAddCnt // Event containing the contract specifics and raw log

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
func (it *IKmanageAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageAddCnt)
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
		it.Event = new(IKmanageAddCnt)
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
func (it *IKmanageAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageAddCnt represents a AddCnt event raised by the IKmanage contract.
type IKmanageAddCnt struct {
	Gi  uint64
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0xe10954b9cf170a127d0c8a4792a7defcf9d39be421f53862a0432503c8d51f46.
//
// Solidity: event AddCnt(uint64 indexed gi, uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) FilterAddCnt(opts *bind.FilterOpts, gi []uint64, ki []uint64) (*IKmanageAddCntIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanage.contract.FilterLogs(opts, "AddCnt", giRule, kiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageAddCntIterator{contract: _IKmanage.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0xe10954b9cf170a127d0c8a4792a7defcf9d39be421f53862a0432503c8d51f46.
//
// Solidity: event AddCnt(uint64 indexed gi, uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *IKmanageAddCnt, gi []uint64, ki []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanage.contract.WatchLogs(opts, "AddCnt", giRule, kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageAddCnt)
				if err := _IKmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
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

// ParseAddCnt is a log parse operation binding the contract event 0xe10954b9cf170a127d0c8a4792a7defcf9d39be421f53862a0432503c8d51f46.
//
// Solidity: event AddCnt(uint64 indexed gi, uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) ParseAddCnt(log types.Log) (*IKmanageAddCnt, error) {
	event := new(IKmanageAddCnt)
	if err := _IKmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the IKmanage contract.
type IKmanageAddProfitIterator struct {
	Event *IKmanageAddProfit // Event containing the contract specifics and raw log

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
func (it *IKmanageAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageAddProfit)
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
		it.Event = new(IKmanageAddProfit)
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
func (it *IKmanageAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageAddProfit represents a AddProfit event raised by the IKmanage contract.
type IKmanageAddProfit struct {
	Gi    uint64
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0x93434bbdd89b6a302c587aa2dbb0dec681f9b80612b95ab867475055d260819f.
//
// Solidity: event AddProfit(uint64 indexed gi, uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) FilterAddProfit(opts *bind.FilterOpts, gi []uint64, ti []uint8) (*IKmanageAddProfitIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanage.contract.FilterLogs(opts, "AddProfit", giRule, tiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageAddProfitIterator{contract: _IKmanage.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0x93434bbdd89b6a302c587aa2dbb0dec681f9b80612b95ab867475055d260819f.
//
// Solidity: event AddProfit(uint64 indexed gi, uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *IKmanageAddProfit, gi []uint64, ti []uint8) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanage.contract.WatchLogs(opts, "AddProfit", giRule, tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageAddProfit)
				if err := _IKmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
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

// ParseAddProfit is a log parse operation binding the contract event 0x93434bbdd89b6a302c587aa2dbb0dec681f9b80612b95ab867475055d260819f.
//
// Solidity: event AddProfit(uint64 indexed gi, uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) ParseAddProfit(log types.Log) (*IKmanageAddProfit, error) {
	event := new(IKmanageAddProfit)
	if err := _IKmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageGetterABI is the input ABI used to generate the binding from.
const IKmanageGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IKmanageGetterFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageGetterFuncSigs = map[string]string{
	"f57d6e17": "balanceOf(uint64,uint64,uint8)",
	"da41f894": "getK(uint64,uint64)",
	"4cfd1e2d": "getKCnt(uint64)",
	"681e4819": "getPf(uint8,uint64)",
}

// IKmanageGetter is an auto generated Go binding around an Ethereum contract.
type IKmanageGetter struct {
	IKmanageGetterCaller     // Read-only binding to the contract
	IKmanageGetterTransactor // Write-only binding to the contract
	IKmanageGetterFilterer   // Log filterer for contract events
}

// IKmanageGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageGetterSession struct {
	Contract     *IKmanageGetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageGetterCallerSession struct {
	Contract *IKmanageGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKmanageGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageGetterTransactorSession struct {
	Contract     *IKmanageGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKmanageGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageGetterRaw struct {
	Contract *IKmanageGetter // Generic contract binding to access the raw methods on
}

// IKmanageGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageGetterCallerRaw struct {
	Contract *IKmanageGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageGetterTransactorRaw struct {
	Contract *IKmanageGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanageGetter creates a new instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetter(address common.Address, backend bind.ContractBackend) (*IKmanageGetter, error) {
	contract, err := bindIKmanageGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetter{IKmanageGetterCaller: IKmanageGetterCaller{contract: contract}, IKmanageGetterTransactor: IKmanageGetterTransactor{contract: contract}, IKmanageGetterFilterer: IKmanageGetterFilterer{contract: contract}}, nil
}

// NewIKmanageGetterCaller creates a new read-only instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterCaller(address common.Address, caller bind.ContractCaller) (*IKmanageGetterCaller, error) {
	contract, err := bindIKmanageGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterCaller{contract: contract}, nil
}

// NewIKmanageGetterTransactor creates a new write-only instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageGetterTransactor, error) {
	contract, err := bindIKmanageGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterTransactor{contract: contract}, nil
}

// NewIKmanageGetterFilterer creates a new log filterer instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageGetterFilterer, error) {
	contract, err := bindIKmanageGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterFilterer{contract: contract}, nil
}

// bindIKmanageGetter binds a generic wrapper to an already deployed contract.
func bindIKmanageGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageGetter *IKmanageGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageGetter.Contract.IKmanageGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageGetter *IKmanageGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.IKmanageGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageGetter *IKmanageGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.IKmanageGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageGetter *IKmanageGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageGetter *IKmanageGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageGetter *IKmanageGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf57d6e17.
//
// Solidity: function balanceOf(uint64 _gi, uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterCaller) BalanceOf(opts *bind.CallOpts, _gi uint64, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "balanceOf", _gi, _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf57d6e17.
//
// Solidity: function balanceOf(uint64 _gi, uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterSession) BalanceOf(_gi uint64, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanageGetter.Contract.BalanceOf(&_IKmanageGetter.CallOpts, _gi, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf57d6e17.
//
// Solidity: function balanceOf(uint64 _gi, uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterCallerSession) BalanceOf(_gi uint64, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanageGetter.Contract.BalanceOf(&_IKmanageGetter.CallOpts, _gi, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xda41f894.
//
// Solidity: function getK(uint64 _gi, uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterCaller) GetK(opts *bind.CallOpts, _gi uint64, _i uint64) (uint64, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getK", _gi, _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xda41f894.
//
// Solidity: function getK(uint64 _gi, uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterSession) GetK(_gi uint64, _i uint64) (uint64, error) {
	return _IKmanageGetter.Contract.GetK(&_IKmanageGetter.CallOpts, _gi, _i)
}

// GetK is a free data retrieval call binding the contract method 0xda41f894.
//
// Solidity: function getK(uint64 _gi, uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetK(_gi uint64, _i uint64) (uint64, error) {
	return _IKmanageGetter.Contract.GetK(&_IKmanageGetter.CallOpts, _gi, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x4cfd1e2d.
//
// Solidity: function getKCnt(uint64 _gi) view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCaller) GetKCnt(opts *bind.CallOpts, _gi uint64) (uint8, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getKCnt", _gi)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x4cfd1e2d.
//
// Solidity: function getKCnt(uint64 _gi) view returns(uint8)
func (_IKmanageGetter *IKmanageGetterSession) GetKCnt(_gi uint64) (uint8, error) {
	return _IKmanageGetter.Contract.GetKCnt(&_IKmanageGetter.CallOpts, _gi)
}

// GetKCnt is a free data retrieval call binding the contract method 0x4cfd1e2d.
//
// Solidity: function getKCnt(uint64 _gi) view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetKCnt(_gi uint64) (uint8, error) {
	return _IKmanageGetter.Contract.GetKCnt(&_IKmanageGetter.CallOpts, _gi)
}

// GetPf is a free data retrieval call binding the contract method 0x681e4819.
//
// Solidity: function getPf(uint8 _ti, uint64 _gi) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterCaller) GetPf(opts *bind.CallOpts, _ti uint8, _gi uint64) (uint64, *big.Int, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getPf", _ti, _gi)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x681e4819.
//
// Solidity: function getPf(uint8 _ti, uint64 _gi) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterSession) GetPf(_ti uint8, _gi uint64) (uint64, *big.Int, error) {
	return _IKmanageGetter.Contract.GetPf(&_IKmanageGetter.CallOpts, _ti, _gi)
}

// GetPf is a free data retrieval call binding the contract method 0x681e4819.
//
// Solidity: function getPf(uint8 _ti, uint64 _gi) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetPf(_ti uint8, _gi uint64) (uint64, *big.Int, error) {
	return _IKmanageGetter.Contract.GetPf(&_IKmanageGetter.CallOpts, _ti, _gi)
}

// IKmanageSetterABI is the input ABI used to generate the binding from.
const IKmanageSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pe\",\"type\":\"uint64\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKmanageSetterFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageSetterFuncSigs = map[string]string{
	"49d2a1b4": "addCnt(uint64,uint64,uint64)",
	"a666bc94": "addKeeper(uint64,uint64)",
	"2d1664e9": "addProfit(uint8,uint64,uint256)",
	"309cb5ea": "setPeriod(uint64)",
	"9cad4378": "withdraw(uint64,uint8,uint64,uint256)",
}

// IKmanageSetter is an auto generated Go binding around an Ethereum contract.
type IKmanageSetter struct {
	IKmanageSetterCaller     // Read-only binding to the contract
	IKmanageSetterTransactor // Write-only binding to the contract
	IKmanageSetterFilterer   // Log filterer for contract events
}

// IKmanageSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageSetterSession struct {
	Contract     *IKmanageSetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageSetterCallerSession struct {
	Contract *IKmanageSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKmanageSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageSetterTransactorSession struct {
	Contract     *IKmanageSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKmanageSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageSetterRaw struct {
	Contract *IKmanageSetter // Generic contract binding to access the raw methods on
}

// IKmanageSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageSetterCallerRaw struct {
	Contract *IKmanageSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageSetterTransactorRaw struct {
	Contract *IKmanageSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanageSetter creates a new instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetter(address common.Address, backend bind.ContractBackend) (*IKmanageSetter, error) {
	contract, err := bindIKmanageSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetter{IKmanageSetterCaller: IKmanageSetterCaller{contract: contract}, IKmanageSetterTransactor: IKmanageSetterTransactor{contract: contract}, IKmanageSetterFilterer: IKmanageSetterFilterer{contract: contract}}, nil
}

// NewIKmanageSetterCaller creates a new read-only instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterCaller(address common.Address, caller bind.ContractCaller) (*IKmanageSetterCaller, error) {
	contract, err := bindIKmanageSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterCaller{contract: contract}, nil
}

// NewIKmanageSetterTransactor creates a new write-only instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageSetterTransactor, error) {
	contract, err := bindIKmanageSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterTransactor{contract: contract}, nil
}

// NewIKmanageSetterFilterer creates a new log filterer instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageSetterFilterer, error) {
	contract, err := bindIKmanageSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterFilterer{contract: contract}, nil
}

// bindIKmanageSetter binds a generic wrapper to an already deployed contract.
func bindIKmanageSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageSetter *IKmanageSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageSetter.Contract.IKmanageSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageSetter *IKmanageSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.IKmanageSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageSetter *IKmanageSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.IKmanageSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageSetter *IKmanageSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageSetter *IKmanageSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageSetter *IKmanageSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.contract.Transact(opts, method, params...)
}

// AddCnt is a paid mutator transaction binding the contract method 0x49d2a1b4.
//
// Solidity: function addCnt(uint64 _gi, uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddCnt(opts *bind.TransactOpts, _gi uint64, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addCnt", _gi, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x49d2a1b4.
//
// Solidity: function addCnt(uint64 _gi, uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddCnt(_gi uint64, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddCnt(&_IKmanageSetter.TransactOpts, _gi, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x49d2a1b4.
//
// Solidity: function addCnt(uint64 _gi, uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddCnt(_gi uint64, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddCnt(&_IKmanageSetter.TransactOpts, _gi, _ki, _cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0xa666bc94.
//
// Solidity: function addKeeper(uint64 _gi, uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddKeeper(opts *bind.TransactOpts, _gi uint64, _ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addKeeper", _gi, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0xa666bc94.
//
// Solidity: function addKeeper(uint64 _gi, uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddKeeper(_gi uint64, _ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddKeeper(&_IKmanageSetter.TransactOpts, _gi, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0xa666bc94.
//
// Solidity: function addKeeper(uint64 _gi, uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddKeeper(_gi uint64, _ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddKeeper(&_IKmanageSetter.TransactOpts, _gi, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x2d1664e9.
//
// Solidity: function addProfit(uint8 _ti, uint64 _gi, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _gi uint64, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addProfit", _ti, _gi, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x2d1664e9.
//
// Solidity: function addProfit(uint8 _ti, uint64 _gi, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddProfit(_ti uint8, _gi uint64, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddProfit(&_IKmanageSetter.TransactOpts, _ti, _gi, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x2d1664e9.
//
// Solidity: function addProfit(uint8 _ti, uint64 _gi, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddProfit(_ti uint8, _gi uint64, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddProfit(&_IKmanageSetter.TransactOpts, _ti, _gi, _money)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x309cb5ea.
//
// Solidity: function setPeriod(uint64 _pe) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) SetPeriod(opts *bind.TransactOpts, _pe uint64) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "setPeriod", _pe)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x309cb5ea.
//
// Solidity: function setPeriod(uint64 _pe) returns()
func (_IKmanageSetter *IKmanageSetterSession) SetPeriod(_pe uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.SetPeriod(&_IKmanageSetter.TransactOpts, _pe)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x309cb5ea.
//
// Solidity: function setPeriod(uint64 _pe) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) SetPeriod(_pe uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.SetPeriod(&_IKmanageSetter.TransactOpts, _pe)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9cad4378.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint64 _gi, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, _gi uint64, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "withdraw", _ki, _ti, _gi, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9cad4378.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint64 _gi, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterSession) Withdraw(_ki uint64, _ti uint8, _gi uint64, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.Withdraw(&_IKmanageSetter.TransactOpts, _ki, _ti, _gi, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9cad4378.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint64 _gi, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterTransactorSession) Withdraw(_ki uint64, _ti uint8, _gi uint64, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.Withdraw(&_IKmanageSetter.TransactOpts, _ki, _ti, _gi, money)
}

// IKmanageSetterAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the IKmanageSetter contract.
type IKmanageSetterAddCntIterator struct {
	Event *IKmanageSetterAddCnt // Event containing the contract specifics and raw log

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
func (it *IKmanageSetterAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageSetterAddCnt)
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
		it.Event = new(IKmanageSetterAddCnt)
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
func (it *IKmanageSetterAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageSetterAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageSetterAddCnt represents a AddCnt event raised by the IKmanageSetter contract.
type IKmanageSetterAddCnt struct {
	Gi  uint64
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0xe10954b9cf170a127d0c8a4792a7defcf9d39be421f53862a0432503c8d51f46.
//
// Solidity: event AddCnt(uint64 indexed gi, uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) FilterAddCnt(opts *bind.FilterOpts, gi []uint64, ki []uint64) (*IKmanageSetterAddCntIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.FilterLogs(opts, "AddCnt", giRule, kiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterAddCntIterator{contract: _IKmanageSetter.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0xe10954b9cf170a127d0c8a4792a7defcf9d39be421f53862a0432503c8d51f46.
//
// Solidity: event AddCnt(uint64 indexed gi, uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *IKmanageSetterAddCnt, gi []uint64, ki []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.WatchLogs(opts, "AddCnt", giRule, kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageSetterAddCnt)
				if err := _IKmanageSetter.contract.UnpackLog(event, "AddCnt", log); err != nil {
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

// ParseAddCnt is a log parse operation binding the contract event 0xe10954b9cf170a127d0c8a4792a7defcf9d39be421f53862a0432503c8d51f46.
//
// Solidity: event AddCnt(uint64 indexed gi, uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) ParseAddCnt(log types.Log) (*IKmanageSetterAddCnt, error) {
	event := new(IKmanageSetterAddCnt)
	if err := _IKmanageSetter.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageSetterAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the IKmanageSetter contract.
type IKmanageSetterAddProfitIterator struct {
	Event *IKmanageSetterAddProfit // Event containing the contract specifics and raw log

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
func (it *IKmanageSetterAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageSetterAddProfit)
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
		it.Event = new(IKmanageSetterAddProfit)
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
func (it *IKmanageSetterAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageSetterAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageSetterAddProfit represents a AddProfit event raised by the IKmanageSetter contract.
type IKmanageSetterAddProfit struct {
	Gi    uint64
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0x93434bbdd89b6a302c587aa2dbb0dec681f9b80612b95ab867475055d260819f.
//
// Solidity: event AddProfit(uint64 indexed gi, uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) FilterAddProfit(opts *bind.FilterOpts, gi []uint64, ti []uint8) (*IKmanageSetterAddProfitIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.FilterLogs(opts, "AddProfit", giRule, tiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterAddProfitIterator{contract: _IKmanageSetter.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0x93434bbdd89b6a302c587aa2dbb0dec681f9b80612b95ab867475055d260819f.
//
// Solidity: event AddProfit(uint64 indexed gi, uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *IKmanageSetterAddProfit, gi []uint64, ti []uint8) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.WatchLogs(opts, "AddProfit", giRule, tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageSetterAddProfit)
				if err := _IKmanageSetter.contract.UnpackLog(event, "AddProfit", log); err != nil {
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

// ParseAddProfit is a log parse operation binding the contract event 0x93434bbdd89b6a302c587aa2dbb0dec681f9b80612b95ab867475055d260819f.
//
// Solidity: event AddProfit(uint64 indexed gi, uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) ParseAddProfit(log types.Log) (*IKmanageSetterAddProfit, error) {
	event := new(IKmanageSetterAddProfit)
	if err := _IKmanageSetter.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOwnerABI is the input ABI used to generate the binding from.
const IOwnerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"}]"

// IOwner is an auto generated Go binding around an Ethereum contract.
type IOwner struct {
	IOwnerCaller     // Read-only binding to the contract
	IOwnerTransactor // Write-only binding to the contract
	IOwnerFilterer   // Log filterer for contract events
}

// IOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnerSession struct {
	Contract     *IOwner           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnerCallerSession struct {
	Contract *IOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnerTransactorSession struct {
	Contract     *IOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnerRaw struct {
	Contract *IOwner // Generic contract binding to access the raw methods on
}

// IOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnerCallerRaw struct {
	Contract *IOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// IOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnerTransactorRaw struct {
	Contract *IOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwner creates a new instance of IOwner, bound to a specific deployed contract.
func NewIOwner(address common.Address, backend bind.ContractBackend) (*IOwner, error) {
	contract, err := bindIOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwner{IOwnerCaller: IOwnerCaller{contract: contract}, IOwnerTransactor: IOwnerTransactor{contract: contract}, IOwnerFilterer: IOwnerFilterer{contract: contract}}, nil
}

// NewIOwnerCaller creates a new read-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerCaller(address common.Address, caller bind.ContractCaller) (*IOwnerCaller, error) {
	contract, err := bindIOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerCaller{contract: contract}, nil
}

// NewIOwnerTransactor creates a new write-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IOwnerTransactor, error) {
	contract, err := bindIOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerTransactor{contract: contract}, nil
}

// NewIOwnerFilterer creates a new log filterer instance of IOwner, bound to a specific deployed contract.
func NewIOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IOwnerFilterer, error) {
	contract, err := bindIOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnerFilterer{contract: contract}, nil
}

// bindIOwner binds a generic wrapper to an already deployed contract.
func bindIOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.IOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transact(opts, method, params...)
}

// IOwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the IOwner contract.
type IOwnerAddOwnerIterator struct {
	Event *IOwnerAddOwner // Event containing the contract specifics and raw log

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
func (it *IOwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOwnerAddOwner)
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
		it.Event = new(IOwnerAddOwner)
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
func (it *IOwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOwnerAddOwner represents a AddOwner event raised by the IOwner contract.
type IOwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*IOwnerAddOwnerIterator, error) {

	logs, sub, err := _IOwner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &IOwnerAddOwnerIterator{contract: _IOwner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *IOwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _IOwner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOwnerAddOwner)
				if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) ParseAddOwner(log types.Log) (*IOwnerAddOwner, error) {
	event := new(IOwnerAddOwner)
	if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPledgeABI is the input ABI used to generate the binding from.
const IPledgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPledgeFuncSigs maps the 4-byte function signature to its string representation.
var IPledgeFuncSigs = map[string]string{
	"724a84a8": "addT(uint8)",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"eb395fde": "getPledge(uint8)",
	"d23f7482": "pledge(uint64,uint256)",
	"14f732de": "rewards(uint64,uint8)",
	"c21a43e4": "totalPledge()",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IPledge is an auto generated Go binding around an Ethereum contract.
type IPledge struct {
	IPledgeCaller     // Read-only binding to the contract
	IPledgeTransactor // Write-only binding to the contract
	IPledgeFilterer   // Log filterer for contract events
}

// IPledgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgeSession struct {
	Contract     *IPledge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgeCallerSession struct {
	Contract *IPledgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IPledgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgeTransactorSession struct {
	Contract     *IPledgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IPledgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgeRaw struct {
	Contract *IPledge // Generic contract binding to access the raw methods on
}

// IPledgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgeCallerRaw struct {
	Contract *IPledgeCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgeTransactorRaw struct {
	Contract *IPledgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledge creates a new instance of IPledge, bound to a specific deployed contract.
func NewIPledge(address common.Address, backend bind.ContractBackend) (*IPledge, error) {
	contract, err := bindIPledge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledge{IPledgeCaller: IPledgeCaller{contract: contract}, IPledgeTransactor: IPledgeTransactor{contract: contract}, IPledgeFilterer: IPledgeFilterer{contract: contract}}, nil
}

// NewIPledgeCaller creates a new read-only instance of IPledge, bound to a specific deployed contract.
func NewIPledgeCaller(address common.Address, caller bind.ContractCaller) (*IPledgeCaller, error) {
	contract, err := bindIPledge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeCaller{contract: contract}, nil
}

// NewIPledgeTransactor creates a new write-only instance of IPledge, bound to a specific deployed contract.
func NewIPledgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgeTransactor, error) {
	contract, err := bindIPledge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeTransactor{contract: contract}, nil
}

// NewIPledgeFilterer creates a new log filterer instance of IPledge, bound to a specific deployed contract.
func NewIPledgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgeFilterer, error) {
	contract, err := bindIPledge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgeFilterer{contract: contract}, nil
}

// bindIPledge binds a generic wrapper to an already deployed contract.
func bindIPledge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledge *IPledgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledge.Contract.IPledgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledge *IPledgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledge.Contract.IPledgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledge *IPledgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledge.Contract.IPledgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledge *IPledgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledge *IPledgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledge *IPledgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledge.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledge.Contract.BalanceOf(&_IPledge.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledge.Contract.BalanceOf(&_IPledge.CallOpts, _i, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCaller) GetPledge(opts *bind.CallOpts, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "getPledge", _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledge.Contract.GetPledge(&_IPledge.CallOpts, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCallerSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledge.Contract.GetPledge(&_IPledge.CallOpts, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledge *IPledgeCaller) Rewards(opts *bind.CallOpts, _index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "rewards", _index, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledge *IPledgeSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledge.Contract.Rewards(&_IPledge.CallOpts, _index, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledge *IPledgeCallerSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledge.Contract.Rewards(&_IPledge.CallOpts, _index, _ti)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledge *IPledgeCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "totalPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledge *IPledgeSession) TotalPledge() (*big.Int, error) {
	return _IPledge.Contract.TotalPledge(&_IPledge.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledge *IPledgeCallerSession) TotalPledge() (*big.Int, error) {
	return _IPledge.Contract.TotalPledge(&_IPledge.CallOpts)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledge *IPledgeTransactor) AddT(opts *bind.TransactOpts, _ti uint8) (*types.Transaction, error) {
	return _IPledge.contract.Transact(opts, "addT", _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledge *IPledgeSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledge.Contract.AddT(&_IPledge.TransactOpts, _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledge *IPledgeTransactorSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledge.Contract.AddT(&_IPledge.TransactOpts, _ti)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledge *IPledgeTransactor) Pledge(opts *bind.TransactOpts, _i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledge.contract.Transact(opts, "pledge", _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledge *IPledgeSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Pledge(&_IPledge.TransactOpts, _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledge *IPledgeTransactorSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Pledge(&_IPledge.TransactOpts, _i, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledge *IPledgeTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledge.contract.Transact(opts, "withdraw", _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledge *IPledgeSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Withdraw(&_IPledge.TransactOpts, _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledge *IPledgeTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Withdraw(&_IPledge.TransactOpts, _i, _ti, money, lock)
}

// IPledgeAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IPledge contract.
type IPledgeAddTIterator struct {
	Event *IPledgeAddT // Event containing the contract specifics and raw log

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
func (it *IPledgeAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPledgeAddT)
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
		it.Event = new(IPledgeAddT)
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
func (it *IPledgeAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPledgeAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPledgeAddT represents a AddT event raised by the IPledge contract.
type IPledgeAddT struct {
	Bal    *big.Int
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledge *IPledgeFilterer) FilterAddT(opts *bind.FilterOpts) (*IPledgeAddTIterator, error) {

	logs, sub, err := _IPledge.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &IPledgeAddTIterator{contract: _IPledge.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledge *IPledgeFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *IPledgeAddT) (event.Subscription, error) {

	logs, sub, err := _IPledge.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPledgeAddT)
				if err := _IPledge.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledge *IPledgeFilterer) ParseAddT(log types.Log) (*IPledgeAddT, error) {
	event := new(IPledgeAddT)
	if err := _IPledge.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPledgeGetterABI is the input ABI used to generate the binding from.
const IPledgeGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IPledgeGetterFuncSigs maps the 4-byte function signature to its string representation.
var IPledgeGetterFuncSigs = map[string]string{
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"eb395fde": "getPledge(uint8)",
	"14f732de": "rewards(uint64,uint8)",
	"c21a43e4": "totalPledge()",
}

// IPledgeGetter is an auto generated Go binding around an Ethereum contract.
type IPledgeGetter struct {
	IPledgeGetterCaller     // Read-only binding to the contract
	IPledgeGetterTransactor // Write-only binding to the contract
	IPledgeGetterFilterer   // Log filterer for contract events
}

// IPledgeGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgeGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgeGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgeGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgeGetterSession struct {
	Contract     *IPledgeGetter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgeGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgeGetterCallerSession struct {
	Contract *IPledgeGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPledgeGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgeGetterTransactorSession struct {
	Contract     *IPledgeGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPledgeGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgeGetterRaw struct {
	Contract *IPledgeGetter // Generic contract binding to access the raw methods on
}

// IPledgeGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgeGetterCallerRaw struct {
	Contract *IPledgeGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgeGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgeGetterTransactorRaw struct {
	Contract *IPledgeGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledgeGetter creates a new instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetter(address common.Address, backend bind.ContractBackend) (*IPledgeGetter, error) {
	contract, err := bindIPledgeGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetter{IPledgeGetterCaller: IPledgeGetterCaller{contract: contract}, IPledgeGetterTransactor: IPledgeGetterTransactor{contract: contract}, IPledgeGetterFilterer: IPledgeGetterFilterer{contract: contract}}, nil
}

// NewIPledgeGetterCaller creates a new read-only instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetterCaller(address common.Address, caller bind.ContractCaller) (*IPledgeGetterCaller, error) {
	contract, err := bindIPledgeGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetterCaller{contract: contract}, nil
}

// NewIPledgeGetterTransactor creates a new write-only instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgeGetterTransactor, error) {
	contract, err := bindIPledgeGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetterTransactor{contract: contract}, nil
}

// NewIPledgeGetterFilterer creates a new log filterer instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgeGetterFilterer, error) {
	contract, err := bindIPledgeGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetterFilterer{contract: contract}, nil
}

// bindIPledgeGetter binds a generic wrapper to an already deployed contract.
func bindIPledgeGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgeGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeGetter *IPledgeGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeGetter.Contract.IPledgeGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeGetter *IPledgeGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.IPledgeGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeGetter *IPledgeGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.IPledgeGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeGetter *IPledgeGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeGetter *IPledgeGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeGetter *IPledgeGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.BalanceOf(&_IPledgeGetter.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.BalanceOf(&_IPledgeGetter.CallOpts, _i, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCaller) GetPledge(opts *bind.CallOpts, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "getPledge", _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.GetPledge(&_IPledgeGetter.CallOpts, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.GetPledge(&_IPledgeGetter.CallOpts, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledgeGetter *IPledgeGetterCaller) Rewards(opts *bind.CallOpts, _index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "rewards", _index, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledgeGetter *IPledgeGetterSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledgeGetter.Contract.Rewards(&_IPledgeGetter.CallOpts, _index, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledgeGetter.Contract.Rewards(&_IPledgeGetter.CallOpts, _index, _ti)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "totalPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledgeGetter *IPledgeGetterSession) TotalPledge() (*big.Int, error) {
	return _IPledgeGetter.Contract.TotalPledge(&_IPledgeGetter.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) TotalPledge() (*big.Int, error) {
	return _IPledgeGetter.Contract.TotalPledge(&_IPledgeGetter.CallOpts)
}

// IPledgeSetterABI is the input ABI used to generate the binding from.
const IPledgeSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPledgeSetterFuncSigs maps the 4-byte function signature to its string representation.
var IPledgeSetterFuncSigs = map[string]string{
	"724a84a8": "addT(uint8)",
	"d23f7482": "pledge(uint64,uint256)",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IPledgeSetter is an auto generated Go binding around an Ethereum contract.
type IPledgeSetter struct {
	IPledgeSetterCaller     // Read-only binding to the contract
	IPledgeSetterTransactor // Write-only binding to the contract
	IPledgeSetterFilterer   // Log filterer for contract events
}

// IPledgeSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgeSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgeSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgeSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgeSetterSession struct {
	Contract     *IPledgeSetter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgeSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgeSetterCallerSession struct {
	Contract *IPledgeSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPledgeSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgeSetterTransactorSession struct {
	Contract     *IPledgeSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPledgeSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgeSetterRaw struct {
	Contract *IPledgeSetter // Generic contract binding to access the raw methods on
}

// IPledgeSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgeSetterCallerRaw struct {
	Contract *IPledgeSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgeSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgeSetterTransactorRaw struct {
	Contract *IPledgeSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledgeSetter creates a new instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetter(address common.Address, backend bind.ContractBackend) (*IPledgeSetter, error) {
	contract, err := bindIPledgeSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetter{IPledgeSetterCaller: IPledgeSetterCaller{contract: contract}, IPledgeSetterTransactor: IPledgeSetterTransactor{contract: contract}, IPledgeSetterFilterer: IPledgeSetterFilterer{contract: contract}}, nil
}

// NewIPledgeSetterCaller creates a new read-only instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetterCaller(address common.Address, caller bind.ContractCaller) (*IPledgeSetterCaller, error) {
	contract, err := bindIPledgeSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterCaller{contract: contract}, nil
}

// NewIPledgeSetterTransactor creates a new write-only instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgeSetterTransactor, error) {
	contract, err := bindIPledgeSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterTransactor{contract: contract}, nil
}

// NewIPledgeSetterFilterer creates a new log filterer instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgeSetterFilterer, error) {
	contract, err := bindIPledgeSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterFilterer{contract: contract}, nil
}

// bindIPledgeSetter binds a generic wrapper to an already deployed contract.
func bindIPledgeSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgeSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeSetter *IPledgeSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeSetter.Contract.IPledgeSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeSetter *IPledgeSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.IPledgeSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeSetter *IPledgeSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.IPledgeSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeSetter *IPledgeSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeSetter *IPledgeSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeSetter *IPledgeSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.contract.Transact(opts, method, params...)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledgeSetter *IPledgeSetterTransactor) AddT(opts *bind.TransactOpts, _ti uint8) (*types.Transaction, error) {
	return _IPledgeSetter.contract.Transact(opts, "addT", _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledgeSetter *IPledgeSetterSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.AddT(&_IPledgeSetter.TransactOpts, _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledgeSetter *IPledgeSetterTransactorSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.AddT(&_IPledgeSetter.TransactOpts, _ti)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledgeSetter *IPledgeSetterTransactor) Pledge(opts *bind.TransactOpts, _i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.contract.Transact(opts, "pledge", _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledgeSetter *IPledgeSetterSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Pledge(&_IPledgeSetter.TransactOpts, _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledgeSetter *IPledgeSetterTransactorSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Pledge(&_IPledgeSetter.TransactOpts, _i, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledgeSetter *IPledgeSetterTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.contract.Transact(opts, "withdraw", _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledgeSetter *IPledgeSetterSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Withdraw(&_IPledgeSetter.TransactOpts, _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledgeSetter *IPledgeSetterTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Withdraw(&_IPledgeSetter.TransactOpts, _i, _ti, money, lock)
}

// IPledgeSetterAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IPledgeSetter contract.
type IPledgeSetterAddTIterator struct {
	Event *IPledgeSetterAddT // Event containing the contract specifics and raw log

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
func (it *IPledgeSetterAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPledgeSetterAddT)
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
		it.Event = new(IPledgeSetterAddT)
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
func (it *IPledgeSetterAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPledgeSetterAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPledgeSetterAddT represents a AddT event raised by the IPledgeSetter contract.
type IPledgeSetterAddT struct {
	Bal    *big.Int
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledgeSetter *IPledgeSetterFilterer) FilterAddT(opts *bind.FilterOpts) (*IPledgeSetterAddTIterator, error) {

	logs, sub, err := _IPledgeSetter.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterAddTIterator{contract: _IPledgeSetter.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledgeSetter *IPledgeSetterFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *IPledgeSetterAddT) (event.Subscription, error) {

	logs, sub, err := _IPledgeSetter.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPledgeSetterAddT)
				if err := _IPledgeSetter.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledgeSetter *IPledgeSetterFilterer) ParseAddT(log types.Log) (*IPledgeSetterAddT, error) {
	event := new(IPledgeSetterAddT)
	if err := _IPledgeSetter.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolABI is the input ABI used to generate the binding from.
const IPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Inflow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Outflow\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"inflow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"outflow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IPoolFuncSigs maps the 4-byte function signature to its string representation.
var IPoolFuncSigs = map[string]string{
	"368007fe": "inflow(address,address,uint256)",
	"530345bb": "outflow(address,address,uint256)",
}

// IPool is an auto generated Go binding around an Ethereum contract.
type IPool struct {
	IPoolCaller     // Read-only binding to the contract
	IPoolTransactor // Write-only binding to the contract
	IPoolFilterer   // Log filterer for contract events
}

// IPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolSession struct {
	Contract     *IPool            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolCallerSession struct {
	Contract *IPoolCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolTransactorSession struct {
	Contract     *IPoolTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolRaw struct {
	Contract *IPool // Generic contract binding to access the raw methods on
}

// IPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolCallerRaw struct {
	Contract *IPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolTransactorRaw struct {
	Contract *IPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPool creates a new instance of IPool, bound to a specific deployed contract.
func NewIPool(address common.Address, backend bind.ContractBackend) (*IPool, error) {
	contract, err := bindIPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPool{IPoolCaller: IPoolCaller{contract: contract}, IPoolTransactor: IPoolTransactor{contract: contract}, IPoolFilterer: IPoolFilterer{contract: contract}}, nil
}

// NewIPoolCaller creates a new read-only instance of IPool, bound to a specific deployed contract.
func NewIPoolCaller(address common.Address, caller bind.ContractCaller) (*IPoolCaller, error) {
	contract, err := bindIPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolCaller{contract: contract}, nil
}

// NewIPoolTransactor creates a new write-only instance of IPool, bound to a specific deployed contract.
func NewIPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolTransactor, error) {
	contract, err := bindIPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolTransactor{contract: contract}, nil
}

// NewIPoolFilterer creates a new log filterer instance of IPool, bound to a specific deployed contract.
func NewIPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolFilterer, error) {
	contract, err := bindIPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolFilterer{contract: contract}, nil
}

// bindIPool binds a generic wrapper to an already deployed contract.
func bindIPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.IPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transact(opts, method, params...)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolTransactor) Inflow(opts *bind.TransactOpts, tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "inflow", tAddr, from, money)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolSession) Inflow(tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Inflow(&_IPool.TransactOpts, tAddr, from, money)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolTransactorSession) Inflow(tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Inflow(&_IPool.TransactOpts, tAddr, from, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolTransactor) Outflow(opts *bind.TransactOpts, tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "outflow", tAddr, to, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolSession) Outflow(tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Outflow(&_IPool.TransactOpts, tAddr, to, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolTransactorSession) Outflow(tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Outflow(&_IPool.TransactOpts, tAddr, to, money)
}

// IPoolInflowIterator is returned from FilterInflow and is used to iterate over the raw logs and unpacked data for Inflow events raised by the IPool contract.
type IPoolInflowIterator struct {
	Event *IPoolInflow // Event containing the contract specifics and raw log

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
func (it *IPoolInflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolInflow)
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
		it.Event = new(IPoolInflow)
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
func (it *IPoolInflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolInflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolInflow represents a Inflow event raised by the IPool contract.
type IPoolInflow struct {
	From  common.Address
	TAddr common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInflow is a free log retrieval operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) FilterInflow(opts *bind.FilterOpts, from []common.Address) (*IPoolInflowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Inflow", fromRule)
	if err != nil {
		return nil, err
	}
	return &IPoolInflowIterator{contract: _IPool.contract, event: "Inflow", logs: logs, sub: sub}, nil
}

// WatchInflow is a free log subscription operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) WatchInflow(opts *bind.WatchOpts, sink chan<- *IPoolInflow, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Inflow", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolInflow)
				if err := _IPool.contract.UnpackLog(event, "Inflow", log); err != nil {
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

// ParseInflow is a log parse operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) ParseInflow(log types.Log) (*IPoolInflow, error) {
	event := new(IPoolInflow)
	if err := _IPool.contract.UnpackLog(event, "Inflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolOutflowIterator is returned from FilterOutflow and is used to iterate over the raw logs and unpacked data for Outflow events raised by the IPool contract.
type IPoolOutflowIterator struct {
	Event *IPoolOutflow // Event containing the contract specifics and raw log

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
func (it *IPoolOutflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolOutflow)
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
		it.Event = new(IPoolOutflow)
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
func (it *IPoolOutflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolOutflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolOutflow represents a Outflow event raised by the IPool contract.
type IPoolOutflow struct {
	To    common.Address
	TAddr common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOutflow is a free log retrieval operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) FilterOutflow(opts *bind.FilterOpts, to []common.Address) (*IPoolOutflowIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Outflow", toRule)
	if err != nil {
		return nil, err
	}
	return &IPoolOutflowIterator{contract: _IPool.contract, event: "Outflow", logs: logs, sub: sub}, nil
}

// WatchOutflow is a free log subscription operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) WatchOutflow(opts *bind.WatchOpts, sink chan<- *IPoolOutflow, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Outflow", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolOutflow)
				if err := _IPool.contract.UnpackLog(event, "Outflow", log); err != nil {
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

// ParseOutflow is a log parse operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) ParseOutflow(log types.Log) (*IPoolOutflow, error) {
	event := new(IPoolOutflow)
	if err := _IPool.contract.UnpackLog(event, "Outflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleGetterABI is the input ABI used to generate the binding from.
const IRoleGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getAcc\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getRInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"verifyKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"desc\",\"type\":\"bytes\"}],\"internalType\":\"structRoleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRoleGetterFuncSigs maps the 4-byte function signature to its string representation.
var IRoleGetterFuncSigs = map[string]string{
	"7738515f": "checkIG(uint64,uint8)",
	"7264a551": "getACnt()",
	"caca4a06": "getAcc(address)",
	"9332aa6e": "getAddr(uint64)",
	"441abace": "getRInfo(address)",
}

// IRoleGetter is an auto generated Go binding around an Ethereum contract.
type IRoleGetter struct {
	IRoleGetterCaller     // Read-only binding to the contract
	IRoleGetterTransactor // Write-only binding to the contract
	IRoleGetterFilterer   // Log filterer for contract events
}

// IRoleGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoleGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoleGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoleGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoleGetterSession struct {
	Contract     *IRoleGetter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoleGetterCallerSession struct {
	Contract *IRoleGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IRoleGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoleGetterTransactorSession struct {
	Contract     *IRoleGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IRoleGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoleGetterRaw struct {
	Contract *IRoleGetter // Generic contract binding to access the raw methods on
}

// IRoleGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoleGetterCallerRaw struct {
	Contract *IRoleGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IRoleGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoleGetterTransactorRaw struct {
	Contract *IRoleGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRoleGetter creates a new instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetter(address common.Address, backend bind.ContractBackend) (*IRoleGetter, error) {
	contract, err := bindIRoleGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRoleGetter{IRoleGetterCaller: IRoleGetterCaller{contract: contract}, IRoleGetterTransactor: IRoleGetterTransactor{contract: contract}, IRoleGetterFilterer: IRoleGetterFilterer{contract: contract}}, nil
}

// NewIRoleGetterCaller creates a new read-only instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetterCaller(address common.Address, caller bind.ContractCaller) (*IRoleGetterCaller, error) {
	contract, err := bindIRoleGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleGetterCaller{contract: contract}, nil
}

// NewIRoleGetterTransactor creates a new write-only instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoleGetterTransactor, error) {
	contract, err := bindIRoleGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleGetterTransactor{contract: contract}, nil
}

// NewIRoleGetterFilterer creates a new log filterer instance of IRoleGetter, bound to a specific deployed contract.
func NewIRoleGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoleGetterFilterer, error) {
	contract, err := bindIRoleGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoleGetterFilterer{contract: contract}, nil
}

// bindIRoleGetter binds a generic wrapper to an already deployed contract.
func bindIRoleGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRoleGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleGetter *IRoleGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleGetter.Contract.IRoleGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleGetter *IRoleGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleGetter.Contract.IRoleGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleGetter *IRoleGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleGetter.Contract.IRoleGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleGetter *IRoleGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleGetter *IRoleGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleGetter *IRoleGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleGetter.Contract.contract.Transact(opts, method, params...)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8, uint8)
func (_IRoleGetter *IRoleGetterCaller) CheckIG(opts *bind.CallOpts, _i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, uint8, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "checkIG", _i, _rType)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(uint64), *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8, uint8)
func (_IRoleGetter *IRoleGetterSession) CheckIG(_i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, uint8, error) {
	return _IRoleGetter.Contract.CheckIG(&_IRoleGetter.CallOpts, _i, _rType)
}

// CheckIG is a free data retrieval call binding the contract method 0x7738515f.
//
// Solidity: function checkIG(uint64 _i, uint8 _rType) view returns(address, address, uint64, uint8, uint8)
func (_IRoleGetter *IRoleGetterCallerSession) CheckIG(_i uint64, _rType uint8) (common.Address, common.Address, uint64, uint8, uint8, error) {
	return _IRoleGetter.Contract.CheckIG(&_IRoleGetter.CallOpts, _i, _rType)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRoleGetter *IRoleGetterCaller) GetACnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getACnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRoleGetter *IRoleGetterSession) GetACnt() (uint64, error) {
	return _IRoleGetter.Contract.GetACnt(&_IRoleGetter.CallOpts)
}

// GetACnt is a free data retrieval call binding the contract method 0x7264a551.
//
// Solidity: function getACnt() view returns(uint64)
func (_IRoleGetter *IRoleGetterCallerSession) GetACnt() (uint64, error) {
	return _IRoleGetter.Contract.GetACnt(&_IRoleGetter.CallOpts)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRoleGetter *IRoleGetterCaller) GetAcc(opts *bind.CallOpts, _a common.Address) (uint64, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getAcc", _a)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRoleGetter *IRoleGetterSession) GetAcc(_a common.Address) (uint64, error) {
	return _IRoleGetter.Contract.GetAcc(&_IRoleGetter.CallOpts, _a)
}

// GetAcc is a free data retrieval call binding the contract method 0xcaca4a06.
//
// Solidity: function getAcc(address _a) view returns(uint64)
func (_IRoleGetter *IRoleGetterCallerSession) GetAcc(_a common.Address) (uint64, error) {
	return _IRoleGetter.Contract.GetAcc(&_IRoleGetter.CallOpts, _a)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRoleGetter *IRoleGetterCaller) GetAddr(opts *bind.CallOpts, _i uint64) (common.Address, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getAddr", _i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRoleGetter *IRoleGetterSession) GetAddr(_i uint64) (common.Address, error) {
	return _IRoleGetter.Contract.GetAddr(&_IRoleGetter.CallOpts, _i)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 _i) view returns(address)
func (_IRoleGetter *IRoleGetterCallerSession) GetAddr(_i uint64) (common.Address, error) {
	return _IRoleGetter.Contract.GetAddr(&_IRoleGetter.CallOpts, _i)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((uint8,uint8,uint64,uint64,address,address,bytes,bytes))
func (_IRoleGetter *IRoleGetterCaller) GetRInfo(opts *bind.CallOpts, _a common.Address) (RoleOut, error) {
	var out []interface{}
	err := _IRoleGetter.contract.Call(opts, &out, "getRInfo", _a)

	if err != nil {
		return *new(RoleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(RoleOut)).(*RoleOut)

	return out0, err

}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((uint8,uint8,uint64,uint64,address,address,bytes,bytes))
func (_IRoleGetter *IRoleGetterSession) GetRInfo(_a common.Address) (RoleOut, error) {
	return _IRoleGetter.Contract.GetRInfo(&_IRoleGetter.CallOpts, _a)
}

// GetRInfo is a free data retrieval call binding the contract method 0x441abace.
//
// Solidity: function getRInfo(address _a) view returns((uint8,uint8,uint64,uint64,address,address,bytes,bytes))
func (_IRoleGetter *IRoleGetterCallerSession) GetRInfo(_a common.Address) (RoleOut, error) {
	return _IRoleGetter.Contract.GetRInfo(&_IRoleGetter.CallOpts, _a)
}

// IRoleSetterABI is the input ABI used to generate the binding from.
const IRoleSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AlterPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ConfirmPayee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"QuitRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReAcc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"rType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"ReRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"SetG\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"quitRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"setDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"setG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"}],\"name\":\"setS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IRoleSetterFuncSigs maps the 4-byte function signature to its string representation.
var IRoleSetterFuncSigs = map[string]string{
	"2915824a": "alterPayee(uint64,address)",
	"5bd04c3f": "confirmPayee(uint64,address)",
	"adad6d3a": "quitRole(uint64)",
	"effcafce": "reAcc(address)",
	"b9f6a8ca": "reRole(uint64,uint8,bytes)",
	"722cc4a4": "setDesc(address,bytes)",
	"616bfd1f": "setG(uint64,uint64)",
	"b6b6dc0a": "setS(uint64,uint8)",
}

// IRoleSetter is an auto generated Go binding around an Ethereum contract.
type IRoleSetter struct {
	IRoleSetterCaller     // Read-only binding to the contract
	IRoleSetterTransactor // Write-only binding to the contract
	IRoleSetterFilterer   // Log filterer for contract events
}

// IRoleSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoleSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoleSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoleSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoleSetterSession struct {
	Contract     *IRoleSetter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoleSetterCallerSession struct {
	Contract *IRoleSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IRoleSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoleSetterTransactorSession struct {
	Contract     *IRoleSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IRoleSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoleSetterRaw struct {
	Contract *IRoleSetter // Generic contract binding to access the raw methods on
}

// IRoleSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoleSetterCallerRaw struct {
	Contract *IRoleSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IRoleSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoleSetterTransactorRaw struct {
	Contract *IRoleSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRoleSetter creates a new instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetter(address common.Address, backend bind.ContractBackend) (*IRoleSetter, error) {
	contract, err := bindIRoleSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRoleSetter{IRoleSetterCaller: IRoleSetterCaller{contract: contract}, IRoleSetterTransactor: IRoleSetterTransactor{contract: contract}, IRoleSetterFilterer: IRoleSetterFilterer{contract: contract}}, nil
}

// NewIRoleSetterCaller creates a new read-only instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetterCaller(address common.Address, caller bind.ContractCaller) (*IRoleSetterCaller, error) {
	contract, err := bindIRoleSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterCaller{contract: contract}, nil
}

// NewIRoleSetterTransactor creates a new write-only instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoleSetterTransactor, error) {
	contract, err := bindIRoleSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterTransactor{contract: contract}, nil
}

// NewIRoleSetterFilterer creates a new log filterer instance of IRoleSetter, bound to a specific deployed contract.
func NewIRoleSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoleSetterFilterer, error) {
	contract, err := bindIRoleSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterFilterer{contract: contract}, nil
}

// bindIRoleSetter binds a generic wrapper to an already deployed contract.
func bindIRoleSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRoleSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleSetter *IRoleSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleSetter.Contract.IRoleSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleSetter *IRoleSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleSetter.Contract.IRoleSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleSetter *IRoleSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleSetter.Contract.IRoleSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoleSetter *IRoleSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoleSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoleSetter *IRoleSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoleSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoleSetter *IRoleSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoleSetter.Contract.contract.Transact(opts, method, params...)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactor) AlterPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "alterPayee", _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.AlterPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) AlterPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.AlterPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactor) ConfirmPayee(opts *bind.TransactOpts, _index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "confirmPayee", _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ConfirmPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x5bd04c3f.
//
// Solidity: function confirmPayee(uint64 _index, address _p) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) ConfirmPayee(_index uint64, _p common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ConfirmPayee(&_IRoleSetter.TransactOpts, _index, _p)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_IRoleSetter *IRoleSetterTransactor) QuitRole(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "quitRole", _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_IRoleSetter *IRoleSetterSession) QuitRole(_i uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.QuitRole(&_IRoleSetter.TransactOpts, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) QuitRole(_i uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.QuitRole(&_IRoleSetter.TransactOpts, _i)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRoleSetter *IRoleSetterTransactor) ReAcc(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "reAcc", _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRoleSetter *IRoleSetterSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReAcc(&_IRoleSetter.TransactOpts, _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReAcc(&_IRoleSetter.TransactOpts, _a)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRoleSetter *IRoleSetterTransactor) ReRole(opts *bind.TransactOpts, _i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "reRole", _i, _rtype, extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRoleSetter *IRoleSetterSession) ReRole(_i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReRole(&_IRoleSetter.TransactOpts, _i, _rtype, extra)
}

// ReRole is a paid mutator transaction binding the contract method 0xb9f6a8ca.
//
// Solidity: function reRole(uint64 _i, uint8 _rtype, bytes extra) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) ReRole(_i uint64, _rtype uint8, extra []byte) (*types.Transaction, error) {
	return _IRoleSetter.Contract.ReRole(&_IRoleSetter.TransactOpts, _i, _rtype, extra)
}

// SetDesc is a paid mutator transaction binding the contract method 0x722cc4a4.
//
// Solidity: function setDesc(address a, bytes _desc) returns()
func (_IRoleSetter *IRoleSetterTransactor) SetDesc(opts *bind.TransactOpts, a common.Address, _desc []byte) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "setDesc", a, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x722cc4a4.
//
// Solidity: function setDesc(address a, bytes _desc) returns()
func (_IRoleSetter *IRoleSetterSession) SetDesc(a common.Address, _desc []byte) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetDesc(&_IRoleSetter.TransactOpts, a, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x722cc4a4.
//
// Solidity: function setDesc(address a, bytes _desc) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) SetDesc(a common.Address, _desc []byte) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetDesc(&_IRoleSetter.TransactOpts, a, _desc)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRoleSetter *IRoleSetterTransactor) SetG(opts *bind.TransactOpts, _i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "setG", _i, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRoleSetter *IRoleSetterSession) SetG(_i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetG(&_IRoleSetter.TransactOpts, _i, _gi)
}

// SetG is a paid mutator transaction binding the contract method 0x616bfd1f.
//
// Solidity: function setG(uint64 _i, uint64 _gi) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) SetG(_i uint64, _gi uint64) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetG(&_IRoleSetter.TransactOpts, _i, _gi)
}

// SetS is a paid mutator transaction binding the contract method 0xb6b6dc0a.
//
// Solidity: function setS(uint64 _i, uint8 _s) returns()
func (_IRoleSetter *IRoleSetterTransactor) SetS(opts *bind.TransactOpts, _i uint64, _s uint8) (*types.Transaction, error) {
	return _IRoleSetter.contract.Transact(opts, "setS", _i, _s)
}

// SetS is a paid mutator transaction binding the contract method 0xb6b6dc0a.
//
// Solidity: function setS(uint64 _i, uint8 _s) returns()
func (_IRoleSetter *IRoleSetterSession) SetS(_i uint64, _s uint8) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetS(&_IRoleSetter.TransactOpts, _i, _s)
}

// SetS is a paid mutator transaction binding the contract method 0xb6b6dc0a.
//
// Solidity: function setS(uint64 _i, uint8 _s) returns()
func (_IRoleSetter *IRoleSetterTransactorSession) SetS(_i uint64, _s uint8) (*types.Transaction, error) {
	return _IRoleSetter.Contract.SetS(&_IRoleSetter.TransactOpts, _i, _s)
}

// IRoleSetterAlterPayeeIterator is returned from FilterAlterPayee and is used to iterate over the raw logs and unpacked data for AlterPayee events raised by the IRoleSetter contract.
type IRoleSetterAlterPayeeIterator struct {
	Event *IRoleSetterAlterPayee // Event containing the contract specifics and raw log

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
func (it *IRoleSetterAlterPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterAlterPayee)
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
		it.Event = new(IRoleSetterAlterPayee)
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
func (it *IRoleSetterAlterPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterAlterPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterAlterPayee represents a AlterPayee event raised by the IRoleSetter contract.
type IRoleSetterAlterPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAlterPayee is a free log retrieval operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) FilterAlterPayee(opts *bind.FilterOpts, payee []common.Address) (*IRoleSetterAlterPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterAlterPayeeIterator{contract: _IRoleSetter.contract, event: "AlterPayee", logs: logs, sub: sub}, nil
}

// WatchAlterPayee is a free log subscription operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) WatchAlterPayee(opts *bind.WatchOpts, sink chan<- *IRoleSetterAlterPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "AlterPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterAlterPayee)
				if err := _IRoleSetter.contract.UnpackLog(event, "AlterPayee", log); err != nil {
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

// ParseAlterPayee is a log parse operation binding the contract event 0x3b4058e93b2a019add56609a3c2463a7b3d8476abf30c934cfcdb9f2158b5e68.
//
// Solidity: event AlterPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) ParseAlterPayee(log types.Log) (*IRoleSetterAlterPayee, error) {
	event := new(IRoleSetterAlterPayee)
	if err := _IRoleSetter.contract.UnpackLog(event, "AlterPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterConfirmPayeeIterator is returned from FilterConfirmPayee and is used to iterate over the raw logs and unpacked data for ConfirmPayee events raised by the IRoleSetter contract.
type IRoleSetterConfirmPayeeIterator struct {
	Event *IRoleSetterConfirmPayee // Event containing the contract specifics and raw log

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
func (it *IRoleSetterConfirmPayeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterConfirmPayee)
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
		it.Event = new(IRoleSetterConfirmPayee)
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
func (it *IRoleSetterConfirmPayeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterConfirmPayeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterConfirmPayee represents a ConfirmPayee event raised by the IRoleSetter contract.
type IRoleSetterConfirmPayee struct {
	Payee common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterConfirmPayee is a free log retrieval operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) FilterConfirmPayee(opts *bind.FilterOpts, payee []common.Address) (*IRoleSetterConfirmPayeeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterConfirmPayeeIterator{contract: _IRoleSetter.contract, event: "ConfirmPayee", logs: logs, sub: sub}, nil
}

// WatchConfirmPayee is a free log subscription operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) WatchConfirmPayee(opts *bind.WatchOpts, sink chan<- *IRoleSetterConfirmPayee, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "ConfirmPayee", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterConfirmPayee)
				if err := _IRoleSetter.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
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

// ParseConfirmPayee is a log parse operation binding the contract event 0x4a21d873c33448e8dfb90cdbfdb12f4aed5f5e2abde5525215b1655d519784dc.
//
// Solidity: event ConfirmPayee(address indexed payee, address addr)
func (_IRoleSetter *IRoleSetterFilterer) ParseConfirmPayee(log types.Log) (*IRoleSetterConfirmPayee, error) {
	event := new(IRoleSetterConfirmPayee)
	if err := _IRoleSetter.contract.UnpackLog(event, "ConfirmPayee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterQuitRoleIterator is returned from FilterQuitRole and is used to iterate over the raw logs and unpacked data for QuitRole events raised by the IRoleSetter contract.
type IRoleSetterQuitRoleIterator struct {
	Event *IRoleSetterQuitRole // Event containing the contract specifics and raw log

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
func (it *IRoleSetterQuitRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterQuitRole)
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
		it.Event = new(IRoleSetterQuitRole)
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
func (it *IRoleSetterQuitRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterQuitRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterQuitRole represents a QuitRole event raised by the IRoleSetter contract.
type IRoleSetterQuitRole struct {
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterQuitRole is a free log retrieval operation binding the contract event 0x574022eb6ac5f8376ecb1e0a563dabd19419e1f59f38069278c6c67caeef005b.
//
// Solidity: event QuitRole(uint64 indexed index)
func (_IRoleSetter *IRoleSetterFilterer) FilterQuitRole(opts *bind.FilterOpts, index []uint64) (*IRoleSetterQuitRoleIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "QuitRole", indexRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterQuitRoleIterator{contract: _IRoleSetter.contract, event: "QuitRole", logs: logs, sub: sub}, nil
}

// WatchQuitRole is a free log subscription operation binding the contract event 0x574022eb6ac5f8376ecb1e0a563dabd19419e1f59f38069278c6c67caeef005b.
//
// Solidity: event QuitRole(uint64 indexed index)
func (_IRoleSetter *IRoleSetterFilterer) WatchQuitRole(opts *bind.WatchOpts, sink chan<- *IRoleSetterQuitRole, index []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "QuitRole", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterQuitRole)
				if err := _IRoleSetter.contract.UnpackLog(event, "QuitRole", log); err != nil {
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

// ParseQuitRole is a log parse operation binding the contract event 0x574022eb6ac5f8376ecb1e0a563dabd19419e1f59f38069278c6c67caeef005b.
//
// Solidity: event QuitRole(uint64 indexed index)
func (_IRoleSetter *IRoleSetterFilterer) ParseQuitRole(log types.Log) (*IRoleSetterQuitRole, error) {
	event := new(IRoleSetterQuitRole)
	if err := _IRoleSetter.contract.UnpackLog(event, "QuitRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterReAccIterator is returned from FilterReAcc and is used to iterate over the raw logs and unpacked data for ReAcc events raised by the IRoleSetter contract.
type IRoleSetterReAccIterator struct {
	Event *IRoleSetterReAcc // Event containing the contract specifics and raw log

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
func (it *IRoleSetterReAccIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterReAcc)
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
		it.Event = new(IRoleSetterReAcc)
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
func (it *IRoleSetterReAccIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterReAccIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterReAcc represents a ReAcc event raised by the IRoleSetter contract.
type IRoleSetterReAcc struct {
	Addr  common.Address
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReAcc is a free log retrieval operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) FilterReAcc(opts *bind.FilterOpts) (*IRoleSetterReAccIterator, error) {

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return &IRoleSetterReAccIterator{contract: _IRoleSetter.contract, event: "ReAcc", logs: logs, sub: sub}, nil
}

// WatchReAcc is a free log subscription operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) WatchReAcc(opts *bind.WatchOpts, sink chan<- *IRoleSetterReAcc) (event.Subscription, error) {

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "ReAcc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterReAcc)
				if err := _IRoleSetter.contract.UnpackLog(event, "ReAcc", log); err != nil {
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

// ParseReAcc is a log parse operation binding the contract event 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752.
//
// Solidity: event ReAcc(address addr, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) ParseReAcc(log types.Log) (*IRoleSetterReAcc, error) {
	event := new(IRoleSetterReAcc)
	if err := _IRoleSetter.contract.UnpackLog(event, "ReAcc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterReRoleIterator is returned from FilterReRole and is used to iterate over the raw logs and unpacked data for ReRole events raised by the IRoleSetter contract.
type IRoleSetterReRoleIterator struct {
	Event *IRoleSetterReRole // Event containing the contract specifics and raw log

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
func (it *IRoleSetterReRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterReRole)
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
		it.Event = new(IRoleSetterReRole)
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
func (it *IRoleSetterReRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterReRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterReRole represents a ReRole event raised by the IRoleSetter contract.
type IRoleSetterReRole struct {
	RType uint8
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReRole is a free log retrieval operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) FilterReRole(opts *bind.FilterOpts, rType []uint8) (*IRoleSetterReRoleIterator, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterReRoleIterator{contract: _IRoleSetter.contract, event: "ReRole", logs: logs, sub: sub}, nil
}

// WatchReRole is a free log subscription operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) WatchReRole(opts *bind.WatchOpts, sink chan<- *IRoleSetterReRole, rType []uint8) (event.Subscription, error) {

	var rTypeRule []interface{}
	for _, rTypeItem := range rType {
		rTypeRule = append(rTypeRule, rTypeItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "ReRole", rTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterReRole)
				if err := _IRoleSetter.contract.UnpackLog(event, "ReRole", log); err != nil {
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

// ParseReRole is a log parse operation binding the contract event 0x5550fbab402c3b014955818e5b6e3cefef20641aa6bed43ae83cd5d3258b8922.
//
// Solidity: event ReRole(uint8 indexed rType, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) ParseReRole(log types.Log) (*IRoleSetterReRole, error) {
	event := new(IRoleSetterReRole)
	if err := _IRoleSetter.contract.UnpackLog(event, "ReRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoleSetterSetGIterator is returned from FilterSetG and is used to iterate over the raw logs and unpacked data for SetG events raised by the IRoleSetter contract.
type IRoleSetterSetGIterator struct {
	Event *IRoleSetterSetG // Event containing the contract specifics and raw log

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
func (it *IRoleSetterSetGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoleSetterSetG)
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
		it.Event = new(IRoleSetterSetG)
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
func (it *IRoleSetterSetGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoleSetterSetGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoleSetterSetG represents a SetG event raised by the IRoleSetter contract.
type IRoleSetterSetG struct {
	Gi    uint64
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetG is a free log retrieval operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) FilterSetG(opts *bind.FilterOpts, gi []uint64) (*IRoleSetterSetGIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IRoleSetter.contract.FilterLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return &IRoleSetterSetGIterator{contract: _IRoleSetter.contract, event: "SetG", logs: logs, sub: sub}, nil
}

// WatchSetG is a free log subscription operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) WatchSetG(opts *bind.WatchOpts, sink chan<- *IRoleSetterSetG, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IRoleSetter.contract.WatchLogs(opts, "SetG", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoleSetterSetG)
				if err := _IRoleSetter.contract.UnpackLog(event, "SetG", log); err != nil {
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

// ParseSetG is a log parse operation binding the contract event 0xe7329265e8e4e2fc8e892a75f1809767d09c472bb2eef16f889eceb5583291a9.
//
// Solidity: event SetG(uint64 indexed gi, uint64 index)
func (_IRoleSetter *IRoleSetterFilterer) ParseSetG(log types.Log) (*IRoleSetterSetG, error) {
	event := new(IRoleSetterSetG)
	if err := _IRoleSetter.contract.UnpackLog(event, "SetG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenGetterABI is the input ABI used to generate the binding from.
const ITokenGetterABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITokenGetterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenGetterFuncSigs = map[string]string{
	"81abb8fe": "checkT(uint8)",
	"8bb4a637": "getTA(uint8)",
	"7600b86a": "getTCnt()",
	"2df2685f": "getTI(address)",
}

// ITokenGetter is an auto generated Go binding around an Ethereum contract.
type ITokenGetter struct {
	ITokenGetterCaller     // Read-only binding to the contract
	ITokenGetterTransactor // Write-only binding to the contract
	ITokenGetterFilterer   // Log filterer for contract events
}

// ITokenGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenGetterSession struct {
	Contract     *ITokenGetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenGetterCallerSession struct {
	Contract *ITokenGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenGetterTransactorSession struct {
	Contract     *ITokenGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenGetterRaw struct {
	Contract *ITokenGetter // Generic contract binding to access the raw methods on
}

// ITokenGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenGetterCallerRaw struct {
	Contract *ITokenGetterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenGetterTransactorRaw struct {
	Contract *ITokenGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenGetter creates a new instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetter(address common.Address, backend bind.ContractBackend) (*ITokenGetter, error) {
	contract, err := bindITokenGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenGetter{ITokenGetterCaller: ITokenGetterCaller{contract: contract}, ITokenGetterTransactor: ITokenGetterTransactor{contract: contract}, ITokenGetterFilterer: ITokenGetterFilterer{contract: contract}}, nil
}

// NewITokenGetterCaller creates a new read-only instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterCaller(address common.Address, caller bind.ContractCaller) (*ITokenGetterCaller, error) {
	contract, err := bindITokenGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterCaller{contract: contract}, nil
}

// NewITokenGetterTransactor creates a new write-only instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenGetterTransactor, error) {
	contract, err := bindITokenGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterTransactor{contract: contract}, nil
}

// NewITokenGetterFilterer creates a new log filterer instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenGetterFilterer, error) {
	contract, err := bindITokenGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterFilterer{contract: contract}, nil
}

// bindITokenGetter binds a generic wrapper to an already deployed contract.
func bindITokenGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenGetter *ITokenGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenGetter.Contract.ITokenGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenGetter *ITokenGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenGetter.Contract.ITokenGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenGetter *ITokenGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenGetter.Contract.ITokenGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenGetter *ITokenGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenGetter *ITokenGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenGetter *ITokenGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenGetter.Contract.contract.Transact(opts, method, params...)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterCaller) CheckT(opts *bind.CallOpts, _ti uint8) (common.Address, bool, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "checkT", _ti)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _ITokenGetter.Contract.CheckT(&_ITokenGetter.CallOpts, _ti)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterCallerSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _ITokenGetter.Contract.CheckT(&_ITokenGetter.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterCaller) GetTA(opts *bind.CallOpts, _ti uint8) (common.Address, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTA", _ti)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterSession) GetTA(_ti uint8) (common.Address, error) {
	return _ITokenGetter.Contract.GetTA(&_ITokenGetter.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterCallerSession) GetTA(_ti uint8) (common.Address, error) {
	return _ITokenGetter.Contract.GetTA(&_ITokenGetter.CallOpts, _ti)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterCaller) GetTCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterSession) GetTCnt() (uint8, error) {
	return _ITokenGetter.Contract.GetTCnt(&_ITokenGetter.CallOpts)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterCallerSession) GetTCnt() (uint8, error) {
	return _ITokenGetter.Contract.GetTCnt(&_ITokenGetter.CallOpts)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterCaller) GetTI(opts *bind.CallOpts, _t common.Address) (uint8, bool, bool, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTI", _t)

	if err != nil {
		return *new(uint8), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _ITokenGetter.Contract.GetTI(&_ITokenGetter.CallOpts, _t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterCallerSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _ITokenGetter.Contract.GetTI(&_ITokenGetter.CallOpts, _t)
}

// OwnerABI is the input ABI used to generate the binding from.
const OwnerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OwnerFuncSigs maps the 4-byte function signature to its string representation.
var OwnerFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"de9375f2": "auth()",
	"022914a7": "owners(address)",
}

// OwnerBin is the compiled bytecode used for deploying new contracts.
var OwnerBin = "0x608060405234801561001057600080fd5b5060405161054f38038061054f83398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6104bc806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063022914a7146100465780634bf1b1341461007e578063de9375f214610093575b600080fd5b610069610054366004610257565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61009161008c3660046102e9565b6100be565b005b6001546100a6906001600160a01b031681565b6040516001600160a01b039091168152602001610075565b823b600081900361010a5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b604482015260640160405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d9061019790849087906004016103ff565b600060405180830381600087803b1580156101b157600080fd5b505af11580156101c5573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b80356001600160a01b038116811461025257600080fd5b919050565b60006020828403121561026957600080fd5b6102728261023b565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156102b2576102b2610279565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156102e1576102e1610279565b604052919050565b6000806000606084860312156102fe57600080fd5b6103078461023b565b9250602080850135801515811461031d57600080fd5b9250604085013567ffffffffffffffff8082111561033a57600080fd5b8187019150601f888184011261034f57600080fd5b61035761028f565b8060a085018b81111561036957600080fd5b855b818110156103ed578035868111156103835760008081fd5b87018581018e136103945760008081fd5b8035878111156103a6576103a6610279565b6103b7818801601f19168b016102b8565b8181528f8b8385010111156103cc5760008081fd5b818b84018c83013760009181018b019190915285525092870192870161036b565b50508096505050505050509250925092565b8281526040602080830182905260009160e08401919084018584805b600581101561047857878603603f1901845282518051808852835b81811015610451578281018801518982018901528701610436565b508781018701849052601f01601f191690960185019550928401929184019160010161041b565b50939897505050505050505056fea264697066735822122051c6cfc7790ace9fc35fd7d24035222c18b0fbba24d8c520a60c25602823548b64736f6c63430008100033"

// DeployOwner deploys a new Ethereum contract, binding an instance of Owner to it.
func DeployOwner(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *Owner, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnerBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// Owner is an auto generated Go binding around an Ethereum contract.
type Owner struct {
	OwnerCaller     // Read-only binding to the contract
	OwnerTransactor // Write-only binding to the contract
	OwnerFilterer   // Log filterer for contract events
}

// OwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerSession struct {
	Contract     *Owner            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerCallerSession struct {
	Contract *OwnerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerTransactorSession struct {
	Contract     *OwnerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerRaw struct {
	Contract *Owner // Generic contract binding to access the raw methods on
}

// OwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerCallerRaw struct {
	Contract *OwnerCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerTransactorRaw struct {
	Contract *OwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwner creates a new instance of Owner, bound to a specific deployed contract.
func NewOwner(address common.Address, backend bind.ContractBackend) (*Owner, error) {
	contract, err := bindOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// NewOwnerCaller creates a new read-only instance of Owner, bound to a specific deployed contract.
func NewOwnerCaller(address common.Address, caller bind.ContractCaller) (*OwnerCaller, error) {
	contract, err := bindOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerCaller{contract: contract}, nil
}

// NewOwnerTransactor creates a new write-only instance of Owner, bound to a specific deployed contract.
func NewOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerTransactor, error) {
	contract, err := bindOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerTransactor{contract: contract}, nil
}

// NewOwnerFilterer creates a new log filterer instance of Owner, bound to a specific deployed contract.
func NewOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerFilterer, error) {
	contract, err := bindOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerFilterer{contract: contract}, nil
}

// bindOwner binds a generic wrapper to an already deployed contract.
func bindOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.OwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCallerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// OwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Owner contract.
type OwnerAddOwnerIterator struct {
	Event *OwnerAddOwner // Event containing the contract specifics and raw log

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
func (it *OwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAddOwner)
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
		it.Event = new(OwnerAddOwner)
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
func (it *OwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAddOwner represents a AddOwner event raised by the Owner contract.
type OwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*OwnerAddOwnerIterator, error) {

	logs, sub, err := _Owner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &OwnerAddOwnerIterator{contract: _Owner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *OwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _Owner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAddOwner)
				if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) ParseAddOwner(log types.Log) (*OwnerAddOwner, error) {
	event := new(OwnerAddOwner)
	if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecoverABI is the input ABI used to generate the binding from.
const RecoverABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// RecoverFuncSigs maps the 4-byte function signature to its string representation.
var RecoverFuncSigs = map[string]string{
	"19045a25": "recover(bytes32,bytes)",
}

// RecoverBin is the compiled bytecode used for deploying new contracts.
var RecoverBin = "0x61029c61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806319045a251461003a575b600080fd5b61004d610048366004610184565b610069565b6040516001600160a01b03909116815260200160405180910390f35b6000815160411461007c57506000610168565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156100c25760009350505050610168565b601b8160ff1610156100dc576100d981601b61023f565b90505b8060ff16601b141580156100f457508060ff16601c14155b156101055760009350505050610168565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610158573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561019757600080fd5b82359150602083013567ffffffffffffffff808211156101b657600080fd5b818501915085601f8301126101ca57600080fd5b8135818111156101dc576101dc61016e565b604051601f8201601f19908116603f011681019083821181831017156102045761020461016e565b8160405282815288602084870101111561021d57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60ff818116838216019081111561016857634e487b7160e01b600052601160045260246000fdfea2646970667358221220f76d0c480a62ad6ae5d1be3be88e99365d71ea8815b586481a71e66314a77cc264736f6c63430008100033"

// DeployRecover deploys a new Ethereum contract, binding an instance of Recover to it.
func DeployRecover(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recover, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RecoverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// Recover is an auto generated Go binding around an Ethereum contract.
type Recover struct {
	RecoverCaller     // Read-only binding to the contract
	RecoverTransactor // Write-only binding to the contract
	RecoverFilterer   // Log filterer for contract events
}

// RecoverCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecoverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecoverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecoverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecoverSession struct {
	Contract     *Recover          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecoverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecoverCallerSession struct {
	Contract *RecoverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RecoverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecoverTransactorSession struct {
	Contract     *RecoverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RecoverRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecoverRaw struct {
	Contract *Recover // Generic contract binding to access the raw methods on
}

// RecoverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecoverCallerRaw struct {
	Contract *RecoverCaller // Generic read-only contract binding to access the raw methods on
}

// RecoverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecoverTransactorRaw struct {
	Contract *RecoverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecover creates a new instance of Recover, bound to a specific deployed contract.
func NewRecover(address common.Address, backend bind.ContractBackend) (*Recover, error) {
	contract, err := bindRecover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// NewRecoverCaller creates a new read-only instance of Recover, bound to a specific deployed contract.
func NewRecoverCaller(address common.Address, caller bind.ContractCaller) (*RecoverCaller, error) {
	contract, err := bindRecover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverCaller{contract: contract}, nil
}

// NewRecoverTransactor creates a new write-only instance of Recover, bound to a specific deployed contract.
func NewRecoverTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoverTransactor, error) {
	contract, err := bindRecover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverTransactor{contract: contract}, nil
}

// NewRecoverFilterer creates a new log filterer instance of Recover, bound to a specific deployed contract.
func NewRecoverFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoverFilterer, error) {
	contract, err := bindRecover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoverFilterer{contract: contract}, nil
}

// bindRecover binds a generic wrapper to an already deployed contract.
func bindRecover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.RecoverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transact(opts, method, params...)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCaller) Recover(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Recover.contract.Call(opts, &out, "recover", hash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCallerSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}
