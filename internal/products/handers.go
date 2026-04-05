package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Piyush-Singh-coder/ecom/internal/json"
	"github.com/go-chi/chi/v5"
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
		log.Println("Error listing products", "error", err)
		json.Write(w, http.StatusInternalServerError, err)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) ProductById(w http.ResponseWriter, r *http.Request){

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Println("Error converting id to int", "error", err)
		json.Write(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.service.ProductById(r.Context(), int64(id))

	if err != nil {
		log.Println("Error getting product by id", "error", err)
		json.Write(w, http.StatusInternalServerError, err)
		return
	}

	json.Write(w, http.StatusOK, product)
}