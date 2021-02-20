package main

import (
	"fmt"
	"time"
)

func sendLetters(channel chan string) {
	// отправляем 4 значения ожидая 1 сек пред каждым
	time.Sleep(1 * time.Second)
	channel <- "a"
	time.Sleep(1 * time.Second)
	channel <- "b"
	time.Sleep(1 * time.Second)
	channel <- "c"
	time.Sleep(1 * time.Second)
	channel <- "d"
}

func main() {
	// время запуска программы
	fmt.Println(time.Now())
	// создаем небуферизованный канал
	channel1 := make(chan string)
	// запуск функции в горутине
	go sendLetters(channel1)
	// приостанавливаем горутину main на 5 сек
	time.Sleep(5 * time.Second)
	// получение и вывод 4 значений с текущим временем
	fmt.Println(<-channel1, time.Now())
	fmt.Println(<-channel1, time.Now())
	fmt.Println(<-channel1, time.Now())
	fmt.Println(<-channel1, time.Now())
	// время завершения программы
	fmt.Println(time.Now())
}
