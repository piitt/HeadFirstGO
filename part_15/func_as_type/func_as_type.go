package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	// в переменной может храниться функция без параметров и возвращаемого значения
	var greeterFunction func()
	// в этой переменной может храниться функция
	// с двумя параметрами int и возвращаемым значением float64
	var mathFunction func(int, int) float64
	greeterFunction = sayHi
	mathFunction = divide
	greeterFunction()
	fmt.Println(mathFunction(5, 2))
}
