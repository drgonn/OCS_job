package router

import (
	"word-card-app/handlers"
	"word-card-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", handlers.Login)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{

		v1.GET("/word-cards", handlers.GetWordCards)
		v1.POST("/word-cards", middleware.AuthModifyMiddleware(), handlers.CreateWordCard)
		v1.PUT("/word-cards/:id", middleware.AuthModifyMiddleware(), handlers.UpdateWordCard)
		v1.DELETE("/word-cards/:id", middleware.AuthModifyMiddleware(), handlers.DeleteWordCard)
	}

}
