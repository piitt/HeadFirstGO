package main

import (
	"fmt"
	"headfirstgo/part_06/hfdatafileslice"
	"log"
)

func main() {
	numbers, err := hfdatafileslice.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
