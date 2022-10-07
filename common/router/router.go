package router

import (
	"github.com/gin-gonic/gin"
	admin "wangStoreServer/app/admin/router"
	crawler "wangStoreServer/app/crawler/router"
	"wangStoreServer/common/captcha"
	"wangStoreServer/common/controller"
)

func InitGinRouter(r *gin.Engine) {
	admin.InitRouter(r, controller.Auth{})
	captcha.InitCaptcha(r)
	crawler.InitCrawlerRouter(r)
}
