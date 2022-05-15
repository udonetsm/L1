package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func Work(data []int, lendata int, ctx context.Context) {
	for i := 0; i < lendata; i++ {
		log.Println(data[i])
		for {
			select {
			case <-ctx.Done():
				return
			default:

			}
		}
	}
}

func makechan() {
	slc := make([]int, 0)
	slc = append(slc, 12, 32, 65, 34, 54, 87)
	lendata := len(slc)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			Work(slc, lendata, ctx)
		}
	}
}
