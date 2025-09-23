package main

import (
	"fmt"
	"strings"

	calcmath "calculator-math"
	calcutils "calculator-utils"
)

func main() {
	// Criar instâncias dos módulos
	calculator := calcmath.NewCalculator("Calculadora Avançada")
	formatter := calcutils.NewFormatter()

	// Mostrar informações dos módulos
	fmt.Println(formatter.FormatHeader("CALCULADORA WORKSPACE DEMO"))
	fmt.Println(calculator.GetInfo())
	fmt.Println(formatter.GetInfo())

	// Demonstrações de operações
	fmt.Println(formatter.FormatHeader("OPERACOES MATEMATICAS"))

	// Soma
	a, b := 15.5, 4.3
	result := calculator.Add(a, b)
	fmt.Println(formatter.FormatOperation("+", a, b, result))

	// Subtração
	result = calculator.Subtract(a, b)
	fmt.Println(formatter.FormatOperation("-", a, b, result))

	// Multiplicação
	result = calculator.Multiply(a, b)
	fmt.Println(formatter.FormatOperation("*", a, b, result))

	// Divisão
	result, err := calculator.Divide(a, b)
	if err != nil {
		fmt.Println(formatter.FormatError(err))
	} else {
		fmt.Println(formatter.FormatOperation("/", a, b, result))
	}

	// Potência
	base, exp := 2.0, 3.0
	result = calculator.Power(base, exp)
	fmt.Println(formatter.FormatOperation("^", base, exp, result))

	// Raiz quadrada
	number := 16.0
	result, err = calculator.SquareRoot(number)
	if err != nil {
		fmt.Println(formatter.FormatError(err))
	} else {
		fmt.Printf("√%.2f = %.2f\n", number, result)
	}

	// Formatação de moeda
	fmt.Println(formatter.FormatHeader("FORMATACAO DE MOEDA"))
	value := 1234.5678
	fmt.Println(formatter.FormatCurrency(value, "R$"))
	fmt.Println(formatter.FormatCurrency(value, "USD"))
	fmt.Println(formatter.FormatCurrency(value, "EUR"))

	// Formatação de números
	fmt.Println(formatter.FormatHeader("FORMATACAO DE NUMEROS"))
	fmt.Printf("Número com 2 casas: %s\n", formatter.FormatNumber(value, 2))
	fmt.Printf("Número com 4 casas: %s\n", formatter.FormatNumber(value, 4))

	// Mensagens de sucesso
	fmt.Println(formatter.FormatSuccess("Todas as operações executadas com sucesso!"))
	fmt.Println(formatter.FormatSuccess("Workspace funcionando perfeitamente!"))

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🎉 Demo do Go Workspace concluída!")
	fmt.Println("📁 Módulos: math, utils, calculator")
	fmt.Println("🔗 Todos os módulos compartilham dependências")
	fmt.Println("⚡ Desenvolvimento local integrado")
}
