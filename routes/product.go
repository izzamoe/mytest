package routes

import (
	"testku/handlers"
	"testku/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, productHandler *handlers.ProductHandler) {
	// Grouping routes under /api/v1/products
	productGroup := app.Group("/api/v1/products")
	productGroup.Get("/", productHandler.GetAllProducts).Name("getAllProducts")
	productGroup.Get("/:id", productHandler.GetProductByID).Name("getProductByID")

	// set for admin only
	productGroup.Use(middlewares.JWTMiddleware(), middlewares.AdminOnly())
	productGroup.Post("/", productHandler.CreateProduct).Name("createProduct")
	productGroup.Put("/:id", productHandler.UpdateProduct).Name("updateProduct")
	productGroup.Delete("/:id", productHandler.DeleteProduct).Name("deleteProduct")
}
