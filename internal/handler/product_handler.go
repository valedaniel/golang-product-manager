package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/valedaniel/golang-product-manager/internal/models"
	"github.com/valedaniel/golang-product-manager/internal/storage"
)

type API struct {
	storage storage.ProductStorage
	logger  *log.Logger
}

func NewRouter(storage storage.ProductStorage, logger *log.Logger) http.Handler {
	logger = log.New(os.Stdout, "HANDLER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	api := &API{
		storage: storage,
		logger:  logger,
	}

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API de Gerenciamento de Produtos"))
	}))

	serveMux.HandleFunc("/products", api.productsHandler)
	serveMux.HandleFunc("/products/", api.productsIdHandler)

	return serveMux
}

func (api *API) productsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		api.handleCreateProduct(writer, request)
	case http.MethodGet:
		api.handleListProducts(writer, request)
	default:
		http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func (api *API) productsIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		api.handleGetByIdProduct(writer, request)
	case http.MethodPut:
		api.handleUpdateProduct(writer, request)
	case http.MethodDelete:
		api.handleDeleteProduct(writer, request)
	default:
		http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func (api *API) handleCreateProduct(writer http.ResponseWriter, request *http.Request) {
	var product models.Product

	if err := readJSON(request, &product); err != nil {
		api.logger.Printf("Erro ao decodificar JSON: %v", err)
		writeJSON(writer, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
		return
	}

	if product.Name == "" || product.Price <= 0 {
		writeJSON(writer, http.StatusBadRequest, map[string]string{"error": "Nome e preço são obrigatórios e o preço deve ser maior que zero"})
		return
	}

	if err := api.storage.Create(request.Context(), &product); err != nil {
		api.logger.Printf("Erro ao criar produto: %v", err)
		writeJSON(writer, http.StatusInternalServerError, map[string]string{"error": "Erro ao criar produto"})
		return
	}

	writeJSON(writer, http.StatusCreated, product)
}

func (api *API) handleGetByIdProduct(writer http.ResponseWriter, request *http.Request) {
	idStr := strings.TrimPrefix(request.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSON(writer, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	product, err := api.storage.Get(request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeJSON(writer, http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
			return
		}

		api.logger.Printf("Erro ao buscar o produto: %v", err)
		writeJSON(writer, http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar produto"})
		return
	}

	writeJSON(writer, http.StatusOK, product)
}

func (api *API) handleListProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := api.storage.List(request.Context())
	if err != nil {
		api.logger.Printf("Erro ao listar produtos: %v", err)
		writeJSON(writer, http.StatusInternalServerError, map[string]string{"error": "Erro ao listar produtos"})
		return
	}

	if products == nil {
		writeJSON(writer, http.StatusOK, []models.Product{})
		return
	}

	writeJSON(writer, http.StatusOK, products)
}

func (api *API) handleUpdateProduct(writer http.ResponseWriter, request *http.Request) {
	var product *models.Product

	idStr := strings.TrimPrefix(request.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSON(writer, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	if err := readJSON(request, &product); err != nil {
		api.logger.Printf("Erro ao decodificar JSON: %v", err)
		writeJSON(writer, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
		return
	}

	response, err := api.storage.Update(request.Context(), product, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeJSON(writer, http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
			return
		}

		api.logger.Printf("Erro ao atualizar o produto: %v", err)
		writeJSON(writer, http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar produto"})
		return
	}

	if !response {
		writeJSON(writer, http.StatusInternalServerError, map[string]string{"error": "Não foi possível atualizar o produto"})
		return
	}

	writeJSON(writer, http.StatusOK, true)
}

func (api *API) handleDeleteProduct(writer http.ResponseWriter, request *http.Request) {
	idStr := strings.TrimPrefix(request.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSON(writer, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	deleteErr := api.storage.Delete(request.Context(), id)
	if deleteErr != nil {
		if errors.Is(deleteErr, sql.ErrNoRows) {
			writeJSON(writer, http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
			return
		}

		api.logger.Printf("Erro ao deletar o produto: %v", deleteErr)
		writeJSON(writer, http.StatusInternalServerError, map[string]string{"error": "Erro ao deletar produto"})
		return
	}

	writeJSON(writer, http.StatusOK, true)

}

func writeJSON(writer http.ResponseWriter, status int, v any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(v)
}

func readJSON(request *http.Request, v any) error {
	return json.NewDecoder(request.Body).Decode(v)
}
