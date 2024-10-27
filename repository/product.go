package repository

import (
	"gorm.io/gorm"
	"testku/entities"
)

type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create creates a new product
func (r *ProductRepository) Create(product *entities.Product) error {
	return r.db.Create(product).Error
}

// FindAll returns all products
func (r *ProductRepository) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Find(&products).Error
	return products, err
}

// FindByID returns a product by ID
func (r *ProductRepository) FindByID(id uint) (*entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

// Update updates a product
func (r *ProductRepository) Update(product *entities.Product) error {
	return r.db.Save(product).Error
}

// Delete deletes a product
func (r *ProductRepository) Delete(product *entities.Product) error {
	return r.db.Delete(product).Error
}

// FindPriceRange returns products with price range
func (r *ProductRepository) FindPriceRange(min, max int) ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Where("price BETWEEN ? AND ?", min, max).Find(&products).Error
	return products, err
}

// FindByAvailability returns products by availability
func (r *ProductRepository) FindByAvailability(availability bool) ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Where("availability = ?", availability).Find(&products).Error
	return products, err
}

// Paginate returns products with pagination
func (r *ProductRepository) Paginate(limit int, offset int) ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}
