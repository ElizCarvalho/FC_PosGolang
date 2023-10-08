package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d | cap=%d | valor=%v \n", len(s), cap(s), s)

	//s[:5] -> do inicio e dropa o resto
	//capacidade se mantém
	fmt.Printf("len=%d | cap=%d | valor=%v \n", len(s[:0]), cap(s[:0]), s[:0])

	//s[:5] -> do inicio, pega as 5 primeiras posições e dropa o resto
	//capacidade se mantém
	fmt.Printf("len=%d | cap=%d | valor=%v \n", len(s[:5]), cap(s[:4]), s[:5])

	//s[3:] -> ignora as 3 primeiras posições (pula elas) e pega o resto
	//nesse caso, a capacidade diminui, pq eu ignorei uma parte do começo
	fmt.Printf("len=%d | cap=%d | valor=%v \n", len(s[3:]), cap(s[3:]), s[3:])

	//aumentando a capacidade do slice
	s = append(s, 110)
	fmt.Printf("len=%d | cap=%d | valor=%v \n", len(s[3:]), cap(s[3:]), s[3:])
	fmt.Printf("len=%d | cap=%d | valor=%v \n", len(s), cap(s), s)

	for i, v := range s {
		fmt.Printf("O valor do indice %d é %d \n", i, v)
	}
}
