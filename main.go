//go:generate abigen --abi ./abis/new_swap.json --pkg contracts --type NewSwap --out contracts/new_swap.go

package main

import (
	"fmt"
	"log"

	"indexer/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://node.koyo.finance/rpc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	new_swap, err := contracts.NewNewSwap(common.HexToAddress("0x9F0a572be1Fcfe96E94C0a730C5F4bc2993fe3F6"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a NewSwap contract: %v", err)
	}
	a, err := new_swap.A(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve swap amplification name: %v", err)
	}
	fmt.Println("Amplification:", a)
}
