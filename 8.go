package main

import (
	"fmt"
	"log"
	"strconv"
)

func IntToBits(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

func bitsToInt(s string) int64 {
	newVal, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return newVal
}

func changeOneBit(num, targetBit, newBit int) int64 {
	if targetBit > 0 && num == 0 {
		return 0
	}

	if newBit > 1 || newBit < 0 {
		log.Println("Новый бит может быть только 1 или 0")
		return 0
	}

	bits := []byte(IntToBits(num))
	fmt.Printf("Было \t %s(%v)\n", string(bits), num)

	if targetBit > len(bits) {
		log.Println("Целевой бит вышел за рамки двоичного числа")
		return 0
	}

	newB := []byte(IntToBits(newBit))
	bits[targetBit] = newB[0]
	newVal := bitsToInt(string(bits))
	fmt.Printf("Стало \t %s(%v)\n", string(bits), newVal)
	return newVal
}
