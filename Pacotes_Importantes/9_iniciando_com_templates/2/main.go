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
	curso := Curso{"C Sharp", 80}
	tmp := template.Must(
		template.New("CursoTemplate").
			Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas."))
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
