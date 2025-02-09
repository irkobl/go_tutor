package main

import (
	"sync"
	"time"
)

type Counter struct {
	mu sync.RWMutex
	c  map[string]int
}

func (c *Counter) CountMe() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *Counter) CountMeAgain() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c
}

func main() {
	//key := "test"

	//c := Counter{c: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		//go c.Inc(key)
	}

	time.Sleep(time.Second * 3)
	//fmt.Println(c.Value(key))
}
