package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/rockiecn/backend/contracts/role"
)

func main() {

	// dial block chain
	client, err := ethclient.Dial("https://testchain.metamemo.one:24180")
	if err != nil {
		log.Fatal(err)
	}

	// role address
	roleAddress := common.HexToAddress("0xC6fE0358b6d8Cb064a0a7db15dbdb7D80F02bD4E")

	// create a query with contract address
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(8228622),
		Addresses: []common.Address{
			roleAddress,
		},
	}

	// filter logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	println("logs number of role contract:", len(logs))

	// get role contract abi
	roleABI, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Fatal(err)
	}

	// event struct
	type eventStruct struct {
		topics []interface{}
		data   []interface{}
	}

	// create map structure
	reAccMap := make(map[common.Hash][]eventStruct)

	// scan each log
	fmt.Println("Travelling logs..")
	for i, vLog := range logs {

		// check log hash, should be ReAcc's hash
		// log Hash: 0xe22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752
		byteHash, err := hex.DecodeString("e22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752")
		if err != nil {
			log.Fatal(err)
		}
		var recTopic common.Hash
		recTopic.SetBytes(byteHash)

		// skip other events
		if vLog.Topics[0] != recTopic {
			continue
		}

		fmt.Println("-- log index:", i)
		fmt.Println("-- block number:", vLog.BlockNumber) // 2490757

		// parse topics, 1 topic in all, the hash of the event
		fmt.Println("parsing ReAcc event..")
		topics := make([]string, 1)
		for t := range vLog.Topics {
			topics[t] = vLog.Topics[t].Hex()
			fmt.Printf("topic %d: %s\n", t, topics[t])
		}

		// parse data
		fmt.Println("parsing data..")
		data, err := roleABI.Unpack("ReAcc", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		// address and index of account
		addr, ok := data[0].(common.Address)
		if !ok {
			log.Fatal("addr assertion failed")
		}
		fmt.Println("address: ", addr)

		index, ok := data[1].(uint64)
		if !ok {
			log.Fatal("index assertion failed")
		}
		fmt.Println("index: ", index)

		// record event info into map
		fmt.Println("recording event..")
		var newEvent eventStruct
		newEvent.topics = append(newEvent.topics, topics[0])
		newEvent.data = append(newEvent.data, data[0])
		newEvent.data = append(newEvent.data, data[1])
		reAccMap[recTopic] = append(reAccMap[recTopic], newEvent)
		fmt.Println("recording event ok")

		fmt.Println()
	}

	// eventSignature := []byte("Transfer(address indexed, address indexed, uint256)")
	// hash := crypto.Keccak256Hash(eventSignature)
	// fmt.Println("event hash:", hash.Hex()) // 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
}
