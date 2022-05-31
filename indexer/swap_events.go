package indexer

import (
	"sync"

	"indexer/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func SwapTokenExchange(new_swap *contracts.NewSwap, blocks <-chan uint64) <-chan contracts.NewSwapTokenExchange {
	out := make(chan contracts.NewSwapTokenExchange)
	go func() {
		for block := range blocks {
			opts := bind.FilterOpts{
				Start: block,
				End:   &block,
			}

			event, _ := new_swap.FilterTokenExchange(&opts, []common.Address{})

			for event.Next() {
				out <- *event.Event
			}
		}
		close(out)
	}()
	return out
}

func MergeSwapTokenExchange(cs ...<-chan contracts.NewSwapTokenExchange) <-chan contracts.NewSwapTokenExchange {
	var wg sync.WaitGroup
	out := make(chan contracts.NewSwapTokenExchange)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan contracts.NewSwapTokenExchange) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
