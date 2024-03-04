package main

import "fmt"

type Pessoa interface {
	Desativar()
}

func Desativacao(p Pessoa) {
	p.Desativar()
}

type Aluno struct {
	Nome  string
	Idade int
	Ativo bool
}

func (a *Aluno) Desativar() {
	a.Ativo = false
}

func main() {
	joao := Aluno{
		Nome:  "João",
		Idade: 18,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", joao.Nome, joao.Idade, joao.Ativo)

	Desativacao(&joao)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", joao.Nome, joao.Idade, joao.Ativo)

}
