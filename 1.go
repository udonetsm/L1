// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).
package main

import (
	"log"
)

type Human struct {
	Name string
	Age  int
	Sex  string
}

type Action struct {
	Some string
	Human
}

// human walk <s> steps
func (h *Human) Walk(s int) int {
	return s
}

// human sleeps <hrs> hours
func (h *Human) Sleep(hrs int) int {
	return hrs
}

// human say something <words>
func (h *Human) Say(words string) string {
	return words
}

// human returns to <job> job
func (h *Human) ToJob(job string) string {
	return job
}

// use all methods from Human for object Action
func life(steps, sleep_hours int, saying, job string) {
	act := new(Action) //initialize Action object
	act.Some = "using methods from Human for object Action"
	log.Printf("Human walk %v steps, lay down to sleep for "+
		"%v hours and say %s when wake up and go to job to %s"+
		" <-- %s",
		act.Walk(steps), act.Sleep(sleep_hours),
		act.Say(saying), act.ToJob(job), act.Some)
}
