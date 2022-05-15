package main

import "log"

func QuickSort(nums []int) []int {
	// find len of arr
	lennums := len(nums)
	// if arr contains less then 2 items, nothing to sort
	if lennums < 2 {
		return nums
	}
	toehold := nums[0] //or nums[lennums/2]
	// items which less than toehold
	left := make([]int, 0)
	// items which more than toehold
	right := make([]int, 0)
	for _, num := range nums[1:] {
		// move to right if item more than toehold
		if num > toehold {
			right = append(right, num)
		} else {
			// move to left if item less than toehold
			left = append(left, num)
		}
	}
	// recoursive sorting left side
	nums = append(QuickSort(left), toehold)
	// recoursive sorting right side
	nums = append(nums, QuickSort(right)...)
	return nums
}

func make_arr_and_sort() {
	var nums []int
	nums = append(nums, 43, 1233, 321, 998, 22200399, 1, 3, 2, 6, 5, 4, 0, 9, 7, 12)
	log.Println(QuickSort(nums))
}
