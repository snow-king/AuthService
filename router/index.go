package router

import (
	"AuthService/app/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup) {
	router.POST("/sign-up", controllers.SignUp)
	router.POST("/sign-in", controllers.SignIn)
	router.GET("/O2Auth/:service/index", controllers.Index)
	router.GET("/O2Auth/:service/callback", controllers.Callback)
	router.POST("/user/appendRole", controllers.AppendRole)
	router.POST("/user/appendNetwork", controllers.AppendSocNetworks)
	router.GET("/jwk", controllers.JWK)
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"response": "It's Alive! Alive!!!!"})
	})
}
