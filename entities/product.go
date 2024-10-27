package entities

import "time"

type Product struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Price        int       `json:"price"`
	Description  string    `json:"description"`
	Availability int       `json:"availability"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// NewProduct creates a new product
func NewProduct(name, description string, price int, availability int) *Product {
	return &Product{
		Name:         name,
		Price:        price,
		Description:  description,
		Availability: availability,
	}
}
