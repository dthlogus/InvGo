package main

import (
	"github.com/dthlogus/InvGo/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	configs.ConnectDB()

	router.Run(":9081")
}
