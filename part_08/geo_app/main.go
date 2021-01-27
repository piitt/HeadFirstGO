package main

import (
	"fmt"
	"headfirstgo/part_08/geo"
)

func main() {
	location := geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
