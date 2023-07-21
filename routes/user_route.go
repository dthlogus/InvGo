package routes

import (
	"github.com/dthlogus/InvGo/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetUser())
	router.PUT("/user/:userId", controllers.UpdateUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.GET("/user", controllers.GetAllUser())
}
