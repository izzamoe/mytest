package handlers

import (
	"errors"
	"net/http"
	"testku/entities"
	"testku/entities/errs"
	"testku/helpers"
	"testku/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user entities.UserRegisterRequest
	if err := c.BodyParser(&user); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, err.Error())
	}

	confirmPassword := user.ConfirmPassword
	NewUser, err := h.service.Register(&user, confirmPassword)
	if err != nil {
		if errors.Is(err, errs.ErrPasswordMismatch) {
			return helpers.ErrorResponse(c, "Password mismatch", http.StatusBadRequest, err.Error())
		}
		if errors.Is(err, errs.ErrorUserExists) {
			return helpers.ErrorResponse(c, "User already exists", http.StatusConflict, err.Error())
		}
		return helpers.ErrorResponse(c, "Internal server error", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "User registered successfully", NewUser, http.StatusCreated)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var credentials entities.UserLoginRequest
	if err := c.BodyParser(&credentials); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, err.Error())
	}

	user, err := h.service.Login(credentials.Email, credentials.Password)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrInvalidCredentials) {
			return helpers.ErrorResponse(c, "Invalid credentials", http.StatusUnauthorized, err.Error())
		}
		return helpers.ErrorResponse(c, "Internal server error", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Login successful", user, http.StatusOK)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return helpers.ErrorResponse(c, "Internal server error", http.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "Users retrieved successfully", users, http.StatusOK)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid user ID", http.StatusBadRequest, err.Error())
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, "User not found", http.StatusNotFound, err.Error())
	}
	return helpers.SuccessResponse(c, "User retrieved successfully", user, http.StatusOK)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, err.Error())
	}

	if err := h.service.UpdateUser(&user); err != nil {
		return helpers.ErrorResponse(c, "Internal server error", http.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "User updated successfully", user, http.StatusOK)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid user ID", http.StatusBadRequest, err.Error())
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, "User not found", http.StatusNotFound, err.Error())
	}

	if err := h.service.DeleteUser(user); err != nil {
		return helpers.ErrorResponse(c, "Internal server error", http.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(http.StatusNoContent)
}
