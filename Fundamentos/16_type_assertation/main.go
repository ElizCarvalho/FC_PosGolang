package main

import "fmt"

func main() {
	var pepe1 interface{} = "Pos Go Lang"
	println(pepe1)          //como ele é interface entao na hora de print ele nao consegue printar o tipo
	println(pepe1.(string)) //consigo pegar uma interface e converter para o tipo que eu quiser (se for possivel)

	resultado, ok := pepe1.(int) //resultado vai receber o valor da variavel pepe1 e ok vai receber true ou false se a conversao foi bem sucedida
	fmt.Printf("O valor de resultado é %v e o valor de ok é %v\n", resultado, ok)
}
