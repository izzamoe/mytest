package repository

import (
	"gorm.io/gorm"
	"testku/entities"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user
func (r *UserRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

// FindAll returns all users
func (r *UserRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users).Error
	return users, err
}

// FindByID returns a user by ID
func (r *UserRepository) FindByID(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// FindByEmail find by email
func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Update updates a user
func (r *UserRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user
func (r *UserRepository) Delete(user *entities.User) error {
	return r.db.Delete(user).Error
}

// Paginate returns users with pagination
func (r *UserRepository) Paginate(limit int, offset int) ([]entities.User, error) {
	var users []entities.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
