package main

import "fmt"

// функция с аргументом канал в котором передаются строковые значения
func greeting(myChannel chan string) {
	// отправляет строку по каналу
	myChannel <- "hi"
}

func main() {
	// создаем канал
	myChannel := make(chan string)
	// вызываем горутину
	go greeting(myChannel)
	// получение значения из канала
	//fmt.Println(<-myChannel)

	receivedValue := <-myChannel
	fmt.Println(receivedValue)
}
