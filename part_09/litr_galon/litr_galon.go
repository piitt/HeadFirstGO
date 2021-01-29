package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters

	// переводим литры в галлоны
	carFuel = Gallons(Liters(40.0) * 0.264)
	// переводим галлоны в литры
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
}
