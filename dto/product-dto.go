package dto

type CreateProductDto struct {
	Category     string  `json:"category" validate:"required" example:"Electronics"`
	Title        string  `json:"title" validate:"required" example:"Smartphone"`
	Price        float64 `json:"price" validate:"required,gt=0" example:"999.99"`
	Description  string  `json:"description" validate:"required" example:"A great smartphone"`
	PostalCode   string  `json:"postal_code" validate:"required" example:"12345-678"`
	Image        string  `json:"image" validate:"required" example:"smartphone.jpg"`
	Street       string  `json:"street" validate:"required" example:"Rua Exemplo"`
	Number       string  `json:"number" validate:"required" example:"123"`
	Complement   string  `json:"complement" example:"Apt 101"`
	Neighborhood string  `json:"neighborhood" validate:"required" example:"Centro"`
	City         string  `json:"city" validate:"required" example:"São Paulo"`
	State        string  `json:"state" validate:"required" example:"SP"`
	Country      string  `json:"country" validate:"required" example:"Brasil"`
	AuthorID     uint    `json:"author_id" validate:"required" example:"1"`
}
