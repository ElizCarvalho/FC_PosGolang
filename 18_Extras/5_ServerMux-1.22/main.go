package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("GET /books/{id}", GetBookHandler)         //parâmetro reconhecido pelo mux no handler
	//mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler) //esse {d...} é um parâmetro pegar todos os segmentos da url após o /books/dir/
	//mux.HandleFunc("GET /books/{$}", BooksHandler) // exato, esse {$} é um parâmetro pegar o último segmento da url após o /books/

	//mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHandler)
	//mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Handler)

	mux.HandleFunc("GET /books/{s}", BooksPrecedenceHandler)            //esse {s} é um parâmetro pegar o último segmento da url após o /books/
	mux.HandleFunc("GET /category/{s}/latest", BooksPrecedence2Handler) // Rota mais específica para evitar conflito
	http.ListenAndServe(":9000", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d") // Access captured directory path segments as slice
	fmt.Fprintf(w, "Accessing directory path: %s\n", dirpath)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2"))
}
