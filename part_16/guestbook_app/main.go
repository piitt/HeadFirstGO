package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	// если возвращается ошибка, указывающая на то, что файл не существует
	// вернуть nil вместо сегмента строк
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	// получаем данные из поля формы
	signature := request.FormValue("signature")
	// с какими параметрами открыть файл
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// открытие файла
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	// записывам текст с новой строки в файл
	_, err = fmt.Fprintln(file, signature)
	check(err)
	// закрываем файл
	err = file.Close()
	check(err)
	time.Sleep(1 * time.Second)
	// переход на другую страницу
	// "/guestbook" куда переходить
	// http.StatusFound код ответа, состояния сервера 302
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
