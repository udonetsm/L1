package main

import "log"

// cover target element by all elements(shift elements to left side)
func deleteIFromSlice(slc []int, target int) []int {
	slc = append(slc[0:target], slc[target+1:]...)
	return slc
}

// swap last item  with target item and cut length of array(-1)
func delete2(slc []int, target int) []int {
	lenslc := len(slc)
	slc[target] = slc[lenslc-1]
	lenslc -= 1
	slc = slc[:lenslc-1]
	return slc
}

func ddeleteFromSlice() {
	slc := []int{2, 7, 4, 5, 9, 1}
	log.Println(deleteIFromSlice(slc, 3))
	log.Println(delete2(slc, 0))
}
