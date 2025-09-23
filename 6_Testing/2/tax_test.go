package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	assert.Equal(t, expected, result)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)                                   // mockando a função SaveTax para retornar nil
	repository.On("SaveTax", 0.0).Return(errors.New("Erro ao salvar o imposto")) // mockando a função SaveTax para retornar erro

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err2 := CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err2, "Erro ao salvar o imposto")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
