package main

import "log"

func repeats(vals1, vals2 []int) []int {
	// new map
	mapvals := make(map[int]int)
	result := make([]int, 0)
	for _, l1 := range vals1 {
		// insert into map all elements of array
		mapvals[l1] += 1
	}
	for _, l2 := range vals2 {
		// insert repeat of element in map to result slice
		if mapvals[l2] >= 1 {
			result = append(result, l2)
			// remove accounted element from range
			mapvals[l2] -= 1
		}
	}
	return result
}

func makeSlice() {
	slc1 := []int{2, 4, 6, 6, 4, 1, 2, 3}
	slc2 := []int{2, 4, 6, 6, 4, 1, 2}
	log.Println(repeats(slc1, slc2))

}
