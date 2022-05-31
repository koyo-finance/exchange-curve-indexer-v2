//go:generate abigen --abi ./abis/new_swap.json --pkg contracts --type NewSwap --out contracts/new_swap.go

package main

import (
	"fmt"
	"log"

	"indexer/contracts"
	indexer "indexer/indexer"
	generators "indexer/indexer/generators"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
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

	blocks := generators.GenBlockRange(629170, 632161)

	c1 := indexer.SwapTokenExchange(new_swap, blocks)
	c2 := indexer.SwapTokenExchange(new_swap, blocks)
	c3 := indexer.SwapTokenExchange(new_swap, blocks)
	c4 := indexer.SwapTokenExchange(new_swap, blocks)
	c5 := indexer.SwapTokenExchange(new_swap, blocks)
	c6 := indexer.SwapTokenExchange(new_swap, blocks)
	c7 := indexer.SwapTokenExchange(new_swap, blocks)
	c8 := indexer.SwapTokenExchange(new_swap, blocks)

	// Consume the merged output from c1 and c2.
	for n := range indexer.MergeSwapTokenExchange(c1, c2, c3, c4, c5, c6, c7, c8) {
		fmt.Println(n.BoughtId.String())
	}

}
