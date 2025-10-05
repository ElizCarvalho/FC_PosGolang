package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

// Thread 1
func main() {
	data := make(chan int)

	qntWorkers := 1000000

	//inicia os workers
	for i := 0; i < qntWorkers; i++ {
		go worker(i, data)
	}

	//envia os dados para os workers
	for i := 0; i < 100000000; i++ {
		data <- i
	}
}
