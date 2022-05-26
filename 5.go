package main

import (
	"log"
	"time"
)

func rw(sec int) {
	ch := make(chan int64)
	for i := 0; i < sec; i++ {
		go func() {
			ch <- time.Now().Unix()
		}()
		go func() {
			log.Println(<-ch)
		}()
		time.Sleep(time.Second)
	}
	close(ch)
}
