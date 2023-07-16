package infrastructure

// CorreiosAPIMock é uma implementação de mock para a interface CorreiosAPI
type CorreiosAPIMock struct{}

// ObterPrazoEntrega retorna um resultado de mock para a API dos Correios
func (m *CorreiosAPIMock) GetDeliveryTimeAndPrice(cepOrigem, cepDestino string, isIndustrial bool) (*ResponseGateway, error) {

	if isIndustrial {
		return &ResponseGateway{
			Price:        42.0,
			PrazoEntrega: 3,
		}, nil
	}
	return &ResponseGateway{
		Price:        10.0,
		PrazoEntrega: 3,
	}, nil
}
