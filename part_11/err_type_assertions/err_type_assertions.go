package main

import "fmt"

type TapePlayer struct {
	Batteries string
}

type TapeRecorder struct {
	Microphones int
}

type Player interface {
	Play(string)
	Stop()
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

//func TryOut(player Player) {
//	player.Play("Test Track")
//	player.Stop()
//	recorder := player.(TapeRecorder)
//	recorder.Record()
//}

func main() {
	//TryOut(TapeRecorder{})
	//TryOut(TapePlayer{})
	// panic: interface conversion: main.Player is main.TapePlayer, not main.TapeRecorder

	var player Player = TapePlayer{}
	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}
