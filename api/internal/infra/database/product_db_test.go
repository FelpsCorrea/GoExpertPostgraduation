package database

import (
	"fmt"
	"testing"

	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/entity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// Criando entidade product
	product, _ := entity.NewProduct("Videogame", 100)

	// Criando uma instancia do "model" product
	productDB := NewProduct(db)

	// Criando o product no banco
	err = productDB.Create(product)
	assert.Nil(t, err)

	// Buscando o produto criado
	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

}

func TestFindAllProduct(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)

		db.Create(product)
	}

	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	// Pagina 3 tem só 4 registros
	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)

}

func TestFindProductByID(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// Criando entidade product
	product, _ := entity.NewProduct("Videogame", 100)

	// Criando uma instancia do "model" product
	productDB := NewProduct(db)

	// Criando o product no banco
	err = productDB.Create(product)
	assert.Nil(t, err)

	// Buscando o produto criado a partir do id
	productFound, err := productDB.FindByID(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// Criando entidade product
	product, _ := entity.NewProduct("Videogame", 100)

	// Criando uma instancia do "model" product
	productDB := NewProduct(db)

	// Criando o product no banco
	err = productDB.Create(product)
	assert.Nil(t, err)

	product.Name = "Videogame 2"
	product.Price = 200

	err = productDB.Update(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// Criando entidade product
	product, _ := entity.NewProduct("Videogame", 100)

	// Criando uma instancia do "model" product
	productDB := NewProduct(db)

	// Criando o product no banco
	err = productDB.Create(product)
	assert.Nil(t, err)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
}
