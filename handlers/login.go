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
