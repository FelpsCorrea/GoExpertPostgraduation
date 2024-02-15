package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	// Declaração da struct curso
	curso := Curso{"Go", 40}

	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - CargaHoraria: {{.CargaHoraria}}"))

	// Executar template
	err := t.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}

}
