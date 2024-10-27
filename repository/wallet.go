package repository

import (
	"testku/entities"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

// NewWalletRepository creates a new WalletRepository
func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

// Create creates a new wallet
func (r *WalletRepository) Create(wallet *entities.Wallet) error {
	return r.db.Create(wallet).Error
}

// FindAll returns all wallets
func (r *WalletRepository) FindAll() ([]entities.Wallet, error) {
	var wallets []entities.Wallet
	err := r.db.Find(&wallets).Error
	return wallets, err
}

// FindByID returns a wallet by ID
func (r *WalletRepository) FindByID(id uint) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := r.db.First(&wallet, id).Error
	return &wallet, err
}

// FindByUserID find wallet by user id
func (r *WalletRepository) FindByUserID(userID uint) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := r.db.Where("user_id = ?", userID).First(&wallet).Error
	return &wallet, err
}

// FindBalanceByUserID find balance by user id
func (r *WalletRepository) FindBalanceByUserID(userID uint) (int, error) {
	var wallet entities.Wallet
	err := r.db.Where("user_id = ?", userID).First(&wallet).Error
	return wallet.Balance, err
}

// Update updates a wallet
func (r *WalletRepository) Update(wallet *entities.Wallet) error {
	return r.db.Save(wallet).Error
}

// FindByEmail finds a wallet by the user's email
func (r *WalletRepository) FindByEmail(email string) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := r.db.Joins("JOIN users ON users.id = wallets.user_id").Where("users.email = ?", email).First(&wallet).Error
	return &wallet, err
}

// FindBalanceByEmail finds the balance by the user's email
func (r *WalletRepository) FindBalanceByEmail(email string) (int, error) {
	var wallet entities.Wallet
	err := r.db.Joins("JOIN users ON users.id = wallets.user_id").Where("users.email = ?", email).First(&wallet).Error
	return wallet.Balance, err
}
