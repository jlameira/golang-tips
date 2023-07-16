package application

import (
	"testing"

	"cmd/api/main.go/internal/domain"
	infrastructure "cmd/api/main.go/internal/infraestruture"
)

func TestCalculaFreteIsNotIndustrial(t *testing.T) {
	valorEsperado := 10.0
	// Crie uma instância da CalculaFreteService
	mockCorreiosAPI := &infrastructure.CorreiosAPIMock{}
	calculaFreteService := NewCalculaFreteService(mockCorreiosAPI)

	// Crie um exemplo de request com os dados para o cálculo do frete
	request := domain.CalculateFreightRequest{
		Weight:       10.5,
		Origin:       "Origem A",
		Destination:  "Destino B",
		IsIndustrial: false,
	}

	// Chame a função CalculateFreight para obter o resultado
	resultado, err := calculaFreteService.CalculateFreight(request)
	if err != nil {
		t.Fatalf("Erro ao calcular o frete: %v", err)
	}

	// Verifique se o resultado está correto com base nas suas expectativas

	// Exemplo:
	if resultado.FreightValue != valorEsperado {
		t.Errorf("Valor do frete incorreto. Esperado: %f, Recebido: %f", valorEsperado, resultado.FreightValue)
	}

}

func TestCalculaFreteIsIndustrial(t *testing.T) {
	valorEsperado := 42.0
	// Crie uma instância da CalculaFreteService
	mockCorreiosAPI := &infrastructure.CorreiosAPIMock{}
	calculaFreteService := NewCalculaFreteService(mockCorreiosAPI)

	// Crie um exemplo de request com os dados para o cálculo do frete
	request := domain.CalculateFreightRequest{
		Weight:       10.5,
		Origin:       "Origem A",
		Destination:  "Destino B",
		IsIndustrial: true,
	}

	// Chame a função CalculateFreight para obter o resultado
	resultado, err := calculaFreteService.CalculateFreight(request)
	if err != nil {
		t.Fatalf("Erro ao calcular o frete: %v", err)
	}

	// Verifique se o resultado está correto com base nas suas expectativas

	// Exemplo:
	if resultado.FreightValue != valorEsperado {
		t.Errorf("Valor do frete incorreto. Esperado: %f, Recebido: %f", valorEsperado, resultado.FreightValue)
	}

}
