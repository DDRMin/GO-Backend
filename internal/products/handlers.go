package products

import (
	"net/http"

	"github.com/DDRMin/GO-Backend/internal/json"
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

	err := h.service.ListProducts(r.Context())
	if err != nil {
		json.Write(w, http.StatusInternalServerError, map[string]string{"error": "failed to list products"})
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}	

	json.Write(w, http.StatusOK, products)

}