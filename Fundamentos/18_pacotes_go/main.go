package main

import (
	"fmt"
	"github.com/ElizCarvalho/curso-go/math"
	"github.com/google/uuid"
)

/*
go mod init github.com/ElizCarvalho/curso-go (cria o arquivo go.mod)
go mod tidy (mantem os so pacotes que estao sendo usados e atualiza as versoes no go.mod)
go get github.com/google/uuid (instala o pacote)
*/
func main() {
	s := math.Soma(3, 5)
	fmt.Printf("O resultado da soma é: %v\n", s)
	u := uuid.New()
	fmt.Printf("O UUID gerado foi: %v\n", u)
}
