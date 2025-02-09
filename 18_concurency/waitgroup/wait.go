package main

import (
	"fmt"
	"sync"
)

func printN(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n)
}

func step1(ch chan<- string) {
	//.....
	ch <- "message from Sleep 1"
}

func step2(ch <-chan string) {
	fmt.Println(<-ch)
	//...
}

func main() {
	const N = 10
	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		go printN(i, &wg)
	}

	wg.Wait()

	ch := make(chan string)
	go step1(ch)
	step2(ch)

}
