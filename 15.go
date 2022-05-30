package main

var justString string

func createHugeString(i int) (str string) {
	return
}

func someFunc(val string) string {
	v := createHugeString(1 << 10)
	out := 100
	if out > len(v) {
		return v
	}
	justString = v[:100]
	return justString
}

func mmain() {
	someFunc(justString)
}
