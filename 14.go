package main

import (
	"fmt"
	"reflect"
)

func getTypeSwitch(val interface{}) string {
	t := reflect.ValueOf(val)
	switch t.Kind() {
	case reflect.Chan:
		return "chan"
	case reflect.Int:
		return "int"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	default:
		return reflect.TypeOf(val).String()
	}
}

func callGetType() {
	ch := make(chan int)
	var str string
	var num int
	var b bool
	fmt.Println(getTypeSwitch(ch), getTypeSwitch(str), getTypeSwitch(num), getTypeSwitch(b))
}
