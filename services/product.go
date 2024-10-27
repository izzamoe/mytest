package services

import (
	"testku/entities"
	"testku/repository"
)

// ProductService type
type ProductService struct {
	repo *repository.ProductRepository
}

// NewProductService creates a new ProductService
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(product *entities.Product) error {
	return s.repo.Create(product)
}

// GetAllProducts returns all products
func (s *ProductService) GetAllProducts() ([]entities.Product, error) {
	return s.repo.FindAll()
}

// GetProductByID returns a product by ID
func (s *ProductService) GetProductByID(id uint) (*entities.Product, error) {
	return s.repo.FindByID(id)
}

// UpdateProduct updates a product
func (s *ProductService) UpdateProduct(product *entities.Product) error {
	return s.repo.Update(product)
}

// DeleteProduct deletes a product
func (s *ProductService) DeleteProduct(product *entities.Product) error {
	return s.repo.Delete(product)
}

// GetProductsByPriceRange returns products within a price range
func (s *ProductService) GetProductsByPriceRange(min, max int) ([]entities.Product, error) {
	return s.repo.FindPriceRange(min, max)
}

// GetProductsByAvailability returns products by availability
func (s *ProductService) GetProductsByAvailability(availability bool) ([]entities.Product, error) {
	return s.repo.FindByAvailability(availability)
}
