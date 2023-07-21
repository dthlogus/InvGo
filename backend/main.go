package main

import (
	"github.com/dthlogus/InvGo/configs"
	"github.com/dthlogus/InvGo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)
	routes.PerfilRoute(router)
	routes.StocksRoute(router)

	router.Run(":9081")
}
