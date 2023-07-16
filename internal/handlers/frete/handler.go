package frete

import (
	"encoding/json"
	"net/http"

	"cmd/api/main.go/internal/domain"
)

type FreightHandler struct {
	calculatorService CalculatorService
}
type CalculatorService interface {
	CalculateFreight(request domain.CalculateFreightRequest) (*domain.Freight, error)
}

func NewHandlerFrete(calculatorService CalculatorService) *FreightHandler {
	return &FreightHandler{
		calculatorService,
	}
}

// CalculateFreightHandler is the handler for the /calcula-frete route
func (h *FreightHandler) CalculateFreightHandler(w http.ResponseWriter, r *http.Request) {
	var request domain.CalculateFreightRequest

	// Decode the request body into the CalculateFreightRequest struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error decoding the request body", http.StatusBadRequest)
		return
	}

	// Here you can create an instance of the application service and use it to calculate the freight
	// calculatorService := application.NewCalculateFreightService()
	response, err := h.calculatorService.CalculateFreight(request)
	if err != nil {
		http.Error(w, "Error calculating the freight", http.StatusInternalServerError)
		return
	}

	// Convert the response to JSON and write it to the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}
}
