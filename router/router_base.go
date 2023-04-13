package router

import (
	"github.com/gin-gonic/gin"
	"im_app/middleware"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CorsMiddleware()).Use(middleware.GinLogger()).Use(middleware.GinRecovery(true))
	v1 := r.Group("/api/v1")
	{

		PublicGroup := v1.Group("")
		{
			InitWebSocketRouter(PublicGroup)

		}
	}
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")

	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "not found",
			"status": "404",
		})

	})
	return r

}
