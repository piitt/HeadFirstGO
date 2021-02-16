package main

import "fmt"

func callFunction(passedFunction func()) {
	passedFunction()
}

func callTwice(passsedFunction func()) {
	passsedFunction()
	passsedFunction()
}

func callWithArguments(passsedFunction func(string, bool)) {
	passsedFunction("This sentence is", false)
}

func printReturnValue(passsedFunction func() string) {
	fmt.Println(passsedFunction())
}

func functionA() {
	fmt.Println("function called")
}

func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}

func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func main() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}
