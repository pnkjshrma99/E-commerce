// server.go
package main

import (
	"log"

	"./middleware"
	"./models"
	"example.com/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to the database
	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize controllers
	productController := controllers.NewProductController(db)
	userController := controllers.NewUserController(db)
	cartController := controllers.NewCartController(db)
	orderController := controllers.NewOrderController(db)

	// Apply middleware
	router.Use(middleware.AuthMiddleware())

	// Product routes
	router.GET("/products", productController.ListProducts)
	router.POST("/products", middleware.AuthorizeRole("admin"), productController.CreateProduct)
	// Add more product routes...

	// User routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	// Add more user routes...

	// Cart routes
	router.GET("/cart", cartController.GetCart)
	router.POST("/cart", cartController.AddToCart)
	// Add more cart routes...

	// Order routes
	router.POST("/order", orderController.PlaceOrder)
	router.GET("/orders", orderController.ListOrders)
	// Add more order routes...

	router.Run(":8080")
}
