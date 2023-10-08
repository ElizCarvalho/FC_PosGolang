package main

import "fmt"

func main() {
	salaries := map[string]int{"Wesley": 1200, "João": 1000, "Maria": 500}
	fmt.Println(salaries)
	fmt.Println(salaries["Wesley"])

	//deletar um elemento do map
	delete(salaries, "Maria")
	fmt.Println(salaries)

	//incluir um elemento no map
	salaries["Pedro"] = 1500
	fmt.Println(salaries)

	//posso criar um map usando make
	//sal1 := make(map[string]int)

	//e tbm posso criar um map usando atribuição
	//sal2 := map[string]int{}

	//percorrer o map salaries
	for key, value := range salaries {
		fmt.Printf("%s recebe %d\n", key, value)
	}

	for _, value := range salaries {
		fmt.Printf("O salario é %d", value)
	}
}
