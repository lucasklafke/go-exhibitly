package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes centraliza todas as rotas da aplicação
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api") // Prefixo para as rotas da API

	// Registrar sub-rotas de diferentes módulos
	SetupProductRoutes(api) // Exemplo: rotas de usuários
	// Adicione outros módulos aqui, como SetupProductRoutes(api)
}
