package repository

import (
	"testku/entities"

	"gorm.io/gorm"
)

// TransactionRepository is a struct to store db
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new TransactionRepository
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create creates a new transaction
func (r *TransactionRepository) Create(transaction *entities.Transaction) error {
	return r.db.Create(transaction).Error
}

// FindAll returns all transactions
func (r *TransactionRepository) FindAll() ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

// FindByWalletID find transaction by wallet id
func (r *TransactionRepository) FindByWalletID(walletID uint) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Where("wallet_id = ?", walletID).Find(&transactions).Error
	return transactions, err
}

// FindByUserID finds transactions by user ID
func (r *TransactionRepository) FindByUserID(userID uint) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Where("wallets.user_id = ?", userID).
		Find(&transactions).Error
	return transactions, err
}

// FindByID returns a transaction by ID
func (r *TransactionRepository) FindByID(id uint) (*entities.Transaction, error) {
	var transaction entities.Transaction
	err := r.db.First(&transaction, id).Error
	return &transaction, err
}

// Update updates a transaction
func (r *TransactionRepository) Update(transaction *entities.Transaction) error {
	return r.db.Save(transaction).Error
}

// Delete deletes a transaction
func (r *TransactionRepository) Delete(transaction *entities.Transaction) error {
	return r.db.Delete(transaction).Error
}

// Paginate returns transactions with pagination
func (r *TransactionRepository) Paginate(limit int, offset int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Limit(limit).Offset(offset).Find(&transactions).Error
	return transactions, err
}

// FindByEmail finds transactions by user email
func (r *TransactionRepository) FindByEmail(email string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Joins("JOIN users ON users.id = wallets.user_id").
		Where("users.email = ?", email).
		Find(&transactions).Error
	return transactions, err
}

// FindByEmailWithPagination finds transactions by user email with pagination
func (r *TransactionRepository) FindByEmailWithPagination(email string, limit int, offset int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Joins("JOIN users ON users.id = wallets.user_id").
		Where("users.email = ?", email).
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error
	return transactions, err
}
