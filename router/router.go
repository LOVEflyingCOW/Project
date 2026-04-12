package router

import (
	"exchangeapp/controllers"

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
	return r
}
