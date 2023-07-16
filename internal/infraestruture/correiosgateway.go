package infrastructure

import (
	"fmt"
)

type CorreiosAPI interface {
	// ObterPrazoEntrega recebe os ceps de origem e destino e o tipo de entrega (industrial ou varejo) e retorna o prazo de entrega
	GetDeliveryTimeAndPrice(zipCodeOrigin, zipCodeDestination string, isIndustrial bool) (*ResponseGateway, error)
}

func NewCorreiosAPI() CorreiosAPI {
	return &CorreiosAPIImpl{}
}

type CorreiosAPIImpl struct{}

func (api *CorreiosAPIImpl) GetDeliveryTimeAndPrice(zipCodeOrigin, zipCodeDestination string, isIndustrial bool) (*ResponseGateway, error) {
	// Aqui, fazemos a chamada à API dos Correios para obter o prazo de entrega
	// Implementação fictícia que retorna um prazo de entrega padrão de 3 dias úteis para o tipo "Correios Industrial"
	// e 5 dias úteis para o tipo "Correios Varejo"

	var response *ResponseApiCorreios
	var err error
	// A chamada à API dos Correios deve ser realizada aqui, substituindo as implementações fictícias abaixo.
	if isIndustrial {
		response, err = api.CorreiosCalculateDeliveryAndPrice(zipCodeOrigin, zipCodeDestination, "industrial")
		if err != nil {
			return nil, err
		}
	} else {
		response, err = api.CorreiosCalculateDeliveryAndPrice(zipCodeOrigin, zipCodeDestination, "varejo")
		if err != nil {
			return nil, err
		}
	}

	return &ResponseGateway{
		Price:        response.ValorFrete,
		PrazoEntrega: response.PrazoEntrega,
	}, nil
}

func (api *CorreiosAPIImpl) CorreiosCalculateDeliveryAndPrice(zipCodeOrigin, zipCodeDestination, user string) (*ResponseApiCorreios, error) {
	// Implementação fictícia que retorna um prazo de entrega padrão de 3 dias úteis para o tipo "Correios Industrial"
	// e 5 dias úteis para o tipo "Correios Varejo"
	switch user {
	case "industrial":
		return &ResponseApiCorreios{
			PrazoEntrega: 3,
			ValorFrete:   10.0,
		}, nil
	case "varejo":
		return &ResponseApiCorreios{
			PrazoEntrega: 5,
			ValorFrete:   20.0,
		}, nil
	default:
		return nil, fmt.Errorf("tipo de usuário inválido")
	}
}
