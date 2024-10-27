package main

import (
	"log"
	"os"
	"testku/config"
	"testku/handlers"
	"testku/repository"
	"testku/routes"
	"testku/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	db, err := config.InitializeDB()
	if err != nil {
		return
	}

	// add repository and handler initialization here
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	// add service initialization here
	userService := services.NewUserService(userRepository)
	productService := services.NewProductService(productRepository)
	walletService := services.NewWalletService(walletRepository, transactionRepository)
	transactionService := services.NewTransactionService(transactionRepository)
	orderService := services.NewOrderService(productRepository, walletRepository, transactionRepository, orderRepository)

	// add handler initialization here
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	walletHandler := handlers.NewWalletHandler(walletService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// add routes initialization here
	routes.UserRoutes(app, userHandler)
	routes.ProductRoutes(app, productHandler)
	routes.WalletRoutes(app, walletHandler)
	routes.TransactionRoutes(app, transactionHandler)
	routes.OrderRoutes(app, orderHandler)

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
