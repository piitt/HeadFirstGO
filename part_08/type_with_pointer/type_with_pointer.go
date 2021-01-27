package main

import "fmt"

type subscrib struct {
	name   string
	rate   float64
	active bool
}

func applyDiscounts(s *subscrib) {
	s.rate = 4.99
}

func main() {
	var s subscrib
	applyDiscounts(&s)
	fmt.Println(s.rate)
}
