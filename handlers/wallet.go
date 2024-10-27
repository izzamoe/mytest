package handlers

import (
	"net/http"
	"testku/helpers"
	"testku/services"

	"github.com/gofiber/fiber/v2"
)

type WalletHandler struct {
	service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) TopUp(c *fiber.Ctx) error {
	// get email from local
	email := c.Locals("email").(string)
	var request struct {
		Amount int `json:"amount"`
	}
	if err := c.BodyParser(&request); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, err.Error())
	}

	err := h.service.TopUp(email, request.Amount)
	if err != nil {
		return helpers.ErrorResponse(c, "Failed to top-up", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Top-up successful", nil, http.StatusOK)
}

func (h *WalletHandler) Withdraw(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	var request struct {
		Amount int `json:"amount"`
	}
	if err := c.BodyParser(&request); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, err.Error())
	}

	err := h.service.Deduct(email, request.Amount)
	if err != nil {
		return helpers.ErrorResponse(c, "Failed to withdraw", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Withdraw successful", nil, http.StatusOK)
}

func (h *WalletHandler) GetBalance(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	balance, err := h.service.GetBalanceByEmail(email)
	if err != nil {
		return helpers.ErrorResponse(c, "Failed to get balance", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Balance retrieved successfully", balance, http.StatusOK)
}
