package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewProduct(t *testing.T) {
	p, err := CreateProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := CreateProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := CreateProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenInvalidPrice(t *testing.T) {
	p, err := CreateProduct("Product 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := CreateProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
