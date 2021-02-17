package main

import (
	"log"
	"os"
	"text/template"
)

func checker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplateIf(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	checker(err)
	err = tmpl.Execute(os.Stdout, data)
	checker(err)
}

func main() {
	executeTemplateIf("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplateIf("start {{if .}}Dot is true!{{end}} finish\n", false)
}
