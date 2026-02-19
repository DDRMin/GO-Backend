package products

import (
	"log"
	"net/http"
	"strconv"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
	"github.com/DDRMin/GO-Backend/internal/json"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to list products", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)

}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindProductByID(r.Context(), intID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to find product", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req createProduct

	if err := json.Read(r, &req); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var price pgtype.Numeric
	priceStr := strconv.FormatFloat(req.Price, 'f', -1, 64)
	if err := price.Scan(priceStr); err != nil {
		log.Println(err)
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	params := repo.CreateProductParams{
		Name:     req.Name,
		Price:    price,
		Quantity: req.Quantity,
	}

	id, err := h.service.CreateProduct(r.Context(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, map[string]int64{"id": id})
}
