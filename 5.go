package main

import (
	"log"
	"time"
)

func rw(sec int) {
	ch := make(chan int64)
	for i := 0; i < sec; i++ {
		go func() {
			// read from channel permanently
			log.Println(<-ch)
		}()
		// write in channel
		ch <- time.Now().Unix()
		time.Sleep(time.Second)
	}
	close(ch)
}
