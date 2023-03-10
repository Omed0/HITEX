package routes

import (
	"emad.com/auth"
	"emad.com/controllers"
	"emad.com/middleware"
	"github.com/gin-gonic/gin"
)

var CheckToken = middleware.CheckToken

func SetUpRowters (router *gin.Engine)  {

	router.POST("/api/login", auth.Login)


	router.POST("/api/user/create", controllers.Create)
	router.GET("/api/user/list", CheckToken, controllers.List)
	router.GET("/api/user/one/:id", CheckToken, controllers.One)
	router.PUT("/api/user/update/:id", CheckToken, controllers.Update)
	router.DELETE("/api/user/delete/:id", controllers.Delete)




	
}


