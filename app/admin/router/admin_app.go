package router

import (
	"github.com/gin-gonic/gin"
	"wangStoreServer/app/admin/controller"
)

func appSwiperImg(r *gin.RouterGroup) {
	app := r.Group("/app/swiper")
	{
		app.POST("/upload", controller.UploadSwiperImg)
		app.DELETE("/delete", controller.DeleteSwiperImg)
	}
}
