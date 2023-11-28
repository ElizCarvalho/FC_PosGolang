package main

func main() {
	// variavel  -> ponteiro que tem um endereço na memoria (0x0001) -> valor (10)
	a := 10
	println(a)
	println(&a) // &a é o endereço de memória de a (0xc00003a728)

	var ponteiro *int = &a //ponteiro é uma variável que armazena o endereço de memória de a
	println(ponteiro)      // 0xc00003a728

	*ponteiro = 20 // *ponteiro é o valor que está no endereço de memória de a
	println(a)     // 20

	b := &a
	println(b)  // 0xc00003a728
	println(*b) // 20
	*b = 30
	println(a) // 30
}
