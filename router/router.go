package router

import (
	"exchangeapp/controllers"
	"exchangeapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login) //func(ctx *gin.Context) {
		//ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		//	"msg": "Login Success",
		//	}

		auth.POST("/register", controllers.Register) // func(ctx *gin.Context) {
		//ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		//	"msg": "Register Success",
		//}

	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
	}
	return r
}
