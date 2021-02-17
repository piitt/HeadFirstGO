package main

import (
	"log"
	"os"
	"text/template"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// действия обозначаются {{}}
	// в них указываются данные, которые надо вставить
	// или операция, которая должна быть выполнена шаблоном
	// внутри действия можно обращаться к значению данных,
	// переданых методу Execut, в форме . (точка)
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	checkerr(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	checkerr(err)
	err = tmpl.Execute(os.Stdout, 42)
	checkerr(err)
	err = tmpl.Execute(os.Stdout, true)
	checkerr(err)
}
