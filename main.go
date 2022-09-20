package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/rockiecn/backend/reader"
)

func main() {

	// new reader
	p, err := reader.New("https://testchain.metamemo.one:24180", big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	// scan blocks
	err = p.FilterBlock(
		big.NewInt(3104474),
		//"0xC6fE0358b6d8Cb064a0a7db15dbdb7D80F02bD4E"
		"")
	if err != nil {
		log.Fatal(err)
	}

	//
	fmt.Println("--- show events")
	p.ShowEvents("e22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752")

}
