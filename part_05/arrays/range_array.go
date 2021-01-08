package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for index, note := range notes {
		fmt.Println(index, note)
	}
	fmt.Println("----------------")
	for _, note := range notes {
		fmt.Println(note)
	}
	fmt.Println("----------------")
	for index, _ := range notes {
		fmt.Println(index)
	}
}
