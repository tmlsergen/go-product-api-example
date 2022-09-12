package main

import (
	"app/cache"
	"app/configs"
	"app/handlers"
	"app/middlewares"
	"app/repositories"
	"app/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR ON LOAD ENV")
		os.Exit(500)
	}

	appRoute := fiber.New()

	db, err := configs.ConnectDb()

	rds := cache.Connection()

	if err != nil {
		fmt.Println("ERROR ON CONNECT DB")
		os.Exit(500)
	}

	// generate services
	// user
	userRepository := repositories.NewUserRepositoryDb(db)
	userService := services.NewUserService(userRepository)
	auth := handlers.AuthHandler{UserService: userService}

	// products
	productRepository := repositories.NewProductRepositoryDb(db)
	deliveryOptionRepository := repositories.NewDeliveryOptionDb(db)
	productService := services.NewProductService(productRepository, deliveryOptionRepository, rds)
	productHandler := handlers.ProductHandler{ProductService: productService}

	// routes
	v1 := appRoute.Group("/api/v1")
	authGroup := v1.Group("/auth")
	productGroup := v1.Group("/products")

	// auth actions
	authGroup.Post("/register", auth.Register)
	authGroup.Post("/login", auth.Login)

	// product actions
	productGroup.Post("/", middlewares.ValidateAuthRequest, productHandler.CreateProduct)
	productGroup.Put("/:id", middlewares.ValidateAuthRequest, productHandler.UpdateProduct)

	productGroup.Get("/", productHandler.ListProducts)
	productGroup.Get("/:id", productHandler.ShowProduct)

	appRoute.Listen(":4000")
}
