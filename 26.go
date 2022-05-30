package main

import "strings"

func countLetters(str string) (bool, string) {
	// for ignore register
	lower := strings.ToLower(str)
	// new set
	mpbool := make(map[rune]bool)
	for _, item := range lower {
		// try to get value of current letter from map
		_, ok := mpbool[item]
		// if element exists, then it repeats. String broken(false)
		if ok {
			return false, str
		}
		// if doesn't exists, add letter in map
		mpbool[item] = true
	}
	// finally map will contains only unique letters
	return true, str
}

// or we  can check all current letter with each other
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
