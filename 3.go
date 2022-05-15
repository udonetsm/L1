package main

import (
	"log"
	"time"
)

func ssquare(num int) int {
	return num * num
}

func ddo() {
	var (
		nums []int
		sum  int
	)
	nums = append(nums, 2, 4, 6, 8, 10)
	for _, num := range nums {
		go func() {
			sum += ssquare(num)
		}()
		time.Sleep(1 * time.Millisecond)
	}
	log.Println(sum)
}
