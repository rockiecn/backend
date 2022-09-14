package reader

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Reader struct {
	endpoint string

	// parsed block number
	currentBlock *big.Int

	// record all contracts' address
	allAddrs map[string]common.Address

	// record all events according event hash
	allEvents map[common.Hash][]types.Log
}

// create a new reader
func New(endpoint string, currentBlock *big.Int) (*Reader, error) {
	allAddrs := make(map[string]common.Address)
	//allAddrs["role"] = common.HexToAddress("0xC6fE0358b6d8Cb064a0a7db15dbdb7D80F02bD4E")

	allEvents := make(map[common.Hash][]types.Log)
	// reAcc := common.HexToHash("e22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752")
	// allEvents[reAcc] = nil

	reader := &Reader{
		endpoint:     endpoint,
		currentBlock: currentBlock,
		allAddrs:     allAddrs,
		allEvents:    allEvents,
	}

	return reader, nil
}

// filter a block, record all events of the block
func (reader *Reader) FilterBlock(block *big.Int, contractAddr string) error {

	// dial block chain
	client, err := ethclient.Dial(reader.endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("query, block: %s, address: %s\n", block.String(), contractAddr)
	query := ethereum.FilterQuery{
		FromBlock: block,
		ToBlock:   block.Add(block, big.NewInt(1)),
		Addresses: []common.Address{
			common.HexToAddress(contractAddr),
		},
	}

	// filter logs with contract address and block number
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	if len(logs) == 0 {
		return nil
	}

	fmt.Printf("=> log number of block %d: %d\n", block, len(logs))

	// scan each log

	fmt.Println("recording logs..")
	for _, v := range logs {
		// record a log
		h := v.Topics[0]
		reader.allEvents[h] = append(reader.allEvents[h], v)
	}

	return nil
}

// show all events of a log
func (reader *Reader) ShowEvents(logHash string) error {
	h, err := reader.String2Hash(logHash)
	if err != nil {
		return err
	}

	// show all events
	for i, v := range reader.allEvents[h] {
		if v.Topics[0] == h {
			fmt.Printf("event %d: %v\n", i, v)
		}
	}

	return nil
}

func (reader *Reader) String2Hash(str string) (h common.Hash, err error) {
	// check log hash, should be ReAcc's hash
	byteHash, err := hex.DecodeString(str)
	if err != nil {
		return common.Hash{}, err
	}

	// generate hash
	h.SetBytes(byteHash)
	return h, nil
}
