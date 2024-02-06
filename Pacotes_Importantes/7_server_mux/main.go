package main

import (
	"fmt"
	"net/http"
)

func main() {

	finish := make(chan bool) // Canal para bloquear a finalização da função main

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{Title: "My Blog"})
	fmt.Println("Server running at :8000")

	go func() {
		http.ListenAndServe(":8000", mux)
	}()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi / Mux2!"))
	})
	fmt.Println("Server running at :8001")

	go func() {
		http.ListenAndServe(":8001", mux2)
	}()

	<-finish // Bloqueia a finalização da função main

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi HomeHandler!"))
}

type blog struct {
	Title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi " + b.Title + "!"))
}
