package routes

import (
	"testku/handlers"
	"testku/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WalletRoutes(app *fiber.App, walletHandler *handlers.WalletHandler) {
	// Grouping routes under /api/v1/wallet
	walletGroup := app.Group("/api/v1/wallet")

	walletGroup.Use(middlewares.JWTMiddleware())
	walletGroup.Post("/topup", walletHandler.TopUp)
	walletGroup.Post("/withdraw", walletHandler.Withdraw)
	walletGroup.Get("/balance", walletHandler.GetBalance)
}
