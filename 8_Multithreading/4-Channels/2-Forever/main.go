package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Executando goroutine")
			time.Sleep(1 * time.Second)
		}
		forever <- true
	}()

	<-forever
}
