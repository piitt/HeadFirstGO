package main

import (
	"fmt"
	"headfirstgo/part_10/hfcalendar"
	"log"
)

func main() {
	//date := hfcalendar.Date{}
	//date.year = 2019
	//date.month = 14
	//date.day = 50
	//fmt.Println(date)
	//
	//date = hfcalendar.Date{year: 0, month: 0, day: -2}
	//fmt.Println(date)

	date := hfcalendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
