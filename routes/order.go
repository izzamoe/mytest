package routes

import (
	"testku/handlers"
	"testku/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App, handler *handlers.OrderHandler) {
	order := app.Group("/api/v1/orders")
	order.Use(middlewares.JWTMiddleware()) // Protect all routes in /api/v1/orders with JWT middleware
	order.Post("/", handler.CreateOrder)
	order.Get("/", handler.GetAllOrders)
	order.Get("/:id", handler.GetOrderByID)
}
