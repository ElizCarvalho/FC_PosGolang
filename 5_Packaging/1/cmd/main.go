package main

import (
	"fmt"

	"github.com/ElizCarvalho/FC_PosGolang/5_Packaging/1/math" //importando o pacote math
)

func main() {
	fmt.Println("Hello, World!")
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	fmt.Println(m.Sub())
}
