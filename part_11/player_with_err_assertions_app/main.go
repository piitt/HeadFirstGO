package main

import "headfirstgo/part_11/hfgadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(hfgadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	TryOut(hfgadget.TapeRecorder{})
	TryOut(hfgadget.TapePlayer{})
}
