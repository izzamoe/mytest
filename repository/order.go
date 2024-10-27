package repository

import (
	"testku/entities"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new OrderRepository
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Create creates a new order
func (r *OrderRepository) Create(order *entities.Order) error {
	return r.db.Create(order).Error
}

// FindAll returns all orders
func (r *OrderRepository) FindAll() ([]entities.Order, error) {
	var orders []entities.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

// FindByID returns an order by ID
func (r *OrderRepository) FindByID(id uint) (*entities.Order, error) {
	var order entities.Order
	err := r.db.First(&order, id).Error
	return &order, err
}

// Update updates an order
func (r *OrderRepository) Update(order *entities.Order) error {
	return r.db.Save(order).Error
}

// Delete deletes an order
func (r *OrderRepository) Delete(order *entities.Order) error {
	return r.db.Delete(order).Error
}

// FindByStatus returns orders by status
func (r *OrderRepository) FindByStatus(status string) ([]entities.Order, error) {
	var orders []entities.Order
	err := r.db.Where("status = ?", status).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) FindByUserID(userID uint) ([]entities.Order, error) {
	var orders []entities.Order
	err := r.db.Joins("JOIN transactions ON transactions.id = orders.transaction_id").
		Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Where("wallets.user_id = ?", userID).
		Find(&orders).Error
	return orders, err
}

// Paginate returns orders with pagination for a specific user ID
func (r *OrderRepository) Paginate(userID uint, limit int, offset int) ([]entities.Order, error) {
	var orders []entities.Order
	err := r.db.Joins("JOIN transactions ON transactions.id = orders.transaction_id").
		Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Where("wallets.user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&orders).Error
	return orders, err
}

// find one order by transaction id and email
func (r *OrderRepository) FindOneByTransactionIDAndEmail(transactionID uint, email string) (*entities.Order, error) {
	var order entities.Order
	err := r.db.Joins("JOIN transactions ON transactions.id = orders.transaction_id").
		Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Joins("JOIN users ON users.id = wallets.user_id").
		Where("transactions.id = ? AND users.email = ?", transactionID, email).
		First(&order).Error
	return &order, err
}

// PaginateByEmail returns orders with pagination for a specific email
func (r *OrderRepository) PaginateByEmail(email string, limit int, offset int) ([]entities.Order, error) {
	var orders []entities.Order
	err := r.db.Joins("JOIN transactions ON transactions.id = orders.transaction_id").
		Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Joins("JOIN users ON users.id = wallets.user_id").
		Where("users.email = ?", email).
		Limit(limit).
		Offset(offset).
		Find(&orders).Error
	return orders, err
}
