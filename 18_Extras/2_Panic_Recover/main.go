package main

import "fmt"

func panic1() {
	panic("Panic1: Something went wrong")
}

func panic2() {
	panic("Panic2: Something went wrong again")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "Panic1: Something went wrong" {
				fmt.Println("panic1 recovered")
			}
			if r == "Panic2: Something went wrong again" {
				fmt.Println("panic2 recovered")
			}
		}
	}()

	panic2()
	fmt.Println("After panic")
}
