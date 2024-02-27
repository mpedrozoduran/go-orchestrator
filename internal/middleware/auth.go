package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mpedrozoduran/go-orchestrator/internal/api"
	"github.com/mpedrozoduran/go-orchestrator/internal/config"
	"net/http"
)

type Middleware struct {
	AppConfig config.Config
}

func NewMiddleware(appConfig config.Config) Middleware {
	return Middleware{AppConfig: appConfig}
}

func (m Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &api.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.AppConfig.Auth.SecretKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*api.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		fmt.Println("User:", claims.Username)
		c.Next()
	}
}
