package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Get возвращает объект http.Response
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	// http.Response - структура с полем Body
	// Body поддерживает интерфейс ReadCloser (методы Read и Close)
	defer response.Body.Close()
	// ioutil.ReadAll читает все содержимое и возвращает сегмент значений byte
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
