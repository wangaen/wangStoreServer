package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"wangStoreServer/app/admin/form"
	"wangStoreServer/app/admin/models"
	"wangStoreServer/app/admin/service"
	"wangStoreServer/common/controller"
)

type AdminInfoController struct {
	models.AdminUser
}

// AdminLogin 登录
func (ac AdminInfoController) AdminLogin(c *gin.Context) {
	var form = form.LoginForm{}
	// 校验参数
	if !controller.ValidParams(c, &form) {
		return
	}
	// 判断用户是否存在
	if err := c.ShouldBindBodyWith(&ac, binding.JSON); err != nil {
		controller.ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	if !service.GetAdminUserInfoSer(&ac.AdminUser, 0, ac.Username) {
		controller.OkResponseBody(c, "fail", "用户名不存在", "")
		return
	}
	// 判断密码
	if !ac.Compare(ac.Password, form.Password) {
		controller.OkResponseBody(c, "fail", "密码不正确", "")
		return
	}
	// 生成token
	token := controller.Auth{}.SetAuthToken(c, ac.Username, ac.Password)
	if token == "" {
		return
	}
	// 成功
	controller.OkResponseBody(c, "ok", "登录成功", gin.H{"userInfo": ac, "token": token})
}

// AdminRegister 注册
func (ac AdminInfoController) AdminRegister(c *gin.Context) {
	var form = form.LoginForm{}
	// 校验参数
	if !controller.ValidParams(c, &form) {
		return
	}
	// 判断用户是否存在
	if err := c.ShouldBindBodyWith(&ac, binding.JSON); err != nil {
		controller.ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	if service.GetAdminUserInfoSer(&ac.AdminUser, 0, ac.Username) {
		controller.OkResponseBody(c, "fail", "用户名已存在", "")
		return
	}
	// 创建用户
	ac.Password = form.Password
	if err := service.CreateAdminUserSer(&ac.AdminUser); err != nil {
		controller.ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	// 成功
	controller.OkResponseBody(c, "ok", "注册成功", gin.H{"userInfo": ac})
}

// GetAdminUserInfo 获取admin用户信息
func (ac AdminInfoController) GetAdminUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	id, ok := strconv.ParseInt(userId, 10, 64)
	if ok != nil {
		controller.ErrResponseBody(c, http.StatusBadRequest, "fail", "参数异常")
		return
	}
	if !service.GetAdminUserInfoSer(&ac.AdminUser, int(id), "") {
		controller.OkResponseBody(c, "ok", "该用户信息不存在", "")
		return
	}
	controller.OkResponseBody(c, "ok", "查询成功", gin.H{"userInfo": ac})
}

// UpdateAdminUserInfo 修改admin用户信息
func (ac AdminInfoController) UpdateAdminUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	id, ok := strconv.ParseInt(userId, 10, 64)
	if ok != nil {
		controller.ErrResponseBody(c, http.StatusBadRequest, "fail", "参数异常")
		return
	}
	if err := c.ShouldBindJSON(&ac); err != nil {
		controller.ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	ac.AdminUser.UserId = int(id)
	if err := service.UpdateAdminUserSer(&ac.AdminUser); err != nil {
		controller.ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	controller.OkResponseBody(c, "ok", "修改成功", gin.H{"userInfo": ac})
}

// GetAdminUserList 获取admin用户列表
func (ac AdminInfoController) GetAdminUserList(c *gin.Context) {
	limit, _ := strconv.ParseInt(c.DefaultQuery("pageIndex", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("pageSize", "10"), 10, 64)
	username := c.Query("username")
	phone := c.Query("phone")
	email := c.Query("email")
	status := c.Query("status")
	offset := limit*pageSize - pageSize
	var adminUsers []models.AdminUser
	count, err := service.GetAdminUserListSer(&adminUsers, pageSize, offset, username, email, phone, status)
	if err != nil {
		controller.ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	controller.OkResponseBody(c, "ok", "查询成功", gin.H{
		"list":      adminUsers,
		"pageIndex": limit,
		"pageSize":  pageSize,
		"count":     count,
	})

}

// DeleteAdminUser 删除admin用户
func (ac AdminInfoController) DeleteAdminUser(c *gin.Context) {
	userId := c.Param("userId")
	id, ok := strconv.ParseInt(userId, 10, 64)
	if ok != nil {
		controller.ErrResponseBody(c, http.StatusBadRequest, "fail", "参数异常")
		return
	}
	if !service.DeleteAdminUserSer(&ac.AdminUser, int(id)) {
		controller.OkResponseBody(c, "fail", "删除失败", "")
		return
	}
	controller.OkResponseBody(c, "ok", "删除成功", "")
}
