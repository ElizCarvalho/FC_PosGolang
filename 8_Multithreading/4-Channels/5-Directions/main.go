package main

import "fmt"

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Elis", hello)
	le(hello)
}

// canal receive only
func recebe(nome string, hello chan<- string) {
	hello <- "Hello " + nome
}

// canal send only
func le(data <-chan string) {
	fmt.Println(<-data)
}
