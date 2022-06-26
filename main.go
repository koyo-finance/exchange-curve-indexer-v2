//go:generate abigen --abi ./abis/new_swap.json --pkg contracts --type NewSwap --out contracts/new_swap.go

package main

import (
	"fmt"
	"log"
	"time"

	"indexer/contracts"
	indexer "indexer/indexer"
	generators "indexer/indexer/generators"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	start := time.Now()

	pool4Address := common.HexToAddress("0x9F0a572be1Fcfe96E94C0a730C5F4bc2993fe3F6")

	client, err := ethclient.Dial("https://node.koyo.finance/rpc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	new_swap, err := contracts.NewNewSwap(pool4Address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a NewSwap contract: %v", err)
	}

	// header, err := client.HeaderByNumber(context.Background(), nil)

	// for {
	// 	for {
	// 		header, err = client.HeaderByNumber(context.Background(), nil)
	// 		if err == nil {
	// 			break
	// 		} else {
	// 			time.Sleep(10 * time.Second)
	// 		}
	// 	}

	// 	fmt.Println(header.Number.String())
	// }

	// blocks := generators.GenBlockRange(629170, 632161)
	blocks := generators.GenBlockRange(560609, 632161)

	var swapTokenExchangeChannels [256]<-chan contracts.NewSwapTokenExchange
	for i := range swapTokenExchangeChannels {
		swapTokenExchangeChannels[i] = indexer.SwapTokenExchange(new_swap, blocks)
	}

	for n := range indexer.MergeSwapTokenExchange(swapTokenExchangeChannels[:]...) {
		fmt.Println(n.BoughtId.String())
	}

	elapsed := time.Since(start)
	log.Printf("%d blocks took %s", 632161-560609, elapsed)
}
