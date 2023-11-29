package main

import "fmt"

type Conta struct {
	saldo float64
}

func (c Conta) simularEmprestimo(valor float64) float64 {
	c.saldo += valor
	fmt.Printf("Saldo no simulador de emprestimo: %.2f\n", c.saldo)
	return c.saldo
}

func (c *Conta) realizarEmprestimo(valor float64) float64 {
	c.saldo += valor
	fmt.Printf("Saldo apos emprestimo: %.2f\n", c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 500}
	fmt.Printf("Saldo conta: %.2f\n", conta.saldo)
	conta.simularEmprestimo(1000)
	fmt.Printf("Saldo final: %.2f\n", conta.saldo)
	conta.realizarEmprestimo(1000)
	fmt.Printf("Saldo final: %.2f\n", conta.saldo)
}
