package infrastructure

import "testing"

func TestCorreiosAPIImpl_ObterPrazoEntregaIndustrial(t *testing.T) {
	// Crie uma instância da CorreiosAPIImpl
	correiosAPI := &CorreiosAPIImpl{}

	// Caso 1: Teste com entrega industrial
	cepOrigem := "31035560"
	cepDestino := "66620450"
	isIndustrial := true

	response, err := correiosAPI.GetDeliveryTimeAndPrice(cepOrigem, cepDestino, isIndustrial)
	if err != nil {
		t.Fatalf("Erro ao obter o prazo de entrega: %v", err)
	}

	// Verifique se os valores retornados são os esperados para entrega industrial
	if response.PrazoEntrega != 3 {
		t.Errorf("Prazo de entrega incorreto para entrega industrial. Esperado: 3, Recebido: %d", response.PrazoEntrega)
	}
	if response.Price != 10.0 {
		t.Errorf("Preço incorreto para entrega industrial. Esperado: 10.0, Recebido: %f", response.Price)
	}

	// Caso 2: Teste com entrega de varejo
	isIndustrial = false

	response, err = correiosAPI.GetDeliveryTimeAndPrice(cepOrigem, cepDestino, isIndustrial)
	if err != nil {
		t.Fatalf("Erro ao obter o prazo de entrega: %v", err)
	}

	// Verifique se os valores retornados são os esperados para entrega de varejo
	if response.PrazoEntrega != 5 {
		t.Errorf("Prazo de entrega incorreto para entrega de varejo. Esperado: 5, Recebido: %d", response.PrazoEntrega)
	}
	if response.Price != 20.0 {
		t.Errorf("Preço incorreto para entrega de varejo. Esperado: 20.0, Recebido: %f", response.Price)
	}
}

func TestCorreiosAPIImpl_CalculaPrazoEPreco(t *testing.T) {
	// Crie uma instância da CorreiosAPIImpl
	correiosAPI := &CorreiosAPIImpl{}

	// Caso 1: Teste com usuário industrial
	cepOrigem := "31035560"
	cepDestino := "66620450"
	user := "industrial"

	response, err := correiosAPI.CorreiosCalculateDeliveryAndPrice(cepOrigem, cepDestino, user)
	if err != nil {
		t.Fatalf("Erro ao calcular o prazo e preço: %v", err)
	}

	// Verifique se os valores retornados são os esperados para usuário industrial
	if response.PrazoEntrega != 3 {
		t.Errorf("Prazo de entrega incorreto para usuário industrial. Esperado: 3, Recebido: %d", response.PrazoEntrega)
	}
	if response.ValorFrete != 10.0 {
		t.Errorf("Preço incorreto para usuário industrial. Esperado: 10.0, Recebido: %f", response.ValorFrete)
	}

	// Caso 2: Teste com usuário de varejo
	user = "varejo"

	response, err = correiosAPI.CorreiosCalculateDeliveryAndPrice(cepOrigem, cepDestino, user)
	if err != nil {
		t.Fatalf("Erro ao calcular o prazo e preço: %v", err)
	}

	// Verifique se os valores retornados são os esperados para usuário de varejo
	if response.PrazoEntrega != 5 {
		t.Errorf("Prazo de entrega incorreto para usuário de varejo. Esperado: 5, Recebido: %d", response.PrazoEntrega)
	}
	if response.ValorFrete != 20.0 {
		t.Errorf("Preço incorreto para usuário de varejo. Esperado: 20.0, Recebido: %f", response.ValorFrete)
	}
}
