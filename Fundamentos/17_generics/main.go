package main

import "fmt"

type MyNumber int

type Number interface {
	~int | float64 // ~ significa que pode ser qualquer tipo que seja baseado em int
}

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	return a == b
}

func Compara2[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Joao": 10, "Maria": 20, "Jose": 30}
	println(Soma(m))
	m2 := map[string]float64{"Joao": 10.35, "Maria": 20.45, "Jose": 30.59}
	fmt.Printf("%.2f\n", Soma(m2))
	fmt.Printf("%.2f\n", Soma2(m2))

	m3 := map[string]MyNumber{"Joao": 40, "Maria": 20, "Jose": 30}
	println(Soma2(m3))
	println(Compara(10, 10.0))
	println(Compara2(10, 10.0))
}
