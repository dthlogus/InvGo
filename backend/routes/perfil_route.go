package routes

import (
	"github.com/dthlogus/InvGo/controllers"
	"github.com/gin-gonic/gin"
)

func PerfilRoute(router *gin.Engine) {
	router.POST("/perfil/:userId", controllers.CreatePerfil())
	router.PUT("/perfil/:userId", controllers.UpdatePerfil())
	router.DELETE("/perfil/:userId", controllers.DeletePerfil())
}
