package main

import (
	"log"
	"reflect"
)

func gettypeReflect(val interface{}) {
	log.Println(reflect.TypeOf(val).Kind())
}

func getTypeSwitch(val interface{}) {
	switch val.(type) {
	case int:
		log.Println("chan")
		// do something
	default:
	}
}
