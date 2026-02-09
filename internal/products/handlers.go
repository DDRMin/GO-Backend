package products

import (
	"encoding/json"
	"net/http"

	"github.com/DDRMin/GO-Backend/internal/products"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := []products.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}

	json.NewEncoder(w).Encode(products)

}