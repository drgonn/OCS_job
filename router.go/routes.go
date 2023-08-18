package router

import (
	"net/http"
	"ocs-app/handlers"
	"ocs-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Heath Check OK 1")
	})

	r.POST("/login", handlers.Login)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{

		v1.GET("/stocks", handlers.GetWordCards)
	}

}
