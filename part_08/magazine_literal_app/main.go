package main

import (
	"fmt"
	"headfirstgo/part_08/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active?:", subscriber.Active)

	subscriber1 := magazine.Subscriber{Rate: 4.99}
	fmt.Println("Name:", subscriber1.Name)
	fmt.Println("Rate:", subscriber1.Rate)
	fmt.Println("Active?:", subscriber1.Active)
}
