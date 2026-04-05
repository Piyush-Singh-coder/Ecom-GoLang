package orders

import (
	"net/http"

	"github.com/Piyush-Singh-coder/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request){

	var tempOrder createOrderParams

	if err := json.Read(r, &tempOrder); err != nil{
		json.Write(w, http.StatusBadRequest, err)
		return
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), tempOrder)
	if err != nil{
		json.Write(w, http.StatusInternalServerError, err)
		return
	}

	json.Write(w, http.StatusCreated, createdOrder)
}