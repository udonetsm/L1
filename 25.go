package main

import (
	"log"
	"time"
)

// unblock after d time
func sleep(d time.Duration) {
	<-time.After(d)
}

// unblock after d ticks
func sleep2(d time.Duration) time.Time {
	ticker := time.Tick(d)
	for done := range ticker {
		return done
	}
	return time.Now()
}

func CallSleep() {
	log.Println("Do something")
	sleep(5 * time.Second)
	log.Println("Something else") //out after 5 sec
}
