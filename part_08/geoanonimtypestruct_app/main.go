package main

import (
	"fmt"
	"headfirstgo/part_08/geoanonimtypestruct"
)

func main() {
	location := geoanonimtypestruct.Landmark{}
	location.Name = "The Googleplex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
