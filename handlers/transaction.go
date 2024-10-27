package handlers

import (
	"net/http"
	"strconv"
	"testku/helpers"
	"testku/services"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	transactions, err := h.service.GetTransactionsByEmail(email)
	if err != nil {
		return helpers.ErrorResponse(c, "Failed to fetch transactions", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Transactions retrieved successfully", transactions, http.StatusOK)
}

func (h *TransactionHandler) GetTransactionsWithPagination(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid limit parameter", http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid offset parameter", http.StatusBadRequest, err.Error())
	}

	transactions, err := h.service.GetTransactionsByEmailWithPagination(email, limit, offset)
	if err != nil {
		return helpers.ErrorResponse(c, "Failed to fetch transactions with pagination", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Transactions retrieved successfully", transactions, http.StatusOK)
}
