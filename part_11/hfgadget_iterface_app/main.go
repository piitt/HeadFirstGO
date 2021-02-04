package main

import (
	"headfirstgo/part_11/hfgadget"
)

// интерфейс можно определить и в пакете gadget,
// но определение интерфейса в том пакете, где он используется,
// обеспечивает большую гибкость
type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = hfgadget.TapePlayer{}
	playList(player, mixtape)
	player = hfgadget.TapeRecorder{}
	playList(player, mixtape)
}
