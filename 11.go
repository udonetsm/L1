package main

import "log"

func repeats(vals1, vals2 []int) []int {
	mapvals := make(map[int]int)
	result := make([]int, 0)
	for _, l1 := range vals1 {
		mapvals[l1] += 1
	}
	for _, l2 := range vals2 {
		if mapvals[l2] >= 1 {
			result = append(result, l2)
			mapvals[l2] -= 1
		}
	}
	return result
}

func makeSlice() {
	slc1 := []int{2, 4, 6, 6, 4}
	slc2 := []int{2, 4, 6, 6, 4}
	log.Println(repeats(slc1, slc2))

}
