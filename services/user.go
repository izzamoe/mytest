package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strconv"
	"testku/entities"
	"testku/entities/errs"
	"testku/repository"
	"time"
)

// UserService Type user services
type UserService struct {
	userRepository *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

func (s *UserService) Register(user *entities.UserRegisterRequest, confirmPassword string) (*entities.User, error) {
	if user.Password != confirmPassword {
		return nil, errs.ErrPasswordMismatch
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	newUser := entities.NewUser(user.Name, user.Email, user.Password)
	err = s.userRepository.Create(newUser)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errs.ErrorUserExists
		}
		return nil, err
	}
	return newUser, nil
}

// Login verifies user credentials
func (s *UserService) Login(email, password string) (*entities.LoginResponse, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errs.ErrInvalidCredentials
	}

	expDuration, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		return nil, err
	}
	// Generate JWT token
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * time.Duration(expDuration)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	response := entities.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: t,
	}

	return &response, nil
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.userRepository.FindAll()
}

// GetUserByID returns a user by ID
func (s *UserService) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepository.FindByID(id)
}

// GetUserByEmail returns a user by email
func (s *UserService) GetUserByEmail(email string) (*entities.User, error) {
	return s.userRepository.FindByEmail(email)
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(user *entities.User) error {
	return s.userRepository.Update(user)
}

func (s *UserService) DeleteUser(user *entities.User) error {
	return s.userRepository.Delete(user)
}
