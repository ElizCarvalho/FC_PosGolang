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
	tmp := template.New("CursoTemplate")
	tmp, err := tmp.Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas.")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
