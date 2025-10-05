package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1: A função main é a função principal que inicia o programa
func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	go task("A", &wg) // Thread 2: A função task é chamada como goroutine
	go task("B", &wg) // Thread 3: A função task é chamada como goroutine
	wg.Wait()
}
