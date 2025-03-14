package dto

type CreateProductDto struct {
	Category    string  `json:"category" validate:"required" example:"Electronics"`
	Title       string  `json:"title" validate:"required" example:"Smartphone"`
	Price       float64 `json:"price" validate:"required,gt=0" example:"999.99"`
	Description string  `json:"description"  example:"A great smartphone"`
	PostalCode  string  `json:"postal_code" example:"12345-678"`
	Image       string  `json:"image"  example:"smartphone.jpg"`
	City        string  `json:"city"  example:"São Paulo"`
	State       string  `json:"state"  example:"SP"`
}
