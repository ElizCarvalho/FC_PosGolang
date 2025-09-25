package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	var productFound *entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.Product{})

	//criar 25 produtos direto no banco de dados
	for i := 0; i < 25; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		db.Create(product)
	}

	productDB := NewProductDB(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 0", products[0].Name)
	assert.Equal(t, "Product 9", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 10", products[0].Name)
	assert.Equal(t, "Product 19", products[9].Name)

	products, err = productDB.FindAll(3, 10, "desc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 4", products[0].Name)
	assert.Equal(t, "Product 0", products[4].Name)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	db.Create(product)

	productDB := NewProductDB(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	db.Create(product)

	product.Name = "Product 2"
	product.Price = 20.0
	productDB := NewProductDB(db)
	err = productDB.Update(product)
	assert.NoError(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 10.0)
	db.Create(product)

	productDB := NewProductDB(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDB.FindById(product.ID.String())
	assert.Error(t, err)
	assert.Empty(t, productFound)
}
