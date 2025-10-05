package main

import "fmt"

// Thread 1: A função main é a função principal que inicia o programa
func main() {
	channel := make(chan string) // Canal vazio

	// Thread 2
	go func() {
		channel <- "Hello" // Canal cheio
	}()

	// Thread 1
	message := <-channel // Canal esvaziado
	fmt.Println(message)
}
