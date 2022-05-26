package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"
	"time"
)

const ExitByUser = 900

func Worker(ch chan int64, ctx context.Context, id int) {
	time.Sleep(time.Second * 3)
	for {
		select {
		case <-ctx.Done():
			data, ok := <-ch
			if !ok {
				return
			}
			println("worker", id, "saved", data)
		default:
			println("worker", id, "captured", <-ch)
		}
	}
}

func rW() int {
	var amount int
	flag.IntVar(&amount, "a", 2, "set amount of active workers")
	flag.Parse()
	ch := make(chan int64)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	for i := 0; i < amount; i++ {
		go Worker(ch, ctx, i)
	}
	for {
		select {
		case <-ctx.Done():
			cancel()
			close(ch)
			log.Println("CANCEL")
			return ExitByUser
		default:
			ch <- time.Now().UnixMilli()
			time.Sleep(time.Millisecond)
			ch <- 1243
			ch <- 89898
		}
	}
}
