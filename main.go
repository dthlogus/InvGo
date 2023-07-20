package main

import (
	"github.com/dthlogus/InvGo/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.Run("locahost:9081")
}
