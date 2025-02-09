package main

import (
	"fmt"
	"time"
)

func error1() {
	for i := 0; i < 10; i++ {
		// i := i
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func erro2() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func error3() {
	f := func(i int) {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		go f(i)
	}
}

func main() {
	// error1()
	// erro2()
	error3()
}
