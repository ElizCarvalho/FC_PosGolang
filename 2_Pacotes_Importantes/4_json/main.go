package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"numero"`
	Saldo  float64 `json:"saldo"` //para ignorar o campo na conversão para JSON, basta usar `json:"-"`
}

func main() {
	conta := Conta{123, 100.0}
	fmt.Println("Opcao 1: Convertendo para JSON...")
	res, err := json.Marshal(conta)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

	fmt.Println("Opcao 2: Convertendo para JSON...")
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Convertendo de JSON para struct...")

	jsonConta := []byte(`{"numero":456,"saldo":200.0}`)
	var contaX Conta
	err = json.Unmarshal(jsonConta, &contaX) //o & é para passar o endereço de memoria da variável
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contaX.Saldo)

}
