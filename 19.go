package main

func reverseLettersInString(str string) (newrune []rune) {
	data := []rune(str)
	for i := len(data) - 1; i >= 0; i-- {
		newrune = append(newrune, data[i])
	}
	return newrune
}
