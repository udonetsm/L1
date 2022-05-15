package main

import (
	"log"
)

func rotateByMath(num1, num2 int) (int, int) {
	num1 += num2
	num2 = num1 - num2
	num1 -= num2
	return num1, num2
}

func rotateByXor(num1, num2 int) (int, int) {
	num1 ^= num2
	num2 ^= num1
	num1 ^= num2
	return num1, num2
}

func callrotate_interger() {
	log.Println(rotateByMath(4, 10))
	log.Println(rotateByXor(23, 11))
}
