package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan int)
	go publish(channel) // Thread 2
	reader(channel)
}

func publish(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel) // fecha o canal para indicar que não serão mais enviados mais valores
}

func reader(channel chan int) {
	for i := range channel {
		fmt.Printf("Received: %d\n", i)
	}
}
