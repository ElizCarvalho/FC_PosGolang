package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 70

	fmt.Printf("O valor da ultima posicao: %v \n", meuArray[len(meuArray)-1])
	fmt.Printf("Tamanho do meu array: %v \n", len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d \n", i, v)
	}
}
