package main

import "fmt"

func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	var primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes[0], primes[2], primes[4])

	notesShot := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notesShot[5], notesShot[4], notesShot[1])
	primesShot := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primesShot[1], primesShot[3], primesShot[4])

	text := [3]string {
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line",
	}
	fmt.Println(text)
}
