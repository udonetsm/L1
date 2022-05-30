package main

import (
	"sync"
)

// description of object
type Mapping struct {
	sync.Mutex
	mp map[int]int
}

// return new object
func Init() *Mapping {
	mp := make(map[int]int)
	return &Mapping{
		mp: mp,
	}
}

//insert into
func (m *Mapping) Set(key, value int) {
	m.Lock()
	m.mp[key] = value
	m.Unlock()
}

func make_slice_and_insert_into_map() map[int]int {
	slc := []int{333333, 544545, 65765757, 23423423, 6567567456}
	// need to wait for inserting finished for all items of map
	wg := sync.WaitGroup{}
	mp := Init()
	lenslc := len(slc)
	for i := 0; i < lenslc; i++ {
		wg.Add(1)
		go func(i int) {
			mp.Set(i, slc[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	return mp.mp
}
