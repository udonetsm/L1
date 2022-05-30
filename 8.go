package main

import (
	"fmt"
	"strconv"
)

// create mask mask and shift 1 to target position
func changeBitTo1(val, position int) int {
	mask := 1 << position
	// change target bit
	return val | mask
}

// create mask, shift 1 to target position and invert 1 to 0
func changeBitTo0(val, position int) int {
	mask := ^(1 << position)
	// change target bit to 0
	return val & mask
}

// if we try to change 63 bit, number will be incorrect(max value for type not negative number).
// Then need to invert number and increment once
func negative(val int) int {
	return ^val + 1
}

func changeBit() (result int) {
	var value, position, oneorzero int
	fmt.Println("type: <value> <position to change> <1 or 0>")
	fmt.Scanln(&value, &position, &oneorzero)
	println(strconv.FormatInt(int64(value), 2))
	if oneorzero == 1 && position == 63 {
		result = negative(value)
	}
	if oneorzero != 0 && oneorzero != 1 {
		result = 0
	}
	if oneorzero == 1 && position != 63 {
		result = changeBitTo1(value, position)
	}
	if oneorzero == 0 {
		result = changeBitTo0(value, position)
	}
	println(strconv.FormatInt(int64(result), 2))
	return
}
