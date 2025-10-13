package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entrou no RecoverMiddleware")
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
				debug.PrintStack()
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic in panic handler")
	})

	fmt.Println("Server running at :3000")
	if err := http.ListenAndServe(":3000", RecoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on port 3000: %v", err)
	}
}
