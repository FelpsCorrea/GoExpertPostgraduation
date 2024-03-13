package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/dto"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/entity"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/infra/database"
	entityPkg "github.com/FelpsCorrea/GoExpertPostgraduation/API/pkg/entity"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// CreateProduct Product godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductInput true "product request"
// @Success 201
// @Failure 500	{object} Error
// @Router /products [post]
// @Security ApiKeyAuth
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

// GetProduct Product godoc
// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Success 200 {object} Product
// @Failure 404
// @Failure 500	{object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	// Pegar o id no corpo da requisição
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}

// GetProducts Product godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Param page query int false "page number"
// @Param limit query int false "limit per page"
// @Param sort query string false "sort by asc or desc"
// @Success 200 {object} []Product
// @Failure 500	{object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	// Pegar os dados de paginação e ordenação
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	// Converte os valores para inteiros
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	// Faz a busca
	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Devolve a resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct Product godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Param product body dto.CreateProductInput true "product request"
// @Success 200
// @Failure 500	{object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Objeto para receber os parametros da requisição
	var product entity.Product

	// Decoda o corpo da requisição e armazena em nosso objeto
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verifica se o ID é válido
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verifica se o produto existe
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Atualiza o produto no banco de dados
	err = h.ProductDB.Update(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// DeleteProduct Product godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Success 200
// @Failure 500	{object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verifica se o produto existe
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Deleta o produto no banco de dados
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
