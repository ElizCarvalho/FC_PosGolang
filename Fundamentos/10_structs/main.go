package main

import "fmt"

// uma struct pode ser formada por outras (composição) mas não trabalha com herança
type ClienteComComposicao struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Local Endereco
}

func (c *Cliente) DesativarCliente() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

func main() {
	julia := ClienteComComposicao{
		Nome:  "Julia",
		Idade: 36,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", julia.Nome, julia.Idade, julia.Ativo)

	julia.Ativo = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", julia.Nome, julia.Idade, julia.Ativo)

	julia.Cidade = "São Paulo" //ou
	fmt.Printf("Nome: %s, Cidade: %s\n", julia.Nome, julia.Cidade)

	julia.Endereco.Cidade = "Rio de Janeiro"
	fmt.Printf("Nome: %s, Cidade: %s\n", julia.Nome, julia.Endereco.Cidade)

	ana := Cliente{
		Nome:  "Ana",
		Idade: 61,
		Ativo: true,
	}
	ana.Local.Cidade = "Natal" //como criou a struct sem composicao então precisa acessar pelo endereco
	fmt.Printf("Nome: %s, Cidade: %s\n", ana.Nome, ana.Local.Cidade)

	ana.DesativarCliente()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ana.Nome, ana.Idade, ana.Ativo)
}
