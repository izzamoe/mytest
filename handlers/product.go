package handlers

import (
	"net/http"
	"testku/entities"
	"testku/helpers"
	"testku/services"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, "Invalid request")
	}

	if err := h.service.CreateProduct(&product); err != nil {
		return helpers.ErrorResponse(c, "Failed to create product", http.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Product created successfully", product, http.StatusCreated)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.service.GetAllProducts()
	if err != nil {
		return helpers.ErrorResponse(c, "Failed to retrieve products", http.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "Products retrieved successfully", products, http.StatusOK)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid product ID", http.StatusBadRequest, err.Error())
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, "Product not found", http.StatusNotFound, err.Error())
	}
	return helpers.SuccessResponse(c, "Product retrieved successfully", product, http.StatusOK)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid product ID", http.StatusBadRequest, err.Error())
	}

	var product entities.Product
	product.ID = uint(id)
	if err := c.BodyParser(&product); err != nil {
		return helpers.ErrorResponse(c, "Invalid request", http.StatusBadRequest, err.Error())
	}

	if err := h.service.UpdateProduct(&product); err != nil {
		return helpers.ErrorResponse(c, "Failed to update product", http.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "Product updated successfully", product, http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return helpers.ErrorResponse(c, "Invalid product ID", http.StatusBadRequest, err.Error())
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, "Product not found", http.StatusNotFound, err.Error())
	}

	if err := h.service.DeleteProduct(product); err != nil {
		return helpers.ErrorResponse(c, "Failed to delete product", http.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "Product deleted successfully", nil, http.StatusOK)
}
