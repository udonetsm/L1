package main

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/udonetsm/help/helper"
)

type Counter struct {
	num int64
	sync.RWMutex
}

func New() *Counter {
	defer helper.PanicCapture("new")
	return &Counter{
		num: -1,
	}
}

func (c *Counter) Count(wg *sync.WaitGroup, finish, start int64) int64 {
	defer helper.PanicCapture("Count")
	for i := start; i <= finish; i++ {
		wg.Add(1)
		go func(i int64) {
			c.Lock()
			atomic.AddInt64(&c.num, 1) //or c.num+=1
			c.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return c.num
}

func Counting(start, finish int64) {
	defer helper.PanicCapture("Counting")
	c := New()
	wg := &sync.WaitGroup{}
	log.Println("last counter value is", c.Count(wg, finish, start))
}
