package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Servidor: request iniciada")
	defer log.Println("Servidor: request finalizada")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Servidor: request processada com sucesso")
		w.Write([]byte("Cliente: request processada com sucesso"))
		return
	case <-ctx.Done():
		log.Println("Servidor: request cancelada pelo cliente")
		return
	}
}
