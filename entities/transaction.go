package entities

import "gorm.io/gorm"

// Transaction represents a transaction in the system
type Transaction struct {
	gorm.Model      `json:"-"`
	WalletID        uint   `json:"-"` // Foreign key to Wallet
	Amount          int    `json:"amount"`
	TransactionType string `json:"transaction_type"`                                        // Enum: deposit, withdrawal
	Timestamp       int64  `json:"timestamp"`                                               // Unix timestamp for the transaction time
	Wallet          Wallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"` // Many-to-One relationship with Wallet
}

// make enum for transaction type
const (
	TransactionTypeDeposit    = "deposit"
	TransactionTypeWithdrawal = "withdrawal"
	TransactionTypePurchase   = "purchase"
)

// NewTransaction creates a new transaction
func NewTransaction(walletID uint, amount int, transactionType string, timestamp int64) *Transaction {
	return &Transaction{
		WalletID:        walletID,
		Amount:          amount,
		TransactionType: transactionType,
		Timestamp:       timestamp,
	}
}
