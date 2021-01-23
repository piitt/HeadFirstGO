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
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
