package main

import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// создаем два канала
	channel1 := make(chan string)
	channel2 := make(chan string)
	// каждый канал передается функции выполняемой в новой горутине
	go abc(channel1)
	go def(channel2)
	// получение и вывод значений из каналов (по порядку)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}
