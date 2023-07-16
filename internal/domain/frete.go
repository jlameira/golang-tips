package domain

// Weight represents the weight of the merchandise
type Weight float64

// Origin represents the origin location of the merchandise
type Origin string

// Destination represents the destination location of the merchandise
type Destination string

// Freight represents the freight calculation
type Freight struct {
	FreightValue float64
	DeliveryTime int
}

// NewFreight creates a new instance of Freight
func NewFreight(freightValue float64, deliveryTime int) *Freight {
	return &Freight{
		FreightValue: freightValue,
		DeliveryTime: deliveryTime,
	}
}

// CalculateFreightRequest represents the request body for the freight calculation service
type CalculateFreightRequest struct {
	Weight       float64 `json:"weight"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	IsIndustrial bool    `json:"isIndustrial"`
}
