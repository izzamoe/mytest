package services

import (
	"testku/entities"
	"testku/repository"
)

// TransactionService type
type TransactionService struct {
	repo *repository.TransactionRepository
}

// NewTransactionService creates a new TransactionService
func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// CreateTransaction creates a new transaction
func (s *TransactionService) CreateTransaction(transaction *entities.Transaction) error {
	return s.repo.Create(transaction)
}

// GetAllTransactions returns all transactions
func (s *TransactionService) GetAllTransactions() ([]entities.Transaction, error) {
	return s.repo.FindAll()
}

// GetTransactionByID returns a transaction by ID
func (s *TransactionService) GetTransactionByID(id uint) (*entities.Transaction, error) {
	return s.repo.FindByID(id)
}

// UpdateTransaction updates a transaction
func (s *TransactionService) UpdateTransaction(transaction *entities.Transaction) error {
	return s.repo.Update(transaction)
}

// DeleteTransaction deletes a transaction
func (s *TransactionService) DeleteTransaction(transaction *entities.Transaction) error {
	return s.repo.Delete(transaction)
}

// GetTransactionsByWalletID returns transactions by wallet ID
func (s *TransactionService) GetTransactionsByWalletID(walletID uint) ([]entities.Transaction, error) {
	return s.repo.FindByWalletID(walletID)
}

// GetTransactionsByUserID returns transactions by user ID
func (s *TransactionService) GetTransactionsByUserID(userID uint) ([]entities.Transaction, error) {
	return s.repo.FindByUserID(userID)
}

// GetTransactionsWithPagination returns transactions with pagination
func (s *TransactionService) GetTransactionsWithPagination(limit int, offset int) ([]entities.Transaction, error) {
	return s.repo.Paginate(limit, offset)
}

// GetTransactionsByEmail returns transactions by user email
func (s *TransactionService) GetTransactionsByEmail(email string) ([]entities.Transaction, error) {
	return s.repo.FindByEmail(email)
}

// GetTransactionsByEmailWithPagination returns transactions by user email with pagination
func (s *TransactionService) GetTransactionsByEmailWithPagination(email string, limit int, offset int) ([]entities.Transaction, error) {
	return s.repo.FindByEmailWithPagination(email, limit, offset)
}
