package routes

import (
	"exhibitly/config"
	"exhibitly/controllers"
	"exhibitly/repositories"
	"exhibitly/services"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes define as rotas relacionadas a usuários
func SetupProductRoutes(router *gin.RouterGroup) {
	productRoutes := router.Group("/products")
	productRepository := repositories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	{
		// @Router
		productRoutes.GET("/", productController.GetProducts)    // get products
		productRoutes.POST("/", productController.CreateProduct) // create product
		productRoutes.GET("/:id", productController.GetProduct)  // get product
	}
}
