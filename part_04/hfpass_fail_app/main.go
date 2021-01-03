package main

import (
	"fmt"
	"headfirstgo/part_04/hfkeyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := hfkeyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
