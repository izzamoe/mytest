package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model    `json:"-"`
	ProductID     uint        `json:"product_id"`                        // Foreign key referring to the Product
	TransactionID uint        `json:"-"`                                 // Foreign key referring to the Transaction
	Quantity      int         `json:"quantity"`                          // Quantity of the product purchased
	TotalPrice    int         `json:"total_price"`                       // Total price for the transaction
	Status        string      `json:"status"`                            // Enum (e.g., success, failed)
	Product       Product     `gorm:"foreignKey:ProductID" json:"-"`     // Relationship with Product
	Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"-"` // Relationship with Transaction
}

// make enum for status
const (
	OrderStatusSuccess = "success"
	OrderStatusFailed  = "failed"
)
