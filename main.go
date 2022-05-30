//go:generate abigen --abi ./abis/new_swap.json --pkg contracts --type NewSwap --out contracts/new_swap.go

package main

import (
	"fmt"
	"log"

	"indexer/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	pool4Address := common.HexToAddress("0x9F0a572be1Fcfe96E94C0a730C5F4bc2993fe3F6")
	filterBlock := uint64(632160)

	client, err := ethclient.Dial("https://node.koyo.finance/rpc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	new_swap, err := contracts.NewNewSwap(pool4Address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a NewSwap contract: %v", err)
	}

	opts := bind.FilterOpts{
		Start: filterBlock,
		End:   &filterBlock,
	}

	event, err := new_swap.FilterTokenExchange(&opts, []common.Address{})

	for event.Next() {
		fmt.Println(event.Event.Raw.BlockHash.Hex())
		fmt.Println(event.Event.Raw.BlockNumber)
		fmt.Println(event.Event.Raw.TxHash.Hex())

		fmt.Println(string(event.Event.BoughtId.String()))
	}

}
