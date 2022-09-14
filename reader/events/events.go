package events

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/backend/contracts/role"
)

// register account event
type ReAccEvent struct {
	// common.Hash
	Topics []common.Hash
	// common.Address, uint64
	Data []interface{}
}

// parse a register account event from a log
func (e *ReAccEvent) Parse(l types.Log) (*ReAccEvent, error) {
	fmt.Printf("parsing reAccEvent, topic len: %d, data len: %d", len(l.Topics), len(l.Data))

	topics := make([]common.Hash, len(l.Topics))
	data := make([]interface{}, len(l.Data))
	e.Topics = topics
	e.Data = data

	// get role contract abi
	roleABI, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Fatal(err)
	}

	// parse data
	fmt.Println("parsing data...")
	data, err = roleABI.Unpack("ReAcc", l.Data)
	if err != nil {
		return nil, err
	}

	// address and index of account
	addr, ok := data[0].(common.Address)
	if !ok {
		return nil, fmt.Errorf("addr assertion failed")
	}
	fmt.Println("account address: ", addr)

	index, ok := data[1].(uint64)
	if !ok {
		return nil, fmt.Errorf("index assertion failed")
	}
	fmt.Println("registered index: ", index)

	e.Topics = l.Topics
	e.Data = data

	return e, nil
}
