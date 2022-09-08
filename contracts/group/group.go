// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package group

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

// GroupABI is the input ABI used to generate the binding from.
const GroupABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inst\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"Create\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_kc\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_pm\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_kpPB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getDesc\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getPInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getPlePerB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inst\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"setDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pp\",\"type\":\"uint256\"}],\"name\":\"setP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"}],\"name\":\"setPerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"kcnt\",\"type\":\"uint8\"}],\"name\":\"setS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GroupFuncSigs maps the 4-byte function signature to its string representation.
var GroupFuncSigs = map[string]string{
	"ad578aeb": "activate(uint64,uint8)",
	"4bf1b134": "add(address,bool,bytes[5])",
	"4acc2468": "add(uint8,uint64,uint256)",
	"de9375f2": "auth()",
	"ae38ca77": "create(uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes)",
	"2a2722ad": "getDesc(uint64)",
	"059e783d": "getGCnt()",
	"f6070374": "getGS(uint64)",
	"770609a8": "getLevel(uint64)",
	"50aa2329": "getPInfo(uint64)",
	"60b5903d": "getPlePerB(uint64)",
	"e795ea8a": "getRate(uint64)",
	"bd6b2222": "inst()",
	"022914a7": "owners(address)",
	"773854ed": "setDesc(uint64,bytes)",
	"936945bd": "setLevel(uint64,uint8)",
	"2a0212e5": "setP(uint64,uint256,uint256)",
	"cfb9854f": "setPerB(uint64,uint256,uint256)",
	"1e752182": "setRate(uint64,uint8,uint8)",
	"5f9dbbf9": "setS(uint64,uint8,uint8)",
	"54fd4d50": "version()",
}

// GroupBin is the compiled bytecode used for deploying new contracts.
var GroupBin = "0x60806040526001805461ffff60a01b1916600160a11b1790553480156200002557600080fd5b5060405162001b8638038062001b8683398101604081905262000048916200023c565b600180546001600160a01b03199081166001600160a01b0385811691909117835560028054909216908416179055604080516101208101825260006020820181815292820181815260608084018381526080850184815260a0860185815260c0870186815260e08801878152610100808a019687528b8a52600380549c8d01815590985288516006909b027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810180549b519851965160ff90811663010000000263ff0000001998821662010000029890981663ffff0000199a8216909b0261ffff19909d169d169c909c179a909a179690961696909617929092178855517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c870155517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d86015591517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85e850155517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85f840155519092839290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f8609091019062000213908262000319565b505050505050620003e5565b80516001600160a01b03811681146200023757600080fd5b919050565b600080604083850312156200025057600080fd5b6200025b836200021f565b91506200026b602084016200021f565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200029f57607f821691505b602082108103620002c057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200031457600081815260208120601f850160051c81016020861015620002ef5750805b601f850160051c820191505b818110156200031057828155600101620002fb565b5050505b505050565b81516001600160401b0381111562000335576200033562000274565b6200034d816200034684546200028a565b84620002c6565b602080601f8311600181146200038557600084156200036c5750858301515b600019600386901b1c1916600185901b17855562000310565b600085815260208120601f198616915b82811015620003b65788860151825594840194600190910190840162000395565b5085821015620003d55787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61179180620003f56000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806360b5903d116100b8578063ae38ca771161007c578063ae38ca77146102d1578063bd6b2222146102e4578063cfb9854f1461030f578063de9375f214610322578063e795ea8a14610335578063f60703741461036257600080fd5b806360b5903d14610260578063770609a814610273578063773854ed14610298578063936945bd146102ab578063ad578aeb146102be57600080fd5b80634acc2468116100ff5780634acc2468146101d75780634bf1b134146101ea57806350aa2329146101fd57806354fd4d50146102255780635f9dbbf91461024d57600080fd5b8063022914a71461013c578063059e783d146101745780631e7521821461018f5780632a0212e5146101a45780632a2722ad146101b7575b600080fd5b61015f61014a36600461110b565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6003546040516001600160401b03909116815260200161016b565b6101a261019d366004611155565b610375565b005b6101a26101b2366004611198565b61043e565b6101ca6101c53660046111cb565b6104d6565b60405161016b919061122c565b6101a26101e536600461123f565b610595565b6101a26101f8366004611345565b6107f5565b61021061020b3660046111cb565b61096e565b6040805192835260208301919091520161016b565b60015461023a90600160a01b900461ffff1681565b60405161ffff909116815260200161016b565b6101a261025b366004611155565b6109d6565b61021061026e3660046111cb565b610acc565b6102866102813660046111cb565b610b34565b60405160ff909116815260200161016b565b6101a26102a6366004611408565b610b71565b6101a26102b9366004611455565b610bda565b6101a26102cc366004611455565b610c53565b6101a26102df366004611488565b610d73565b6002546102f7906001600160a01b031681565b6040516001600160a01b03909116815260200161016b565b6101a261031d366004611198565b610f9d565b6001546102f7906001600160a01b031681565b6103486103433660046111cb565b611035565b6040805160ff93841681529290911660208301520161016b565b6102866103703660046111cb565b6110b7565b3360009081526020819052604090205460ff166103ad5760405162461bcd60e51b81526004016103a49061151d565b60405180910390fd5b816003846001600160401b0316815481106103ca576103ca611540565b906000526020600020906006020160000160026101000a81548160ff021916908360ff160217905550806003846001600160401b03168154811061041057610410611540565b906000526020600020906006020160000160036101000a81548160ff021916908360ff160217905550505050565b3360009081526020819052604090205460ff1661046d5760405162461bcd60e51b81526004016103a49061151d565b816003846001600160401b03168154811061048a5761048a611540565b906000526020600020906006020160010181905550806003846001600160401b0316815481106104bc576104bc611540565b906000526020600020906006020160020181905550505050565b60606003826001600160401b0316815481106104f4576104f4611540565b9060005260206000209060060201600501805461051090611556565b80601f016020809104026020016040519081016040528092919081815260200182805461053c90611556565b80156105895780601f1061055e57610100808354040283529160200191610589565b820191906000526020600020905b81548152906001019060200180831161056c57829003601f168201915b50505050509050919050565b6003826001600160401b0316815481106105b1576105b1611540565b600091825260209091206006909102015460ff166001036106035760405162461bcd60e51b815260206004820152600c60248201526b1859190e99c818985b9b995960a21b60448201526064016103a4565b8260ff16600103610686576003826001600160401b03168154811061062a5761062a611540565b600091825260209091206006909102015460ff166003146106815760405162461bcd60e51b8152602060048201526011602482015270616464753a67206e6f742061637469766560781b60448201526064016103a4565b505050565b8260ff16600203610777576003826001600160401b0316815481106106ad576106ad611540565b600091825260209091206006909102015460ff166003146107045760405162461bcd60e51b8152602060048201526011602482015270616464503a67206e6f742061637469766560781b60448201526064016103a4565b6003826001600160401b03168154811061072057610720611540565b9060005260206000209060060201600201548110156106815760405162461bcd60e51b81526020600482015260136024820152720e040e0d8cac8ceca40dcdee840cadcdeeaced606b1b60448201526064016103a4565b8260ff16600303610681576003826001600160401b03168154811061079e5761079e611540565b9060005260206000209060060201600101548110156106815760405162461bcd60e51b81526020600482015260136024820152720d640e0d8cac8ceca40dcdee840cadcdeeaced606b1b60448201526064016103a4565b823b600081900361083d5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064016103a4565b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906108ca9084908790600401611590565b600060405180830381600087803b1580156108e457600080fd5b505af11580156108f8573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b6000806003836001600160401b03168154811061098d5761098d611540565b9060005260206000209060060201600101546003846001600160401b0316815481106109bb576109bb611540565b90600052602060002090600602016002015491509150915091565b3360009081526020819052604090205460ff16610a055760405162461bcd60e51b81526004016103a49061151d565b8160ff16600303610a88576003836001600160401b031681548110610a2c57610a2c611540565b600091825260209091206006909102015460ff61010090910481169082161015610a885760405162461bcd60e51b815260206004820152600d60248201526c0d6e640dcdee840cadcdeeaced609b1b60448201526064016103a4565b816003846001600160401b031681548110610aa557610aa5611540565b60009182526020909120600690910201805460ff191660ff92909216919091179055505050565b6000806003836001600160401b031681548110610aeb57610aeb611540565b9060005260206000209060060201600301546003846001600160401b031681548110610b1957610b19611540565b90600052602060002090600602016004015491509150915091565b60006003826001600160401b031681548110610b5257610b52611540565b6000918252602090912060069091020154610100900460ff1692915050565b3360009081526020819052604090205460ff16610ba05760405162461bcd60e51b81526004016103a49061151d565b806003836001600160401b031681548110610bbd57610bbd611540565b906000526020600020906006020160050190816106819190611635565b3360009081526020819052604090205460ff16610c095760405162461bcd60e51b81526004016103a49061151d565b806003836001600160401b031681548110610c2657610c26611540565b906000526020600020906006020160000160016101000a81548160ff021916908360ff1602179055505050565b3360009081526020819052604090205460ff16610c825760405162461bcd60e51b81526004016103a49061151d565b6003826001600160401b031681548110610c9e57610c9e611540565b600091825260209091206006909102015460ff16600103610cf05760405162461bcd60e51b815260206004820152600c60248201526b1858dd0e99c818985b9b995960a21b60448201526064016103a4565b6003826001600160401b031681548110610d0c57610d0c611540565b600091825260209091206006909102015460ff610100909104811690821610610d6f57600380836001600160401b031681548110610d4c57610d4c611540565b60009182526020909120600690910201805460ff191660ff929092169190911790555b5050565b3360009081526020819052604090205460ff16610da25760405162461bcd60e51b81526004016103a49061151d565b60038054604080516101208101825260ff8c811660208301908152608083018b815260a084018b81528e84169585019586528d84166060860190815260c086018c815260e087018c81526101008089018d81526001808b528c018d5560009c909c52885160068c027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101805499519c5196518b1663010000000263ff00000019978c1662010000029790971663ffff0000199d8c1690940261ffff19909a1692909a169190911797909717999099169890981791909117855591517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c840155517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d830155517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85e82015592517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85f84015593519293909283927fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f8600190610f479082611635565b5050507f4b14fb8b0a4d249a53bfa9dadcdeab4977149168de6dd5efa8a37e41107b3fda828b8b8b8b8b8b8b8b604051610f89999897969594939291906116f4565b60405180910390a150505050505050505050565b3360009081526020819052604090205460ff16610fcc5760405162461bcd60e51b81526004016103a49061151d565b816003846001600160401b031681548110610fe957610fe9611540565b906000526020600020906006020160030181905550806003846001600160401b03168154811061101b5761101b611540565b906000526020600020906006020160040181905550505050565b6000806003836001600160401b03168154811061105457611054611540565b906000526020600020906006020160000160029054906101000a900460ff166003846001600160401b03168154811061108f5761108f611540565b906000526020600020906006020160000160039054906101000a900460ff1691509150915091565b60006003826001600160401b0316815481106110d5576110d5611540565b600091825260209091206006909102015460ff1692915050565b80356001600160a01b038116811461110657600080fd5b919050565b60006020828403121561111d57600080fd5b611126826110ef565b9392505050565b80356001600160401b038116811461110657600080fd5b803560ff8116811461110657600080fd5b60008060006060848603121561116a57600080fd5b6111738461112d565b925061118160208501611144565b915061118f60408501611144565b90509250925092565b6000806000606084860312156111ad57600080fd5b6111b68461112d565b95602085013595506040909401359392505050565b6000602082840312156111dd57600080fd5b6111268261112d565b6000815180845260005b8181101561120c576020818501810151868301820152016111f0565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061112660208301846111e6565b60008060006060848603121561125457600080fd5b61125d84611144565b925061126b6020850161112d565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156112b3576112b361127b565b60405290565b600082601f8301126112ca57600080fd5b81356001600160401b03808211156112e4576112e461127b565b604051601f8301601f19908116603f0116810190828211818310171561130c5761130c61127b565b8160405283815286602085880101111561132557600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561135a57600080fd5b611363846110ef565b9250602080850135801515811461137957600080fd5b925060408501356001600160401b038082111561139557600080fd5b818701915087601f8301126113a957600080fd5b6113b1611291565b8060a084018a8111156113c357600080fd5b845b818110156113f7578035858111156113dd5760008081fd5b6113e98d8289016112b9565b8552509286019286016113c5565b505080955050505050509250925092565b6000806040838503121561141b57600080fd5b6114248361112d565b915060208301356001600160401b0381111561143f57600080fd5b61144b858286016112b9565b9150509250929050565b6000806040838503121561146857600080fd5b6114718361112d565b915061147f60208401611144565b90509250929050565b600080600080600080600080610100898b0312156114a557600080fd5b6114ae89611144565b97506114bc60208a01611144565b96506114ca60408a01611144565b9550606089013594506080890135935060a0890135925060c0890135915060e08901356001600160401b0381111561150157600080fd5b61150d8b828c016112b9565b9150509295985092959890939650565b6020808252600990820152683737ba1037bbb732b960b91b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600181811c9082168061156a57607f821691505b60208210810361158a57634e487b7160e01b600052602260045260246000fd5b50919050565b8281526040602080830182905260009160e084019190840185845b60058110156115da57603f198786030183526115c88583516111e6565b945091830191908301906001016115ab565b5092979650505050505050565b601f82111561068157600081815260208120601f850160051c8101602086101561160e5750805b601f850160051c820191505b8181101561162d5782815560010161161a565b505050505050565b81516001600160401b0381111561164e5761164e61127b565b6116628161165c8454611556565b846115e7565b602080601f831160018114611697576000841561167f5750858301515b600019600386901b1c1916600185901b17855561162d565b600085815260208120601f198616915b828110156116c6578886015182559484019460019091019084016116a7565b50858210156116e45787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006101206001600160401b038c16835260ff8b16602084015260ff8a16604084015260ff891660608401528760808401528660a08401528560c08401528460e08401528061010084015261174b818401856111e6565b9c9b50505050505050505050505056fea2646970667358221220a24830cc06e1969cea6be7fd9cb2846e5449b4ffc26330424fe590e564e92edc64736f6c63430008100033"

// DeployGroup deploys a new Ethereum contract, binding an instance of Group to it.
func DeployGroup(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, _inst common.Address) (common.Address, *types.Transaction, *Group, error) {
	parsed, err := abi.JSON(strings.NewReader(GroupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GroupBin), backend, _a, _inst)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Group{GroupCaller: GroupCaller{contract: contract}, GroupTransactor: GroupTransactor{contract: contract}, GroupFilterer: GroupFilterer{contract: contract}}, nil
}

// Group is an auto generated Go binding around an Ethereum contract.
type Group struct {
	GroupCaller     // Read-only binding to the contract
	GroupTransactor // Write-only binding to the contract
	GroupFilterer   // Log filterer for contract events
}

// GroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type GroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GroupSession struct {
	Contract     *Group            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GroupCallerSession struct {
	Contract *GroupCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GroupTransactorSession struct {
	Contract     *GroupTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type GroupRaw struct {
	Contract *Group // Generic contract binding to access the raw methods on
}

// GroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GroupCallerRaw struct {
	Contract *GroupCaller // Generic read-only contract binding to access the raw methods on
}

// GroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GroupTransactorRaw struct {
	Contract *GroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGroup creates a new instance of Group, bound to a specific deployed contract.
func NewGroup(address common.Address, backend bind.ContractBackend) (*Group, error) {
	contract, err := bindGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Group{GroupCaller: GroupCaller{contract: contract}, GroupTransactor: GroupTransactor{contract: contract}, GroupFilterer: GroupFilterer{contract: contract}}, nil
}

// NewGroupCaller creates a new read-only instance of Group, bound to a specific deployed contract.
func NewGroupCaller(address common.Address, caller bind.ContractCaller) (*GroupCaller, error) {
	contract, err := bindGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GroupCaller{contract: contract}, nil
}

// NewGroupTransactor creates a new write-only instance of Group, bound to a specific deployed contract.
func NewGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*GroupTransactor, error) {
	contract, err := bindGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GroupTransactor{contract: contract}, nil
}

// NewGroupFilterer creates a new log filterer instance of Group, bound to a specific deployed contract.
func NewGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*GroupFilterer, error) {
	contract, err := bindGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GroupFilterer{contract: contract}, nil
}

// bindGroup binds a generic wrapper to an already deployed contract.
func bindGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Group *GroupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Group.Contract.GroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Group *GroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Group.Contract.GroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Group *GroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Group.Contract.GroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Group *GroupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Group.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Group *GroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Group.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Group *GroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Group.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_Group *GroupCaller) Add(opts *bind.CallOpts, _rType uint8, _gi uint64, _pm *big.Int) error {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "add", _rType, _gi, _pm)

	if err != nil {
		return err
	}

	return err

}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_Group *GroupSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _Group.Contract.Add(&_Group.CallOpts, _rType, _gi, _pm)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_Group *GroupCallerSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _Group.Contract.Add(&_Group.CallOpts, _rType, _gi, _pm)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Group *GroupCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Group *GroupSession) Auth() (common.Address, error) {
	return _Group.Contract.Auth(&_Group.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Group *GroupCallerSession) Auth() (common.Address, error) {
	return _Group.Contract.Auth(&_Group.CallOpts)
}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_Group *GroupCaller) GetDesc(opts *bind.CallOpts, i uint64) ([]byte, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getDesc", i)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_Group *GroupSession) GetDesc(i uint64) ([]byte, error) {
	return _Group.Contract.GetDesc(&_Group.CallOpts, i)
}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_Group *GroupCallerSession) GetDesc(i uint64) ([]byte, error) {
	return _Group.Contract.GetDesc(&_Group.CallOpts, i)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_Group *GroupCaller) GetGCnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getGCnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_Group *GroupSession) GetGCnt() (uint64, error) {
	return _Group.Contract.GetGCnt(&_Group.CallOpts)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_Group *GroupCallerSession) GetGCnt() (uint64, error) {
	return _Group.Contract.GetGCnt(&_Group.CallOpts)
}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_Group *GroupCaller) GetGS(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getGS", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_Group *GroupSession) GetGS(i uint64) (uint8, error) {
	return _Group.Contract.GetGS(&_Group.CallOpts, i)
}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_Group *GroupCallerSession) GetGS(i uint64) (uint8, error) {
	return _Group.Contract.GetGS(&_Group.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_Group *GroupCaller) GetLevel(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getLevel", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_Group *GroupSession) GetLevel(i uint64) (uint8, error) {
	return _Group.Contract.GetLevel(&_Group.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_Group *GroupCallerSession) GetLevel(i uint64) (uint8, error) {
	return _Group.Contract.GetLevel(&_Group.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 i) view returns(uint256, uint256)
func (_Group *GroupCaller) GetPInfo(opts *bind.CallOpts, i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getPInfo", i)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 i) view returns(uint256, uint256)
func (_Group *GroupSession) GetPInfo(i uint64) (*big.Int, *big.Int, error) {
	return _Group.Contract.GetPInfo(&_Group.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 i) view returns(uint256, uint256)
func (_Group *GroupCallerSession) GetPInfo(i uint64) (*big.Int, *big.Int, error) {
	return _Group.Contract.GetPInfo(&_Group.CallOpts, i)
}

// GetPlePerB is a free data retrieval call binding the contract method 0x60b5903d.
//
// Solidity: function getPlePerB(uint64 i) view returns(uint256, uint256)
func (_Group *GroupCaller) GetPlePerB(opts *bind.CallOpts, i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getPlePerB", i)

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
func (_Group *GroupSession) GetPlePerB(i uint64) (*big.Int, *big.Int, error) {
	return _Group.Contract.GetPlePerB(&_Group.CallOpts, i)
}

// GetPlePerB is a free data retrieval call binding the contract method 0x60b5903d.
//
// Solidity: function getPlePerB(uint64 i) view returns(uint256, uint256)
func (_Group *GroupCallerSession) GetPlePerB(i uint64) (*big.Int, *big.Int, error) {
	return _Group.Contract.GetPlePerB(&_Group.CallOpts, i)
}

// GetRate is a free data retrieval call binding the contract method 0xe795ea8a.
//
// Solidity: function getRate(uint64 i) view returns(uint8, uint8)
func (_Group *GroupCaller) GetRate(opts *bind.CallOpts, i uint64) (uint8, uint8, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getRate", i)

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
func (_Group *GroupSession) GetRate(i uint64) (uint8, uint8, error) {
	return _Group.Contract.GetRate(&_Group.CallOpts, i)
}

// GetRate is a free data retrieval call binding the contract method 0xe795ea8a.
//
// Solidity: function getRate(uint64 i) view returns(uint8, uint8)
func (_Group *GroupCallerSession) GetRate(i uint64) (uint8, uint8, error) {
	return _Group.Contract.GetRate(&_Group.CallOpts, i)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Group *GroupCaller) Inst(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "inst")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Group *GroupSession) Inst() (common.Address, error) {
	return _Group.Contract.Inst(&_Group.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Group *GroupCallerSession) Inst() (common.Address, error) {
	return _Group.Contract.Inst(&_Group.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Group *GroupCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Group *GroupSession) Owners(arg0 common.Address) (bool, error) {
	return _Group.Contract.Owners(&_Group.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Group *GroupCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Group.Contract.Owners(&_Group.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Group *GroupCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Group *GroupSession) Version() (uint16, error) {
	return _Group.Contract.Version(&_Group.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Group *GroupCallerSession) Version() (uint16, error) {
	return _Group.Contract.Version(&_Group.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_Group *GroupTransactor) Activate(opts *bind.TransactOpts, _gi uint64, _kc uint8) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "activate", _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_Group *GroupSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _Group.Contract.Activate(&_Group.TransactOpts, _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_Group *GroupTransactorSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _Group.Contract.Activate(&_Group.TransactOpts, _gi, _kc)
}

// Add0 is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Group *GroupTransactor) Add0(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "add0", _a, isSet, signs)
}

// Add0 is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Group *GroupSession) Add0(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Group.Contract.Add0(&_Group.TransactOpts, _a, isSet, signs)
}

// Add0 is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Group *GroupTransactorSession) Add0(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Group.Contract.Add0(&_Group.TransactOpts, _a, isSet, signs)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPB, uint256 _ppPB, bytes _desc) returns()
func (_Group *GroupTransactor) Create(opts *bind.TransactOpts, _level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPB *big.Int, _ppPB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "create", _level, _mr, _tr, _kr, _pr, _kpPB, _ppPB, _desc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPB, uint256 _ppPB, bytes _desc) returns()
func (_Group *GroupSession) Create(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPB *big.Int, _ppPB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _Group.Contract.Create(&_Group.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPB, _ppPB, _desc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPB, uint256 _ppPB, bytes _desc) returns()
func (_Group *GroupTransactorSession) Create(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPB *big.Int, _ppPB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _Group.Contract.Create(&_Group.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPB, _ppPB, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_Group *GroupTransactor) SetDesc(opts *bind.TransactOpts, _gi uint64, _desc []byte) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "setDesc", _gi, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_Group *GroupSession) SetDesc(_gi uint64, _desc []byte) (*types.Transaction, error) {
	return _Group.Contract.SetDesc(&_Group.TransactOpts, _gi, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_Group *GroupTransactorSession) SetDesc(_gi uint64, _desc []byte) (*types.Transaction, error) {
	return _Group.Contract.SetDesc(&_Group.TransactOpts, _gi, _desc)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_Group *GroupTransactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "setLevel", _gi, _le)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_Group *GroupSession) SetLevel(_gi uint64, _le uint8) (*types.Transaction, error) {
	return _Group.Contract.SetLevel(&_Group.TransactOpts, _gi, _le)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_Group *GroupTransactorSession) SetLevel(_gi uint64, _le uint8) (*types.Transaction, error) {
	return _Group.Contract.SetLevel(&_Group.TransactOpts, _gi, _le)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_Group *GroupTransactor) SetP(opts *bind.TransactOpts, _gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "setP", _gi, _kp, _pp)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_Group *GroupSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _Group.Contract.SetP(&_Group.TransactOpts, _gi, _kp, _pp)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_Group *GroupTransactorSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _Group.Contract.SetP(&_Group.TransactOpts, _gi, _kp, _pp)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_Group *GroupTransactor) SetPerB(opts *bind.TransactOpts, _gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "setPerB", _gi, _kpPerB, _ppPerB)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_Group *GroupSession) SetPerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _Group.Contract.SetPerB(&_Group.TransactOpts, _gi, _kpPerB, _ppPerB)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_Group *GroupTransactorSession) SetPerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _Group.Contract.SetPerB(&_Group.TransactOpts, _gi, _kpPerB, _ppPerB)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_Group *GroupTransactor) SetRate(opts *bind.TransactOpts, _gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "setRate", _gi, _mr, _tr)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_Group *GroupSession) SetRate(_gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _Group.Contract.SetRate(&_Group.TransactOpts, _gi, _mr, _tr)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_Group *GroupTransactorSession) SetRate(_gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _Group.Contract.SetRate(&_Group.TransactOpts, _gi, _mr, _tr)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_Group *GroupTransactor) SetS(opts *bind.TransactOpts, _gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "setS", _gi, _s, kcnt)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_Group *GroupSession) SetS(_gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _Group.Contract.SetS(&_Group.TransactOpts, _gi, _s, kcnt)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_Group *GroupTransactorSession) SetS(_gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _Group.Contract.SetS(&_Group.TransactOpts, _gi, _s, kcnt)
}

// GroupAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Group contract.
type GroupAddOwnerIterator struct {
	Event *GroupAddOwner // Event containing the contract specifics and raw log

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
func (it *GroupAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupAddOwner)
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
		it.Event = new(GroupAddOwner)
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
func (it *GroupAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupAddOwner represents a AddOwner event raised by the Group contract.
type GroupAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Group *GroupFilterer) FilterAddOwner(opts *bind.FilterOpts) (*GroupAddOwnerIterator, error) {

	logs, sub, err := _Group.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &GroupAddOwnerIterator{contract: _Group.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Group *GroupFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *GroupAddOwner) (event.Subscription, error) {

	logs, sub, err := _Group.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupAddOwner)
				if err := _Group.contract.UnpackLog(event, "AddOwner", log); err != nil {
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
func (_Group *GroupFilterer) ParseAddOwner(log types.Log) (*GroupAddOwner, error) {
	event := new(GroupAddOwner)
	if err := _Group.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GroupCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the Group contract.
type GroupCreateIterator struct {
	Event *GroupCreate // Event containing the contract specifics and raw log

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
func (it *GroupCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupCreate)
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
		it.Event = new(GroupCreate)
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
func (it *GroupCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupCreate represents a Create event raised by the Group contract.
type GroupCreate struct {
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
func (_Group *GroupFilterer) FilterCreate(opts *bind.FilterOpts) (*GroupCreateIterator, error) {

	logs, sub, err := _Group.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &GroupCreateIterator{contract: _Group.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x4b14fb8b0a4d249a53bfa9dadcdeab4977149168de6dd5efa8a37e41107b3fda.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 tr, uint256 kr, uint256 pr, uint256 kPerB, uint256 pPerB, bytes _desc)
func (_Group *GroupFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *GroupCreate) (event.Subscription, error) {

	logs, sub, err := _Group.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupCreate)
				if err := _Group.contract.UnpackLog(event, "Create", log); err != nil {
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
func (_Group *GroupFilterer) ParseCreate(log types.Log) (*GroupCreate, error) {
	event := new(GroupCreate)
	if err := _Group.contract.UnpackLog(event, "Create", log); err != nil {
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

// IGroupABI is the input ABI used to generate the binding from.
const IGroupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"Create\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_kc\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_pm\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_kPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getDesc\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getPInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getPlePerB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"setDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pp\",\"type\":\"uint256\"}],\"name\":\"setP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"}],\"name\":\"setPerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"kcnt\",\"type\":\"uint8\"}],\"name\":\"setS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGroupFuncSigs maps the 4-byte function signature to its string representation.
var IGroupFuncSigs = map[string]string{
	"ad578aeb": "activate(uint64,uint8)",
	"4acc2468": "add(uint8,uint64,uint256)",
	"ae38ca77": "create(uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes)",
	"2a2722ad": "getDesc(uint64)",
	"059e783d": "getGCnt()",
	"f6070374": "getGS(uint64)",
	"770609a8": "getLevel(uint64)",
	"50aa2329": "getPInfo(uint64)",
	"60b5903d": "getPlePerB(uint64)",
	"e795ea8a": "getRate(uint64)",
	"773854ed": "setDesc(uint64,bytes)",
	"936945bd": "setLevel(uint64,uint8)",
	"2a0212e5": "setP(uint64,uint256,uint256)",
	"cfb9854f": "setPerB(uint64,uint256,uint256)",
	"1e752182": "setRate(uint64,uint8,uint8)",
	"5f9dbbf9": "setS(uint64,uint8,uint8)",
}

// IGroup is an auto generated Go binding around an Ethereum contract.
type IGroup struct {
	IGroupCaller     // Read-only binding to the contract
	IGroupTransactor // Write-only binding to the contract
	IGroupFilterer   // Log filterer for contract events
}

// IGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGroupSession struct {
	Contract     *IGroup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGroupCallerSession struct {
	Contract *IGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGroupTransactorSession struct {
	Contract     *IGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGroupRaw struct {
	Contract *IGroup // Generic contract binding to access the raw methods on
}

// IGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGroupCallerRaw struct {
	Contract *IGroupCaller // Generic read-only contract binding to access the raw methods on
}

// IGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGroupTransactorRaw struct {
	Contract *IGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGroup creates a new instance of IGroup, bound to a specific deployed contract.
func NewIGroup(address common.Address, backend bind.ContractBackend) (*IGroup, error) {
	contract, err := bindIGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGroup{IGroupCaller: IGroupCaller{contract: contract}, IGroupTransactor: IGroupTransactor{contract: contract}, IGroupFilterer: IGroupFilterer{contract: contract}}, nil
}

// NewIGroupCaller creates a new read-only instance of IGroup, bound to a specific deployed contract.
func NewIGroupCaller(address common.Address, caller bind.ContractCaller) (*IGroupCaller, error) {
	contract, err := bindIGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupCaller{contract: contract}, nil
}

// NewIGroupTransactor creates a new write-only instance of IGroup, bound to a specific deployed contract.
func NewIGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*IGroupTransactor, error) {
	contract, err := bindIGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupTransactor{contract: contract}, nil
}

// NewIGroupFilterer creates a new log filterer instance of IGroup, bound to a specific deployed contract.
func NewIGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*IGroupFilterer, error) {
	contract, err := bindIGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGroupFilterer{contract: contract}, nil
}

// bindIGroup binds a generic wrapper to an already deployed contract.
func bindIGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroup *IGroupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroup.Contract.IGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroup *IGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroup.Contract.IGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroup *IGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroup.Contract.IGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroup *IGroupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroup *IGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroup *IGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroup.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroup *IGroupCaller) Add(opts *bind.CallOpts, _rType uint8, _gi uint64, _pm *big.Int) error {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "add", _rType, _gi, _pm)

	if err != nil {
		return err
	}

	return err

}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroup *IGroupSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroup.Contract.Add(&_IGroup.CallOpts, _rType, _gi, _pm)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroup *IGroupCallerSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroup.Contract.Add(&_IGroup.CallOpts, _rType, _gi, _pm)
}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_IGroup *IGroupCaller) GetDesc(opts *bind.CallOpts, i uint64) ([]byte, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getDesc", i)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_IGroup *IGroupSession) GetDesc(i uint64) ([]byte, error) {
	return _IGroup.Contract.GetDesc(&_IGroup.CallOpts, i)
}

// GetDesc is a free data retrieval call binding the contract method 0x2a2722ad.
//
// Solidity: function getDesc(uint64 i) view returns(bytes)
func (_IGroup *IGroupCallerSession) GetDesc(i uint64) ([]byte, error) {
	return _IGroup.Contract.GetDesc(&_IGroup.CallOpts, i)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroup *IGroupCaller) GetGCnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getGCnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroup *IGroupSession) GetGCnt() (uint64, error) {
	return _IGroup.Contract.GetGCnt(&_IGroup.CallOpts)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroup *IGroupCallerSession) GetGCnt() (uint64, error) {
	return _IGroup.Contract.GetGCnt(&_IGroup.CallOpts)
}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_IGroup *IGroupCaller) GetGS(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getGS", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_IGroup *IGroupSession) GetGS(i uint64) (uint8, error) {
	return _IGroup.Contract.GetGS(&_IGroup.CallOpts, i)
}

// GetGS is a free data retrieval call binding the contract method 0xf6070374.
//
// Solidity: function getGS(uint64 i) view returns(uint8)
func (_IGroup *IGroupCallerSession) GetGS(i uint64) (uint8, error) {
	return _IGroup.Contract.GetGS(&_IGroup.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroup *IGroupCaller) GetLevel(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getLevel", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroup *IGroupSession) GetLevel(i uint64) (uint8, error) {
	return _IGroup.Contract.GetLevel(&_IGroup.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroup *IGroupCallerSession) GetLevel(i uint64) (uint8, error) {
	return _IGroup.Contract.GetLevel(&_IGroup.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroup *IGroupCaller) GetPInfo(opts *bind.CallOpts, _i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getPInfo", _i)

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
func (_IGroup *IGroupSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroup.Contract.GetPInfo(&_IGroup.CallOpts, _i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroup *IGroupCallerSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroup.Contract.GetPInfo(&_IGroup.CallOpts, _i)
}

// GetPlePerB is a free data retrieval call binding the contract method 0x60b5903d.
//
// Solidity: function getPlePerB(uint64 i) view returns(uint256, uint256)
func (_IGroup *IGroupCaller) GetPlePerB(opts *bind.CallOpts, i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getPlePerB", i)

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
func (_IGroup *IGroupSession) GetPlePerB(i uint64) (*big.Int, *big.Int, error) {
	return _IGroup.Contract.GetPlePerB(&_IGroup.CallOpts, i)
}

// GetPlePerB is a free data retrieval call binding the contract method 0x60b5903d.
//
// Solidity: function getPlePerB(uint64 i) view returns(uint256, uint256)
func (_IGroup *IGroupCallerSession) GetPlePerB(i uint64) (*big.Int, *big.Int, error) {
	return _IGroup.Contract.GetPlePerB(&_IGroup.CallOpts, i)
}

// GetRate is a free data retrieval call binding the contract method 0xe795ea8a.
//
// Solidity: function getRate(uint64 i) view returns(uint8, uint8)
func (_IGroup *IGroupCaller) GetRate(opts *bind.CallOpts, i uint64) (uint8, uint8, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getRate", i)

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
func (_IGroup *IGroupSession) GetRate(i uint64) (uint8, uint8, error) {
	return _IGroup.Contract.GetRate(&_IGroup.CallOpts, i)
}

// GetRate is a free data retrieval call binding the contract method 0xe795ea8a.
//
// Solidity: function getRate(uint64 i) view returns(uint8, uint8)
func (_IGroup *IGroupCallerSession) GetRate(i uint64) (uint8, uint8, error) {
	return _IGroup.Contract.GetRate(&_IGroup.CallOpts, i)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroup *IGroupTransactor) Activate(opts *bind.TransactOpts, _gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "activate", _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroup *IGroupSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroup.Contract.Activate(&_IGroup.TransactOpts, _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroup *IGroupTransactorSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroup.Contract.Activate(&_IGroup.TransactOpts, _gi, _kc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kPerB, uint256 _pPerB, bytes _desc) returns()
func (_IGroup *IGroupTransactor) Create(opts *bind.TransactOpts, _level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kPerB *big.Int, _pPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "create", _level, _mr, _tr, _kr, _pr, _kPerB, _pPerB, _desc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kPerB, uint256 _pPerB, bytes _desc) returns()
func (_IGroup *IGroupSession) Create(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kPerB *big.Int, _pPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IGroup.Contract.Create(&_IGroup.TransactOpts, _level, _mr, _tr, _kr, _pr, _kPerB, _pPerB, _desc)
}

// Create is a paid mutator transaction binding the contract method 0xae38ca77.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kPerB, uint256 _pPerB, bytes _desc) returns()
func (_IGroup *IGroupTransactorSession) Create(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kPerB *big.Int, _pPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IGroup.Contract.Create(&_IGroup.TransactOpts, _level, _mr, _tr, _kr, _pr, _kPerB, _pPerB, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_IGroup *IGroupTransactor) SetDesc(opts *bind.TransactOpts, _gi uint64, _desc []byte) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "setDesc", _gi, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_IGroup *IGroupSession) SetDesc(_gi uint64, _desc []byte) (*types.Transaction, error) {
	return _IGroup.Contract.SetDesc(&_IGroup.TransactOpts, _gi, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x773854ed.
//
// Solidity: function setDesc(uint64 _gi, bytes _desc) returns()
func (_IGroup *IGroupTransactorSession) SetDesc(_gi uint64, _desc []byte) (*types.Transaction, error) {
	return _IGroup.Contract.SetDesc(&_IGroup.TransactOpts, _gi, _desc)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_IGroup *IGroupTransactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "setLevel", _gi, _le)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_IGroup *IGroupSession) SetLevel(_gi uint64, _le uint8) (*types.Transaction, error) {
	return _IGroup.Contract.SetLevel(&_IGroup.TransactOpts, _gi, _le)
}

// SetLevel is a paid mutator transaction binding the contract method 0x936945bd.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le) returns()
func (_IGroup *IGroupTransactorSession) SetLevel(_gi uint64, _le uint8) (*types.Transaction, error) {
	return _IGroup.Contract.SetLevel(&_IGroup.TransactOpts, _gi, _le)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_IGroup *IGroupTransactor) SetP(opts *bind.TransactOpts, _gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "setP", _gi, _kp, _pp)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_IGroup *IGroupSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _IGroup.Contract.SetP(&_IGroup.TransactOpts, _gi, _kp, _pp)
}

// SetP is a paid mutator transaction binding the contract method 0x2a0212e5.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp) returns()
func (_IGroup *IGroupTransactorSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int) (*types.Transaction, error) {
	return _IGroup.Contract.SetP(&_IGroup.TransactOpts, _gi, _kp, _pp)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_IGroup *IGroupTransactor) SetPerB(opts *bind.TransactOpts, _gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "setPerB", _gi, _kpPerB, _ppPerB)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_IGroup *IGroupSession) SetPerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _IGroup.Contract.SetPerB(&_IGroup.TransactOpts, _gi, _kpPerB, _ppPerB)
}

// SetPerB is a paid mutator transaction binding the contract method 0xcfb9854f.
//
// Solidity: function setPerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB) returns()
func (_IGroup *IGroupTransactorSession) SetPerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int) (*types.Transaction, error) {
	return _IGroup.Contract.SetPerB(&_IGroup.TransactOpts, _gi, _kpPerB, _ppPerB)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_IGroup *IGroupTransactor) SetRate(opts *bind.TransactOpts, _gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "setRate", _gi, _mr, _tr)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_IGroup *IGroupSession) SetRate(_gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _IGroup.Contract.SetRate(&_IGroup.TransactOpts, _gi, _mr, _tr)
}

// SetRate is a paid mutator transaction binding the contract method 0x1e752182.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr) returns()
func (_IGroup *IGroupTransactorSession) SetRate(_gi uint64, _mr uint8, _tr uint8) (*types.Transaction, error) {
	return _IGroup.Contract.SetRate(&_IGroup.TransactOpts, _gi, _mr, _tr)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_IGroup *IGroupTransactor) SetS(opts *bind.TransactOpts, _gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "setS", _gi, _s, kcnt)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_IGroup *IGroupSession) SetS(_gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _IGroup.Contract.SetS(&_IGroup.TransactOpts, _gi, _s, kcnt)
}

// SetS is a paid mutator transaction binding the contract method 0x5f9dbbf9.
//
// Solidity: function setS(uint64 _gi, uint8 _s, uint8 kcnt) returns()
func (_IGroup *IGroupTransactorSession) SetS(_gi uint64, _s uint8, kcnt uint8) (*types.Transaction, error) {
	return _IGroup.Contract.SetS(&_IGroup.TransactOpts, _gi, _s, kcnt)
}

// IGroupCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the IGroup contract.
type IGroupCreateIterator struct {
	Event *IGroupCreate // Event containing the contract specifics and raw log

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
func (it *IGroupCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGroupCreate)
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
		it.Event = new(IGroupCreate)
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
func (it *IGroupCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGroupCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGroupCreate represents a Create event raised by the IGroup contract.
type IGroupCreate struct {
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
func (_IGroup *IGroupFilterer) FilterCreate(opts *bind.FilterOpts) (*IGroupCreateIterator, error) {

	logs, sub, err := _IGroup.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &IGroupCreateIterator{contract: _IGroup.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x4b14fb8b0a4d249a53bfa9dadcdeab4977149168de6dd5efa8a37e41107b3fda.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 tr, uint256 kr, uint256 pr, uint256 kPerB, uint256 pPerB, bytes _desc)
func (_IGroup *IGroupFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *IGroupCreate) (event.Subscription, error) {

	logs, sub, err := _IGroup.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGroupCreate)
				if err := _IGroup.contract.UnpackLog(event, "Create", log); err != nil {
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
func (_IGroup *IGroupFilterer) ParseCreate(log types.Log) (*IGroupCreate, error) {
	event := new(IGroupCreate)
	if err := _IGroup.contract.UnpackLog(event, "Create", log); err != nil {
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
