package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	println("Criando um arquivo....")
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	println("Escrevendo no arquivo...")
	//tamanho, err := f.WriteString("Hello World")
	tamanho, err := f.Write([]byte("Hello World com array de bytes"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso. Tamanho: %d bytes\n", tamanho)

	println("Fechando o arquivo...")
	err = f.Close()
	if err != nil {
		panic(err)
	}

	println("Lendo um arquivo inteiro...")
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo lido em bytes:")
	fmt.Println(data)
	fmt.Println("Arquivo lido em string:")
	fmt.Println(string(data))

	println("Lendo um arquivo grande pouco a pouco....")
	arquivo, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10)
	for {
		tamanho, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:tamanho]))
	}

	println("Removendo o arquivo...")
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}

	println("Arquivo removido com sucesso.")
}
