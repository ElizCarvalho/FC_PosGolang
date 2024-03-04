package main

import "fmt"

func main() {
	//defer é uma palavra reservada que serve para adiar a execução de uma função até o final da função atual
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")

}
