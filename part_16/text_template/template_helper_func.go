package main

import (
	"log"
	"os"
	"text/template"
)

func checkerror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	checkerror(err)
	err = tmpl.Execute(os.Stdout, data)
	checkerror(err)
}

func main() {
	executeTemplate("Dot is : {{.}}!\n", "ABC")
	executeTemplate("Dot is : {{.}}!\n", 123.5)
}
