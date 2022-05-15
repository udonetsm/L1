package main

import (
	"fmt"
	"time"
)

func square(num int) {
	fmt.Println(num * num)
	/* other way to write to stdout*/
	//fmt.Fprintln(os.Stdout, num*num)
}

func do() {
	nums := [...]int{2, 4, 6, 8, 10}
	for _, num := range nums {
		go square(num)
	}
	// after goroutines called, compiler want to swith to next line. There is end of the programm
	// for hold closing, programm do callback of scanln, to show results
	//fmt.Scanln()
	// or
	time.Sleep(1 * time.Millisecond)
}
