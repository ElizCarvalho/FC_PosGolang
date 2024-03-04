package main

import (
	"html/template"
	"net/http"
	"strings"
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

	templates := []string{"header.html", "content.html", "footer.html"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			{"C#", 80},
			{"Go", 40},
			{"Java", 60}}

		tmp := template.New("content.html")
		tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})
		tmp = template.Must(tmp.ParseFiles(templates...))
		err := tmp.Execute(w, cursos)
		if err != nil {
			http.Error(w, "Não foi possível exibir a página...", http.StatusInternalServerError)
			return
		}
	})

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err.Error())
	}
}
