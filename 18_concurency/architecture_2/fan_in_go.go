package main

import (
	"fmt"
	"sync"
)

func worker(n int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- n * n
	}()
	return ch
}

func fanIn(channels ...<-chan int) <-chan int {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(in <-chan int) {
			defer wg.Done()
			for i := range in {
				ch <- i
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	var chans []<-chan int
	for i := 0; i < 10; i++ {
		chans = append(chans, worker(i))
	}

	ch := fanIn(chans...)
	for val := range ch {
		fmt.Println(val)
	}
}
