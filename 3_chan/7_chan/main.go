package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	val := c.c[key]
	c.mu.Unlock()
	return val
}

func main() {
	key := "test"

	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(c.Value(key))
}
