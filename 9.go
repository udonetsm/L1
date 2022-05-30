package main

// get all values and insert it to channel
func writeX(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range values {
			ch <- val
		}
		// close channel when values finished
		close(ch)
	}()
	return ch
}

// returns multiplicated values from origin chan to pout chan
func multX(ch <-chan int) <-chan int {
	// create channel for capture written values
	out := make(chan int)
	go func() {
		for writen := range ch {
			out <- writen * 2
		}
		// close channel when values finished
		close(out)
	}()
	return out
}

func callWriteAndSum() {
	// make origin channel
	vals := writeX(12, 43, 65, 12, 444, 23)
	// make out channel from origin channel
	out := multX(vals)
	for val := range out {
		println(val)
	}
}
