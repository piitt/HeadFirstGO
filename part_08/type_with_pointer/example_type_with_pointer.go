package main

import "fmt"

type myType struct {
	myField int
}

func main() {
	var value myType
	value.myField = 3
	var pointer *myType = &value
	fmt.Println((*pointer).myField)
	(*pointer).myField = 9
	fmt.Println((*pointer).myField)
}
