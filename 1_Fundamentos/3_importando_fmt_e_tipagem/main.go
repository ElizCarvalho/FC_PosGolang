package main

import "fmt"

type PEPE int

var (
	b bool
	c int
	d         = "PEPE"
	e float32 = 1.2
	f PEPE
)

func main() {
	fmt.Printf("O tipo da variavel B é %T e o valor é %v \n", b, b)
	fmt.Printf("O tipo da variavel C é %T e o valor é %d \n", c, c)
	fmt.Printf("O tipo da variavel D é %T e o valor é %v \n", d, d)
	fmt.Printf("O tipo da variavel E é %T e o valor é %v \n", e, e)
	fmt.Printf("O tipo da variavel F é %T e o valor é %d \n", f, f)
}
