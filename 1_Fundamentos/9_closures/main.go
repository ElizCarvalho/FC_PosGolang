package main

import "fmt"

func main() {

	// Closure é uma função que retorna outra função
	// A função interna tem acesso às variáveis da função externa
	// A função interna é um closure
	t := func() int {
		return sum(0, 1, 1, 2, 3, 5) * 3
	}()
	fmt.Println(t)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
