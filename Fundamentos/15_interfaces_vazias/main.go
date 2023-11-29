package main

import "fmt"

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}

func main() {
	var x interface{} = 10 //funciona meio que um generics
	showType(x)
	x = "string"
	showType(x)
	x = true
	showType(x)
	x = 1.5
	showType(x)
	x = []int{1, 2, 3}
	showType(x)
	x = map[string]int{
		"chave": 10,
	}
	showType(x)
	x = func() {
		fmt.Println("Olá, mundo!")
	}
	showType(x)
}
