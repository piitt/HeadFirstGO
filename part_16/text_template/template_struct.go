package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func chkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplateStruct(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	chkerr(err)
	err = tmpl.Execute(os.Stdout, data)
	chkerr(err)
}

func main() {
	templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplateStruct(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplateStruct(templateText, Part{Name: "Cables", Count: 2})

	println("---------------------------------------------")

	// действие if используется для вывода поля Rate только в том случае,
	// если поле Active содержит true
	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	executeTemplateStruct(templateText, subscriber)

	println("---------------------------------------------")

	subscriber = Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false}
	executeTemplateStruct(templateText, subscriber)

}
