package main

import (
	"sync"
)

type Sums struct {
	sync.RWMutex
	res int
}

func NewSums() *Sums {
	return &Sums{
		res: 0,
	}
}

func (s *Sums) sumSqrts(vals int, wg *sync.WaitGroup) {
	s.RLock()
	defer s.RUnlock()
	s.res += vals
}

func sumSqrts(vals []int) int {
	s := NewSums()
	wg := sync.WaitGroup{}
	vals = sqrt(vals)
	lenvals := len(vals)
	wg.Add(lenvals)
	for i := 0; i < lenvals; i++ {
		go func(i int) {
			s.sumSqrts(vals[i], &wg)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return s.res
}
