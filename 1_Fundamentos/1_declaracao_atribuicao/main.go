package main

const A = "PEPE"

// declaração de variavel  de escopo GLOBAL
var (
	b bool
	c int
	d string
	e float32
)

func main() {
	println("Hello World!")
	println("---- CONSTANTE ESCOPO GLOBAL ----")
	println(A)
	println("---- VAR ESCOPO GLOBAL: PRINT VALORES DEFAULT ----")
	println(b)
	println(c)
	println(d)
	println(e)
	println("---- VAR ESCOPO LOCAL: PRINT VALORES DEFAULT ----")
	var f bool
	println(f)
	println("---- VAR ESCOPO LOCAL: DECLARACAO COM ATRIBUICAO ----")
	var g = "PEPE LOCAL"
	println(g)
	println("---- VAR ESCOPO LOCAL: SHORTHAND ASSIGNMENT ----")
	h := "Shorthand Assignment"
	println(h)
	println("---- VAR ESCOPO PACKAGE ----")
	println(soma)
}
