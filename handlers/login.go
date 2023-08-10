package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	var isAdmin bool
	if username == "Tom" {
		isAdmin = true
	} else if username == "Mary" {
		isAdmin = false
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}

	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"admin":    isAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间
	})

	// 签署 JWT
	tokenString, err := token.SignedString("jwtSecret")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// 全局变量，存储已撤销的 JWT 标识
// 这里是为了简化，正规项目当然是放到 Redis 或者数据库中
var RevokedTokens = make(map[string]bool)

func RevokeToken(tokenID string) {
	RevokedTokens[tokenID] = true
}

func Logout(c *gin.Context) {
	tokenString := c.Param("token")
	if len(tokenString) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}
	RevokeToken(tokenString)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
