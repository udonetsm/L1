package main

import "strings"

func reverseWordsInString(data string) (newstring string) {
	slc := strings.Split(data, " ")
	lenslc := len(slc)
	for i := lenslc - 1; i >= 0; i-- {
		newstring += slc[i] + " "
	}
	return
}
