package router

import (
	"github.com/gin-gonic/gin"
	"wangStoreServer/app/admin/controller"
)

func init() {
	routerNoCheckList = append(routerNoCheckList, login, register)
}

func login(r *gin.RouterGroup) {
	r.POST("/login", controller.AdminInfoController{}.AdminLogin)
}

func register(r *gin.RouterGroup) {
	r.POST("/register", controller.AdminInfoController{}.AdminRegister)
}
