package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	fmt.Println("-------------------")

	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)
	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt))

	fmt.Println("--------------------")

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(&myFloat)
	fmt.Println(reflect.TypeOf(&myFloat))

	fmt.Println("-------------------")

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(&myBool)
	fmt.Println(reflect.TypeOf(&myBool))

}
