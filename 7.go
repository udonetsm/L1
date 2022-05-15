package main

import (
	"sync"
)

type Mapping struct {
	sync.RWMutex
	mp map[int]int
}

func Init() *Mapping {
	mp := make(map[int]int)
	return &Mapping{
		mp: mp,
	}
}

func (m *Mapping) Set(data []int, wg *sync.WaitGroup) {
	lendata := len(data)
	wg.Add(lendata)
	for i := 0; i < lendata; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			m.mp[i] = data[i]
			wg.Done()
		}(i)
	}
}

func make_slice_and_insert_into_map() map[int]int {
	slc := make([]int, 0)
	slc = append(slc, 4333333, 544545, 65765757, 23423423, 6567567456)
	wg := sync.WaitGroup{}
	mp := Init()
	mp.Set(slc, &wg)
	wg.Wait()
	return mp.mp
}
