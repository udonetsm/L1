package main

import "log"

func deleteIFromSlice(slc []int, target int) []int {
	slc = append(slc[0:target], slc[target+1:]...)
	return slc
}

func ddeleteFromSlice() {
	slc := []int{2, 7, 4, 5, 9, 1}
	log.Println(deleteIFromSlice(slc, 3))
}
