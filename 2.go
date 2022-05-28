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
func (m *Mult) Sqrt(val int) {
	m.Lock()
	defer m.Unlock()
	m.res = append(m.res, val*val)
}

func sqrt(vals []int) []int {
	r := NewMult()
	wg := sync.WaitGroup{}
	lenvals := len(vals)
	wg.Add(lenvals)
	for i := 0; i < lenvals; i++ {
		go func(i int) {
			r.Sqrt(vals[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	return r.res
}
