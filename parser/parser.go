package parser

import "github.com/ethereum/go-ethereum/common"

type Parser struct {
	endpoint string
	// instance contract address
	instanceAddr string
	// record all contracts' address
	allAddr map[string]string
	// record all events according event hash
	allEvent map[common.Hash][]interface{}
}
