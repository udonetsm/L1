package main

import (
	"sync"
)

// descroption of object
type Sums struct {
	sync.Mutex
	res int
}

// gives new object
func NewSums() *Sums {
	return &Sums{
		res: 0,
	}
}

// sums all sqrts of object safety
func (s *Sums) sumSqrts(vals int) {
	s.Lock()
	defer s.Unlock()
	s.res += vals
}

// call all functions
func sumSqrts(vals []int) int {
	s := NewSums()
	// we should wait for sum all
	wg := sync.WaitGroup{}
	vals = sqrt(vals)
	lenvals := len(vals)
	wg.Add(lenvals)
	for i := 0; i < lenvals; i++ {
		go func(i int) {
			s.sumSqrts(vals[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	return s.res
}
