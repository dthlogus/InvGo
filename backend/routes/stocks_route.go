package routes

import (
	"github.com/dthlogus/InvGo/backend/controllers"
	"github.com/gin-gonic/gin"
)

func StocksRoute(router *gin.Engine) {
	router.GET("/stocks/:symbols", controllers.GetStocks())
}
