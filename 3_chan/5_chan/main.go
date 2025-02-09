package main

import (
	"fmt"
	"time"
)

func main() {
	ch_str := make(chan string)
	//ch_int := make(chan int)
	go sayHello(ch_str)

	for i := range ch_str {
		//str := <-ch_str
		fmt.Printf("%s: ", i)
	}
	fmt.Println(<-ch_str)
	//fmt.Println(<-ch_int)
}

func sayHello(exit chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- "stroka"
		//num <- i
	}
	close(exit)
	//close(num)
}
