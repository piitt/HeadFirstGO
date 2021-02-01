package main

import (
	"fmt"
	"headfirstgo/part_10/geoset"
)

func main() {
	coordinates := geoset.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
