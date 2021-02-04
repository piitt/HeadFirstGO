package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println(index, notes[index])
	index = 3
	fmt.Println(index, notes[index])

	fmt.Println("--------------------------")

	notesFor := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notesFor[i])
	}

	fmt.Println("--------------------------")

	//for i := 0; i <= 7; i++ {
	//	fmt.Println(i, notesFor[i])
	//}

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(len(notesFor))
	fmt.Println(len(primes))

	fmt.Println("------------------")

	for i := 0; i < len(notesFor); i++ {
		fmt.Println(i, notesFor[i])
	}
}
