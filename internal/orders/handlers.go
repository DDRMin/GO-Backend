package orders

import (
	"log"
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

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.Read(r, &req); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == 0 {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	if len(req.Items) == 0 {
		http.Error(w, "at least one item is required", http.StatusBadRequest)
		return
	}

	orderID, err := h.service.CreateOrder(r.Context(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, CreateOrderResponse{OrderID: orderID})
}
