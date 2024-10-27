package routes

import (
	"testku/handlers"
	"testku/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App, transactionHandler *handlers.TransactionHandler) {
	// Grouping routes under /api/v1/transactions
	transactionGroup := app.Group("/api/v1/transactions")
	transactionGroup.Use(middlewares.JWTMiddleware())

	transactionGroup.Get("/", transactionHandler.GetAllTransactions).Name("getAllTransactions")
	transactionGroup.Get("/paginate", transactionHandler.GetTransactionsWithPagination).Name("getTransactionsWithPagination")
}
