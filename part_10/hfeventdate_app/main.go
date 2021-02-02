package main

import (
	"fmt"
	"headfirstgo/part_10/hfeventdate"
	"log"
)

func main() {
	event := hfeventdate.Event{}
	err := event.SetTitle("Mom's birthday")
	// err := event.SetTitle("An extremely long and impractical title")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

	// fmt.Println(event.Date.Year())
	// fmt.Println(event.Date.Month())
	// fmt.Println(event.Date.Day())
}
