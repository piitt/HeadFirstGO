package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
}

func main() {
	sizes := make(chan int)

	//go responseSize("http://example.com", sizes)
	//go responseSize("http://mail.ru", sizes)
	//go responseSize("http://ya.ru", sizes)
	//fmt.Println(<-sizes)
	//fmt.Println(<-sizes)
	//fmt.Println(<-sizes)

	urls := []string{"http://example.com", "http://mail.ru", "http://ya.ru"}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}
