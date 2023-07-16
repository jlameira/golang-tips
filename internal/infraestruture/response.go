package infrastructure

type ResponseGateway struct {
	Price        float64
	PrazoEntrega int
}

type ResponseApiCorreios struct {
	ValorFrete   float64 `json:"valor_frete"`
	PrazoEntrega int     `json:"prazo_entrega"`
}
