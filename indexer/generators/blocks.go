package indexer

// Generates a channel withing a range
func GenBlockRange(start uint64, end uint64) <-chan uint64 {
	out := make(chan uint64)

	go func() {
		for i := start; i <= end; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}
