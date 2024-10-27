package services

import (
	"testku/entities"
	"testku/entities/errs"
	"testku/repository"
	"time"
)

// WalletService type
type WalletService struct {
	walletRepo      *repository.WalletRepository
	transactionRepo *repository.TransactionRepository
}

// NewWalletService creates a new WalletService
func NewWalletService(walletRepo *repository.WalletRepository, transactionRepo *repository.TransactionRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo, transactionRepo: transactionRepo}
}

// TopUp adds amount to the wallet balance
func (s *WalletService) TopUp(email string, amount int) error {
	wallet, err := s.walletRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	wallet.Balance += amount
	if err := s.walletRepo.Update(wallet); err != nil {
		return err
	}

	// Create a new transaction
	transaction := entities.NewTransaction(wallet.ID, amount, entities.TransactionTypeDeposit, time.Now().Unix())
	return s.transactionRepo.Create(transaction)
}

// Deduct deducts amount from the wallet balance
func (s *WalletService) Deduct(email string, amount int) error {
	wallet, err := s.walletRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	if wallet.Balance < amount {
		return errs.ErrInsufficientBalance
	}
	wallet.Balance -= amount
	if err := s.walletRepo.Update(wallet); err != nil {
		return err
	}

	// Create a new transaction
	transaction := entities.NewTransaction(wallet.ID, amount, entities.TransactionTypeWithdrawal, time.Now().Unix())
	return s.transactionRepo.Create(transaction)
}

// GetBalance returns the wallet balance
func (s *WalletService) GetBalance(userID uint) (int, error) {
	return s.walletRepo.FindBalanceByUserID(userID)
}

// get balance by email
func (s *WalletService) GetBalanceByEmail(email string) (int, error) {
	return s.walletRepo.FindBalanceByEmail(email)
}
