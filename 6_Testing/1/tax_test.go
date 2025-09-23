package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5},
		{1000.0, 10},
		{1500.0, 10},
	}

	for _, table := range table {
		result := CalculateTax(table.amount)
		if result != table.expected {
			t.Errorf("Expected %f but got %f", table.expected, result)
		}
	}
}

// pra rodar o benchmark usamos o comando $ go test -bench=.
// ou $ go test -bench=. -run=^$ (pra rodar apenas o benchmark)
// podemos usar count com o comando $ go test -count=10 (pra rodar 10 vezes)
// o comando go test -benchmem vai mostrar o quanto de memoria é usado em cada benchmark
// o comando go test -benchtime=10s vai rodar o benchmark por 10 segundos
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
