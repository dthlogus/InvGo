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

	router.Run(":9081")
}
