package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	// Encaixar vários arquivos para formar um template
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// Executar com o template criado
	t := template.Must(template.New("content.html").ParseFiles(templates...))

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
