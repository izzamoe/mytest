package entities

import (
	"gorm.io/gorm"
)

// User
type User struct {
	gorm.Model `json:"-"`
	Name       string  `json:"name"`
	Email      string  `json:"email" gorm:"unique"`
	Password   string  `json:"-"`
	Role       string  `json:"role" gorm:"default:user"`
	Wallet     *Wallet `gorm:"foreignKey:UserID" json:"-"`
}

// request user
type UserRegisterRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// response register
type UserRegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// UserLoginRequest user login request
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse response login
type LoginResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

// NewUser new user
func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Wallet = &Wallet{
		Balance: 0,
		UserID:  u.ID,
	}
	return nil
}

// BeforeDelete hook
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.Wallet != nil {
		if err := tx.Delete(u.Wallet).Error; err != nil {
			return err
		}
	}
	return nil
}

// Role admin or user
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
