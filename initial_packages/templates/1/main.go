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

	// Criação do template
	tmp := template.New("CursoTemplate")

	tmp, _ = tmp.Parse("Curso: {{.Nome}} - CargaHoraria: {{.CargaHoraria}}")

	// Executar template
	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}

}
