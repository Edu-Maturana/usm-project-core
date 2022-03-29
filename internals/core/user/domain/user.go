package domain

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint8  `json:"stock"`
	Price       uint16 `json:"price"`
}
