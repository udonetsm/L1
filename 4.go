package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const ClosedByUser = 110

//reads data channel and Done channel of context
func readFromCh(ch chan int64, ctx context.Context, id int, wg *sync.WaitGroup) {
	for {
		// we should wait to finish all go routines
		wg.Add(1)
		select {
		// if done captured, need check data chan
		case <-ctx.Done():
			data, open := <-ch
			// if channel closed, stop reading
			if !open {
				wg.Done()
				return
			} else {
				// if not closed, check value in channel and out it
				println("saved", data, "by", id)
				wg.Done()
				return
			}
		default:
			for data := range ch {
				// out values from channel while done not captured
				println(data, "by", id)
			}
			wg.Done()
		}
	}
}

func writeInChannel(ch chan int64, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// if done signal captured, close channel to write inside noting
			close(ch)
			return
		default:
			// while done not captured, write in channel
			a := time.Now().UnixMilli()
			ch <- a
			println("wrote", a)
			time.Sleep(time.Millisecond)
		}
	}
}

func callReadworkers() {
	ch := make(chan int64)
	amnt := amount()
	// ctx will listen done channel
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	wg := &sync.WaitGroup{}
	defer cancel()
	for i := 0; i < amnt; i++ {
		// start amount of workers
		go readFromCh(ch, ctx, i, wg)
	}
	writeInChannel(ch, ctx)
	wg.Wait()
	log.Println("END")
}

func amount() (amnt int) {
	// get amount of worker as flag
	flag.IntVar(&amnt, "a", 2, "set amount of workers")
	flag.Parse()
	return
}
