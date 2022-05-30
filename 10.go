package main

import (
	"log"
	"math"
	"sync"
)

// description of object
type Set struct {
	sync.RWMutex
	subs map[int]*SubSet
}

type SubSet struct {
	set map[float64]int
}

// return new storage(set)
func NewSet() *Set {
	return &Set{
		subs: map[int]*SubSet{},
	}
}

// returns name of subset from step of temperature(10)
func nameSubset(val float64) (name int, value float64) {
	name = int(math.Trunc(val/10) * 10)
	value = val
	return
}

//add temperature to target subset and subset to target set
func (s *Set) Insert(val float64) {
	name, v := nameSubset(val)
	item, ok := s.subs[name]
	if !ok {
		s.Lock()
		defer s.Unlock()
		s.subs[name] = &SubSet{
			set: map[float64]int{v: 1},
		}
	} else {
		item.set[v]++
	}
}

// returns set by target steps
func (s *Set) Get(SubNames ...int) (vals []*SubSet, subnames []int) {
	for _, SubName := range SubNames {
		s.RLock()
		defer s.RUnlock()
		val, ok := s.subs[SubName]
		if ok {
			vals = append(vals, val)
			subnames = append(subnames, SubName)
		}
	}
	return vals, subnames
}

func celcions(subsequence []float64) {
	set := NewSet()
	for _, item := range subsequence {
		set.Insert(item)
	}
	/* for i, items := range set.subs {
		log.Println(items.set)
	} */
	// or we can get some range of temperature
	val, names := set.Get(20, -20, 10, 30)
	for i, item := range val {
		log.Printf("%v{%v}", names[i], item.set)
	}
}

func celc() {
	cels := make([]float64, 0)
	cels = append(cels, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5)
	celcions(cels)
}
