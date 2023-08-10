package router

import (
	"word-card-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		r.POST("/login", handlers.Login)

		v1.GET("/word-cards", handlers.GetWordCards)
		v1.POST("/word-cards", handlers.CreateWordCard)
		v1.PUT("/word-cards/:id", handlers.UpdateWordCard)
		v1.DELETE("/word-cards/:id", handlers.DeleteWordCard)
	}

}
