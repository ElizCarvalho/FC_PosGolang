package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Received c1:", msg1)
	case msg2 := <-c2:
		fmt.Println("Received c2:", msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout")
	default:
		fmt.Println("Default")
	}
}
