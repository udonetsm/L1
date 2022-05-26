package main

import "strings"

func countLetters(str string) (bool, string) {
	lower := strings.ToLower(str)
	mpbool := make(map[rune]bool, 0)
	for _, item := range lower {
		_, ok := mpbool[item]
		if ok {
			return false, str
		}
		mpbool[item] = true
	}
	return true, str
}

func countLetters1(str string) (string, bool) {
	lower := strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if lower[i] == lower[j] {
				return str, false
			}
		}
	}
	return str, true
}
