package main

import (
	"encoding/json"
	"net/http"

	"github.com/FelpsCorrea/GoExpertPostgraduation/API/configs"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/dto"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/entity"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)

	productHandler := NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product dto.CreateProductInput

	// Decoda o corpo da requisição e armazena em nosso DTO
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Cria uma instância de produto
	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Salva o produto no banco de dados
	err = h.ProductDB.Create(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
