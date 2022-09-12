package main

import (
	"github.com/gin-gonic/gin"
	"wangStoreServer/common/router"
)

func main() {

	engine := gin.Default()
	router.InitGinRouter(engine)
	engine.Run(":8080")

}
