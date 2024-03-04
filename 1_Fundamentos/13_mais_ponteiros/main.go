package main

func soma(num1, num2 int) int {
	num1 = 50
	return num1 + num2
}

func subtrai(num1, num2 *int) int {
	*num1 = 50
	return *num1 - *num2
}

func main() {
	pepe1 := 10
	pepe2 := 20
	print("soma: ")
	println(soma(pepe1, pepe2)) //nesse caso estou passando copia dos valores
	println("pepe1: ", pepe1)

	print("subtraçao: ")
	println(subtrai(&pepe1, &pepe2)) //nesse caso estou passando o endereço de memoria
	println("pepe1: ", pepe1)
}
