package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	done := make(chan bool)

	for i := 0; i < 30000; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("the only once")
			})
			done <- true
		}()
	}

	for i := 0; i < 30; i++ {
		<-done
	}
}
