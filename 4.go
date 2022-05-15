package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const PROGRAMM_STOPPED_BY_USER = 100

func WriteInChannel(amountWorkers int) {
	ctx := context.Background()
	sigs := make(chan os.Signal, 1)
	ch := make(chan int64, amountWorkers)
	for {
		for i := 0; i < amountWorkers; i++ {
			ch <- time.Now().UnixNano()
			go Worker(i, ch, ctx)
			go captureSig(sigs)
		}
		time.Sleep(2 * time.Second)
	}
}

func Worker(id int, ch <-chan int64, ctx context.Context) {
	// reading channel and output data from it
	for data := range ch {
		fmt.Fprintln(os.Stdout, data, "(Worker ", id, ")")
	}
}

func captureSig(sigs chan os.Signal) {
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	log.Println("ALLDONE")
}

func write_and_read() {
	var amountOfWorkers int
	flag.IntVar(&amountOfWorkers, "a", 2, "-a <amount of workers>")
	flag.Parse()
	WriteInChannel(amountOfWorkers)
}
