// internal/application/calculadora.go

package application

import (
	"cmd/api/main.go/internal/domain"
	infrastructure "cmd/api/main.go/internal/infraestruture"
)

// CalculaFreteService represents the freight calculation service
type CalculaFreteService struct {
	correiosAPI infrastructure.CorreiosAPI
}

// NewCalculaFreteService creates an instance of the freight calculation service
func NewCalculaFreteService(correiosAPI infrastructure.CorreiosAPI) *CalculaFreteService {
	return &CalculaFreteService{
		correiosAPI: correiosAPI,
	}
}

// CalculateFreight calculates the freight value based on the request data
func (s *CalculaFreteService) CalculateFreight(request domain.CalculateFreightRequest) (*domain.Freight, error) {
	// Here, we call the Correios API to get the delivery time
	respPrazoEntrega, err := s.correiosAPI.GetDeliveryTimeAndPrice(request.Origin, request.Destination, request.IsIndustrial)
	if err != nil {
		// Error handling (not covered in the example)
	}

	// Creating the freight calculation using the "Frete" Value Object
	return domain.NewFreight(respPrazoEntrega.Price, respPrazoEntrega.PrazoEntrega), nil
}
