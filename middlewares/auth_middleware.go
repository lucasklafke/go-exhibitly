package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4" // Certifique-se de ter o pacote JWT instalado
)

// AuthorizationMiddleware é um middleware para verificar o token JWT
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtendo o token do cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Verificando se o token está no formato correto
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// Validando o token JWT
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Aqui você retorna a chave secreta para validar o token
			return []byte("SUA_CHAVE_SECRETA"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			c.Abort()
			return
		}

		// Salvando as informações do usuário no contexto para usar nas rotas
		c.Set("userID", claims["sub"]) // `sub` é um campo comum em tokens JWT
		c.Next()
	}
}
