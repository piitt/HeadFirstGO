package main

import (
	"log"
	"os"
	"text/template"
)

func chk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplateRange(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	chk(err)
	err = tmpl.Execute(os.Stdout, data)
	chk(err)
}

func main() {
	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplateRange(templateText, []string{"do", "re", "mi"})

	println("-----------------------------------------------")

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplateRange(templateText, []float64{1.25, 0.99, 27})

	println("-----------------------------------------------")

	// если range передать пустое значение или nil цикл выполняться не будет
	executeTemplateRange(templateText, []float64{})
	executeTemplateRange(templateText, nil)
}
