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

	// Executar com o template criado
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

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
