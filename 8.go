package main

func changeBitTo1(val, position int) int {
	mask := 1 << position
	return val | mask
}

func changeBitTo0(val, position int) int {
	mask := ^(1 << position)
	return val & mask
}
