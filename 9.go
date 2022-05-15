package main

import (
	"fmt"
	"os"
)

func getX(nums []int) <-chan int {
	och := make(chan int)
	go func() {
		for _, num := range nums {
			och <- num
		}
		close(och)
	}()
	return och
}

func sum(nums <-chan int) <-chan int {
	och := make(chan int)
	go func() {
		for num := range nums {
			num *= 2
			och <- num
		}
		close(och)
	}()
	return och
}

func getx_and_sum() {
	origin := [...]int{2, 1, 13, 43, 23}
	nums := getX(origin[0:])
	result := sum(nums)
	for sums := range result {
		fmt.Fprintln(os.Stdout, sums)
	}
}
