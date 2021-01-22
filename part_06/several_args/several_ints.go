package main

import "fmt"

func severalInt(numbers ...int) {
	fmt.Println(numbers)
}

func main() {
	severalInt(1)
	severalInt(1, 2, 3)
}
