package router

import (
	"github.com/gin-gonic/gin"
	"wangStoreServer/common/controller"
)

var (
	routerNoCheckList = make([]func(engine *gin.RouterGroup), 0)
	routerCheckList   = make([]func(engine *gin.RouterGroup, authMiddleware controller.Auth), 0)
)

func InitRouter(r *gin.Engine, authMiddleware controller.Auth) {
	noCheckRouter(r)
	checkRouter(r, authMiddleware)
}

func noCheckRouter(r *gin.Engine) {
	routerGroup := r.Group("/admin")
	for _, fun := range routerNoCheckList {
		fun(routerGroup)
	}
}

func checkRouter(r *gin.Engine, authMiddleware controller.Auth) {
	routerGroup := r.Group("/admin")
	for _, fun := range routerCheckList {
		fun(routerGroup, authMiddleware)
	}
}
