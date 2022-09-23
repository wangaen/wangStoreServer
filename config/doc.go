package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"wangStoreServer/docs"
)

func InitDocs(engine *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("API文档地址： ", "http://localhost:8080/swagger/index.html")
}
