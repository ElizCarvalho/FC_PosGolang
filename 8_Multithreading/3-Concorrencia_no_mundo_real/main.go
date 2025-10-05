package main

import (
	"fmt"
	"net/http"
	"sync"
)

var number uint64 = 0

func main() {
	mutex := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock() // bloqueia a região de código para que apenas uma goroutine acesse por vez
		number++
		mutex.Unlock() // desbloqueia a região de código para que outras goroutines possam acessar
		w.Write([]byte(fmt.Sprintf("Voce visitou esta pagina %d vezes", number)))
	})
	http.ListenAndServe(":3000", nil)

	//testar com o comando: ab -n 1000 -c 10 http://localhost:3000/
	// no codigo sem o mutex o numero de visitas nao é atualizado corretamente
	// pois o numero de visitas é atualizado de forma concorrente
	//pode usar o comando: go run -race main.go para verificar se há race conditions

	// com o mutex o numero de visitas é atualizado corretamente
	// mutex é um mecanismo que permite que apenas uma goroutine acesse uma região de código por vez

	//uma outra forma de resolver é usando o package atomic
	// atomic.AddUint64(&number, 1)

}
