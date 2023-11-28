package main

import "fmt"

func main() {
	fmt.Println(soma(0, 1, 1, 2, 3, 5, 8, 13, 21))
}

// '...' diz que vai entrar um número variável de parâmetros, tds do mesmo tipo
// ...int é um slice de int
func soma(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
