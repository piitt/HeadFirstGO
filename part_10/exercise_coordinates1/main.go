package main

import (
	"fmt"
	"headfirstgo/part_10/geosetvalidget"
	"log"
)

func main() {
	coordinates := geosetvalidget.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
