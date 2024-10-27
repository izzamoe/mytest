package services

import (
	"errors"
	"testku/entities"
	"testku/entities/errs"
	"testku/repository"

	"gorm.io/gorm"
)

// OrderService type
type OrderService struct {
	productRepo     *repository.ProductRepository
	walletRepo      *repository.WalletRepository
	transactionRepo *repository.TransactionRepository
	orderRepo       *repository.OrderRepository
}

// NewOrderService creates a new OrderService
func NewOrderService(productRepo *repository.ProductRepository, walletRepo *repository.WalletRepository, transactionRepo *repository.TransactionRepository, orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{
		productRepo:     productRepo,
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
		orderRepo:       orderRepo,
	}
}

func (r *OrderService) CreateOrder(email string, order *entities.Order) error {
	// Check apakah ada product yang di order
	product, err := r.productRepo.FindByID(order.ProductID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrProductNotFound
		}
	}

	// Check apakah stock mencukupi
	if product.Availability < order.Quantity {
		return errs.ErrorInsufficientStock
	}

	// Calculate total price
	order.TotalPrice = product.Price * order.Quantity

	// Get user wallet by email
	wallet, err := r.walletRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	// Check apakah saldo mencukupi
	if wallet.Balance < order.TotalPrice {
		return errs.ErrInsufficientBalance
	}

	// Potong saldo user
	wallet.Balance -= order.TotalPrice
	if err := r.walletRepo.Update(wallet); err != nil {
		return err
	}

	// Buat transaksi
	transaction := entities.NewTransaction(wallet.ID, order.TotalPrice, entities.TransactionTypePurchase, order.CreatedAt.Unix())
	if err := r.transactionRepo.Create(transaction); err != nil {
		return err
	}

	// Update stock product
	product.Availability -= order.Quantity
	if err := r.productRepo.Update(product); err != nil {
		return err
	}

	// Update status order menjadi success
	order.Status = entities.OrderStatusSuccess
	order.TransactionID = transaction.ID

	// Create order
	return r.orderRepo.Create(order)
}

// GetAllOrders returns all orders
func (s *OrderService) GetAllOrders() ([]entities.Order, error) {
	return s.orderRepo.FindAll()
}

// GetOrderByID returns an order by ID
func (s *OrderService) GetOrderByID(id uint) (*entities.Order, error) {
	return s.orderRepo.FindByID(id)
}

// UpdateOrder updates an order
func (s *OrderService) UpdateOrder(order *entities.Order) error {
	return s.orderRepo.Update(order)
}

// DeleteOrder deletes an order
func (s *OrderService) DeleteOrder(order *entities.Order) error {
	return s.orderRepo.Delete(order)
}

// GetOrdersByStatus returns orders by status
func (s *OrderService) GetOrdersByStatus(status string) ([]entities.Order, error) {
	return s.orderRepo.FindByStatus(status)
}

// GetOrdersByUserID returns orders by user ID
func (s *OrderService) GetOrdersByUserID(userID uint) ([]entities.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}

// GetOrdersWithPagination returns orders with pagination for a specific user ID
func (s *OrderService) GetOrdersWithPagination(userID uint, limit int, offset int) ([]entities.Order, error) {
	return s.orderRepo.Paginate(userID, limit, offset)
}

// // PaginateByEmail returns orders with pagination for a specific email
func (s *OrderService) PaginateByEmail(email string, limit int, offset int) ([]entities.Order, error) {
	return s.orderRepo.PaginateByEmail(email, limit, offset)
}

// // GetOrderByTransactionID by email
func (s *OrderService) GetOrderByTransactionIDAndEmail(transactionID uint, email string) (*entities.Order, error) {
	return s.orderRepo.FindOneByTransactionIDAndEmail(transactionID, email)
}
