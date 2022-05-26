package main

import (
	"sync"
)

type Mult struct {
	sync.RWMutex
	res []int
}

func NewMult() *Mult {
	m := make([]int, 0)
	return &Mult{
		res: m,
	}
}
func (m *Mult) Sqrt(val int, wg *sync.WaitGroup) {
	m.Lock()
	defer m.Unlock()
	m.res = append(m.res, val*val)
	wg.Done()
}

func sqrt(vals []int) []int {
	r := NewMult()
	wg := sync.WaitGroup{}
	lenvals := len(vals)
	wg.Add(lenvals)
	for i := 0; i < lenvals; i++ {
		go r.Sqrt(vals[i], &wg)
	}
	wg.Wait()
	return r.res
}
