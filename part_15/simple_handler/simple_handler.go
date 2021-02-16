package main

import (
	"log"
	"net/http"
)

// writer http.ResponseWriter значение которое используется для записи данных в ответ браузеру
// request *http.Request указатель на значение запрос браузера
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	// добавление строки в ответ
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// функция-обработчик при запросе get /hello дёргать фнкцию viewHandler
	http.HandleFunc("/hello", viewHandler)
	// запуск веб сервера
	// nil означает, что запросы будут обрабатывться с использованием функций
	// заданных при помощи HandleFunc
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
