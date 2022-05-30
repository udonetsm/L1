package main

import (
	"sync"
)

// description object
type Counter struct {
	num int64
	sync.RWMutex
}

// returns new object
func New() *Counter {
	return &Counter{
		num: 0,
	}
}

// cuncurrent increment
func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.num += 1
}

func counting(finish int) int {
	// neew wait for all go routines finished
	wg := &sync.WaitGroup{}
	// make new object
	counter := New()
	for i := 0; i < finish; i++ {
		wg.Add(1)
		//start increment in go routine
		go func(counter *Counter, wg *sync.WaitGroup) {
			counter.Increment()
			wg.Done()
		}(counter, wg)
	}
	// waiting to finish
	wg.Wait()
	return int(counter.num)
}
