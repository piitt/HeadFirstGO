package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Here's my template!\n"
	// создание и загрузка шаблона
	tmpl, err := template.New("test").Parse(text)
	check(err)
	// вывести шаблон на стандартный вывод
	// nil птому что нет данных
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}
