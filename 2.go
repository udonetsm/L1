package main

import (
	"sync"
)

// struc of result counting
type Mult struct {
	sync.Mutex
	res []int
}

// returns new safe storage
func NewMult() *Mult {
	m := make([]int, 0)
	return &Mult{
		res: m,
	}
}

//safe find of sqrt
func (m *Mult) Sqrt(val int) {
	m.Lock()
	m.res = append(m.res, val*val)
	m.Unlock()
}

// it create new storage and calls all function
func sqrt(vals []int) []int {
	r := NewMult()
	wg := sync.WaitGroup{}
	lenvals := len(vals)
	wg.Add(lenvals)
	for i := 0; i < lenvals; i++ {
		go func(i int) {
			r.Sqrt(vals[i])
			// decrement of amount
			wg.Done()
		}(i)
	}
	// waiting for finish goroutines
	wg.Wait()
	return r.res
}
