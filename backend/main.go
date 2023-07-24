package main

import (
	"github.com/dthlogus/InvGo/backend/configs"
	"github.com/dthlogus/InvGo/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)
	routes.PerfilRoute(router)
	routes.StocksRoute(router)
	routes.AuthRoute(router)

	router.Run(":9081")
}
