package router

import (
	"net/http"
	"ocs-app/handlers"
	"ocs-app/public"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Heath Check OK 1")
	})
	public.TableInit()

	r.POST("/login", handlers.Login)
	v1 := r.Group("/api/v1")
	// v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/stock", public.StockC.List)
	}

}
