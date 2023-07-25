package routes

import (
	"github.com/dthlogus/InvGo/backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/auth/user", controllers.AuthenticateUser())

}
