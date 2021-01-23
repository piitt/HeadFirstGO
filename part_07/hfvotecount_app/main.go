package main

import (
	"fmt"
	"headfirstgo/part_07/hfdtflvote"
	"log"
)

func main() {
	lines, err := hfdtflvote.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
