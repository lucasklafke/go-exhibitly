package main

import (
	"exhibitly/config"
	"exhibitly/routes"
	"log"
	"os"

	_ "exhibitly/docs"
	"exhibitly/ent"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Exhibitly API
// @version 1.0
// @description API de classificados online para o projeto Exhibitly.
// @termsOfService http://exhibitly.com/terms/

// @contact.name Suporte Exhibitly
// @contact.url http://exhibitly.com/support
// @contact.email suporte@exhibitly.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env, usando variáveis do sistema")
	}

	config.ConnectToDatabase()

	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))

	router := gin.Default()
	routes.SetupRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":8080"))
}
