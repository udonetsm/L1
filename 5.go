package main

import (
	"log"
	"os"
	"time"
)

func insert_in_channel(interval int, ch chan int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for sec := range ticker.C {
		ch <- sec.Second()
	}
}

func read_channel(timesec int) {
	ch := make(chan int)
	timer := time.NewTimer(time.Duration(timesec) * time.Second)
	go insert_in_channel(1, ch)
	for sec := range ch {
		select {
		case <-timer.C:
			os.Exit(0)
		default:
			log.Println(sec)
		}
	}
}
