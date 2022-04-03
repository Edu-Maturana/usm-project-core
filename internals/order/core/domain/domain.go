package domain

type Order struct {
	ID            string    `json:"id"`
	CustomerEmail string    `json:"customer_email"`
	Products      []Product `json:"products"`
	Status        string    `json:"status" enum:"pending,confirmed,rejected"`
	CreatedAt     string    `json:"created_at"`
}

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    uint16 `json:"price"`
	Quantity uint8  `json:"quantity"`
}
