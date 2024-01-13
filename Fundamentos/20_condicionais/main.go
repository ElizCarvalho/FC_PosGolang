package main

func main() {

	a := 10
	b := 20
	c := 30
	if a > b {
		println("a é maior que b")
	} else {
		println("b é maior que a")
	}

	if a > b && a > c {
		println("a é o maior")
	}

	if a > b || a > c {
		println("a é um dos maiores")
	}

	switch a {
	case 10:
		println("a é igual a 10")
	case 20:
		println("a é igual a 20")
	default:
		println("a não é igual a 10 nem a 20")
	}
}
