package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testProductName = "Product 1"
const testProductPrice = 10.0

func TestNewProduct(t *testing.T) {
	p, err := NewProduct(testProductName, testProductPrice)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, testProductName, p.Name)
	assert.Equal(t, testProductPrice, p.Price)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
	assert.False(t, p.CreatedAt.IsZero())
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", testProductPrice)
	assert.Error(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct(testProductName, 0.0)
	assert.Error(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct(testProductName, -1.0)
	assert.Error(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}
