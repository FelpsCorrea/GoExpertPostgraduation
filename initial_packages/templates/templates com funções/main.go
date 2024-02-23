package main

import (
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	// Encaixar vários arquivos para formar um template
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")

	// Disponibilizando funções para usar nos templates
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	// Executar com o template criado
	t = template.Must(t.ParseFiles(templates...))

	// Executar template
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"Python", 20},
	})

	if err != nil {
		panic(err)
	}

}
