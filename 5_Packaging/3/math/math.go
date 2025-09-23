package math

import (
	"fmt"
	"math"

	"github.com/google/uuid"
)

// Calculator representa uma calculadora básica
type Calculator struct {
	ID   string
	Name string
}

// NewCalculator cria uma nova instância da calculadora
func NewCalculator(name string) *Calculator {
	return &Calculator{
		ID:   uuid.New().String(),
		Name: name,
	}
}

// Add realiza soma
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract realiza subtração
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply realiza multiplicação
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide realiza divisão
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não é permitida")
	}
	return a / b, nil
}

// Power calcula potência
func (c *Calculator) Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// SquareRoot calcula raiz quadrada
func (c *Calculator) SquareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("não é possível calcular raiz quadrada de número negativo")
	}
	return math.Sqrt(number), nil
}

// GetInfo retorna informações da calculadora
func (c *Calculator) GetInfo() string {
	return fmt.Sprintf("Calculadora: %s (ID: %s)", c.Name, c.ID)
}
