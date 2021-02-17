package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// формируем html страницу из файла
	html, err := template.ParseFiles("view.html")
	check(err)
	// вставляем страницу в ответ
	// т.к. данных нет то передаем nil
	err = html.Execute(writer, nil)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	// ListenAndServe всегда возвращает ошибку (err всегда != nil)
	log.Fatal(err)
}
