package main

import "fmt"

type NoiseMarker interface {
	MakeSound()
}

type Whistle string

type Horn string

type Robot string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func play(n NoiseMarker) {
	n.MakeSound()
	n.Walk()
}

func main() {
	var toy NoiseMarker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
	play(Robot("Botco Ambler"))
}
