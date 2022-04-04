package domain

type Order struct {
	ID            string    `json:"id"`
	CustomerEmail string    `json:"customer_email" validate:"required, email"`
	Products      []Product `json:"products" validate:"required"`
	Status        string    `json:"status" enum:"pending,confirmed,rejected" validate:"required"`
	CreatedAt     string    `json:"created_at"`
}

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Price    uint16 `json:"price" validate:"required,gte=0"`
	Quantity uint8  `json:"quantity" validate:"required,gte=1"`
}
