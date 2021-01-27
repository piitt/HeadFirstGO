package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	subscriber1 := subscriber{name: "Aman Singh", rate: 4.99, active: true}
	fmt.Println("Name:", subscriber1.name)
	fmt.Println("Rate:", subscriber1.rate)
	fmt.Println("Active?:", subscriber1.active)
}
