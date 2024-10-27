package entities

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model `json:"-"`
	UserID     uint `json:"user_id"` // Foreign key to User
	Balance    int  `json:"balance"`
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-One relationship with User
}

// NewWallet creates a new wallet
func NewWallet(userID uint, balance int) *Wallet {
	return &Wallet{
		UserID:  userID,
		Balance: balance,
	}
}
