package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var counter int

func inc(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10_000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go inc(&wg)
	go inc(&wg)

	wg.Wait()

	fmt.Println(counter)
}
