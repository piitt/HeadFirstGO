package main

import "fmt"

func main() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)

	fmt.Println("----------------")

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	fmt.Println("---------------")

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
}
