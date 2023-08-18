package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"ocs-app/handlers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// 验证 JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return "jwtSecret", nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中，供后续处理函数使用
		claims := token.Claims.(jwt.MapClaims)
		if handlers.RevokedTokens[tokenString] {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized (token revoked)"})
			c.Abort()
			return
		}
		c.Set("username", claims["username"])
		c.Set("admin", claims["admin"])
		c.Next()
	}
}

func AuthModifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		admin, _ := c.Get("admin")
		if admin != true {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}
	}
}
