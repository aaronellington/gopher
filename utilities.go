package gopher

import "sync"

func Runners(maxRunners uint, items []any, worker func(i int, x any)) {
	wg := sync.WaitGroup{}
	wg.Add(len(items))

	guard := make(chan int, maxRunners)
	for i, item := range items {
		guard <- 1 // would block if guard channel is already filled
		go func(n int, t any) {
			worker(n, t)
			wg.Done()
			<-guard
		}(i, item)
	}

	wg.Wait()
}
