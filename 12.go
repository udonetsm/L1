package main

import (
	"log"

	mapset "github.com/deckarep/golang-set"
)

func a() {
	slc := make([]string, 0)
	slc = append(slc, "cat", "dog", "cat", "tree")
	rngvals := make(map[string]struct{}) // Новое множество (либо rngvals := make(map[string]bool)
	for i := 0; i < len(slc); i++ {
		rngvals[slc[i]] = struct{}{} // Создание (либо rngvals[slc[i]]=true)
	}
	for item := range rngvals {
		log.Println(item)
	}
}

func b() {
	slc := make([]string, 0)
	slc = append(slc, "cat", "dog", "cat", "tree")
	set := mapset.NewSet()
	for i := 0; i < len(slc); i++ {
		set.Add(slc[i])
	}
	log.Println(set)
}
