package main

import (
	"log"
	"reflect"
)

func gettypeReflect(num interface{}) {
	log.Println(reflect.TypeOf(num).Kind())
}
