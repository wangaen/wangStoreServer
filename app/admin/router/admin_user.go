package router

import (
	"github.com/gin-gonic/gin"
	"wangStoreServer/app/admin/controller"
	common "wangStoreServer/common/controller"
)

func init() {
	routerCheckList = append(routerCheckList, AdminUser)
}

func AdminUser(r *gin.RouterGroup, middlewareAuth common.Auth) {
	u := r.Group("/user").Use(middlewareAuth.ValidAuthToken())
	{
		u.GET("/list", controller.AdminInfoController{}.GetAdminUserList)
		u.GET("/:userId", controller.AdminInfoController{}.GetAdminUserInfo)
		u.PUT("/:userId", controller.AdminInfoController{}.UpdateAdminUserInfo)
		u.DELETE("/:userId", controller.AdminInfoController{}.DeleteAdminUser)
	}
}
