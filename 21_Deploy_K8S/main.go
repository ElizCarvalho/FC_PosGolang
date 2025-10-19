package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Aula de Deploy K8S"))
	})
	http.ListenAndServe(":8080", nil)
}
