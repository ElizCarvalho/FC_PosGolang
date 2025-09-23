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
