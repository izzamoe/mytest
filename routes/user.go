package routes

import (
	"testku/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	// Grouping routes under /api/v1/users
	userGroup := app.Group("/api/v1/users")

	userGroup.Post("/register", userHandler.Register).Name("register")
	userGroup.Post("/login", userHandler.Login).Name("login")
	//userGroup.Get("/", userHandler.GetAllUsers)
	//userGroup.Get("/:id", userHandler.GetUserByID)
	//userGroup.Put("/:id", userHandler.UpdateUser)
	//userGroup.Delete("/:id", userHandler.DeleteUser)
}
