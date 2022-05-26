package main

import (
	"context"
	"log"
	"time"
)

func workBeforeStopping1(quit chan bool) {
	for {
		select {
		case <-quit:
			// close quit channel, when it get true(block goroutine)
			close(quit)
			println("Closed")
			// finish goroutine
			return
		default:
			// do something when goroutine doesn't get true in channel
			log.Println("Performs")
			time.Sleep(time.Second)
		}
	}
}

func workBeforeStopping2(ctx context.Context) {
	for {
		select {
		// when ctx get Done signal, then goroutine stop
		case <-ctx.Done():
			return
		default:
			// when goroutine doesn't get Done signal, then do some job
			println("Perform")
			time.Sleep(time.Second)
		}
	}
}

func stoppingGoroutine1() {
	quit := make(chan bool)
	go workBeforeStopping1(quit)
	time.Sleep(time.Second * 5)
	quit <- true //stop goroutine by channel with passing true in channel.
}

func stoppingGoroutine2() {
	ctx := context.Background()
	go workBeforeStopping2(ctx)
	time.Sleep(time.Second * 5)
	ctx.Done() //stop goroutine by context with passing Done signal
}
