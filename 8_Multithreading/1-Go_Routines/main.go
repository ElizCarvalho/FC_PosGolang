package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1: A função main é a função principal que inicia o programa
func main() {
	go task("A") // Thread 2: A função task é chamada como goroutine
	go task("B") // Thread 3: A função task é chamada como goroutine
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}() // Thread 4: A função task é chamada como goroutine
	time.Sleep(10 * time.Second) // Thread 4: A função main espera 10 segundos
}
