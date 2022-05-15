package main

import (
	"strings"
)

func reverseLettersInString(data string) (newstr string) {
	slc := strings.Split(data, "")
	lenslc := len(slc)
	for i := lenslc - 1; i >= 0; i-- {
		newstr += slc[i]
	}
	return newstr
}
