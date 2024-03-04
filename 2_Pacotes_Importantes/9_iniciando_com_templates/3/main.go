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
	cursos := Cursos{
		{"C Sharp", 80},
		{"Go", 40},
		{"Java", 60}}

	tmp := template.Must(template.New("content.html").ParseFiles("content.html"))
	err := tmp.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}
