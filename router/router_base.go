package router

import (
	"app_ws/global"
	"app_ws/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"net/http"

	gs "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CorsMiddleware()).Use(middleware.GinLogger()).Use(middleware.GinRecovery(true)).Use(middleware.Sentinel())
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	global.GVA_LOG.Info("register swagger handler")
	v1 := r.Group("/api/v1")
	{

		//无token 访问
		PublicGroup := v1.Group("")
		{
			InitWebSocketRouter(PublicGroup)

		}
		//token 访问
		//PrivateGroup := v1.Group("")
		//PrivateGroup.Use(middleware.JWTAuthMiddleware())
		//{
		//	//InitPostRouter(PrivateGroup)
		//	//InitCommentRouter(PrivateGroup)
		//
		//}
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
