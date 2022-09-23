package main

import (
	"github.com/gin-gonic/gin"
	"wangStoreServer/common/router"
	"wangStoreServer/config"
)

func main() {
	engine := gin.Default()
	router.InitGinRouter(engine)
	config.InitDocs(engine)
	engine.Run(":8080")
}
