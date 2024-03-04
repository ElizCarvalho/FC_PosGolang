package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"zero", "um", "dois", "tres", "quatro", "cinco", "seis", "sete", "oito", "nove"}
	for k, v := range numeros {
		println(k, v)
	}

	for _, v := range numeros {
		println(v)
	}

	for k, _ := range numeros {
		println(k)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("loop infinito")
		//bom para consumir filas de mensageria por exemplo
	}

}
