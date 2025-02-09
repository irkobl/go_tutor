package main

import (
	"fmt"
	"runtime"
	"sync"
)

func processor(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
}

func main() {
	src := make(chan int)
	res := make(chan int)

	n := runtime.NumCPU()

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			processor(src, res)
		}()
	}

	go func() {
		for val := range res {
			fmt.Println(val)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			src <- i
		}
		close(src)
	}()

	wg.Wait()

	close(res)

}
