package handlers

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"testku/entities"
	"testku/entities/errs"
	"testku/helpers"
	"testku/services"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entities.Order
	if err := c.BodyParser(&order); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, "Invalid request")
	}

	email := c.Locals("email").(string) // Assuming email is set in context by middleware

	if err := h.service.CreateOrder(email, &order); err != nil {
		// if
		if errors.Is(err, errs.ErrProductNotFound) {
			return helpers.ErrorResponse(c, "Product not found", http.StatusNotFound, err.Error())
		}
		if errors.Is(err, errs.ErrorInsufficientStock) {
			return helpers.ErrorResponse(c, "Insufficient stock", http.StatusBadRequest, err.Error())
		}

		return helpers.ErrorResponse(c, "Failed to order", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Order created successfully", order, http.StatusCreated)
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid limit parameter", http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid offset parameter", http.StatusBadRequest, err.Error())
	}

	orders, err := h.service.PaginateByEmail(email, limit, offset)
	if err != nil {
		return helpers.ErrorResponse(c, "Failed to retrieve orders", http.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "Orders retrieved successfully", orders, http.StatusOK)
}

func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	id, err := c.ParamsInt("id")
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid order ID", http.StatusBadRequest, err.Error())
	}

	order, err := h.service.GetOrderByTransactionIDAndEmail(uint(id), email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ErrorResponse(c, "Order not found", http.StatusNotFound, nil)
		}
		return helpers.ErrorResponse(c, "Order not found", http.StatusNotFound, err.Error())
	}
	return helpers.SuccessResponse(c, "Order retrieved successfully", order, http.StatusOK)
}
