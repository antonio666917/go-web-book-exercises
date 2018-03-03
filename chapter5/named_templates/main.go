package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t, err := template.New("test").Parse(`{{define "T"}}{{ $variable := "A variable" }}Hello, {{.}}! Hello, {{ $variable }}.{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "World")
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
