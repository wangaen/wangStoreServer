package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"wangStoreServer/common/models"
	"wangStoreServer/config"
)

type AdminUser struct {
	UserId   int    `json:"userId" gorm:"primaryKey;autoIncrement;comment:管理员id"`
	Username string `json:"username" gorm:"size:64;unique;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId   int    `json:"roleId" gorm:"size:20;comment:角色ID"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex      int    `json:"sex" gorm:"size:2;comment:性别"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	DeptId   int    `json:"deptId" gorm:"size:20;comment:部门"`
	PostId   int    `json:"postId" gorm:"size:20;comment:岗位"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
	models.ControlBy
	models.ModelTime
}

func init() {
	adminUser := &AdminUser{}
	// 检测表名是否存在
	if !config.DB.Migrator().HasTable(adminUser.TableName()) {
		config.DB.AutoMigrate(&adminUser)
	}
}

func (AdminUser) TableName() string {
	return "admin_user"
}

func (a *AdminUser) GetId() interface{} {
	return a.UserId
}

// Encrypt 密码加密
func (a *AdminUser) Encrypt() (err error) {
	if a.Password == "" {
		return
	}
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		a.Password = string(hash)
		return
	}
}

// Compare 密码对比
func (a *AdminUser) Compare(enPwd string, newPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(enPwd), []byte(newPwd))
	if err != nil {
		return false
	} else {
		return true
	}
}

// BeforeCreate 钩子函数 - 创建前执行
func (a *AdminUser) BeforeCreate(_ *gorm.DB) error {
	return a.Encrypt()
}

// BeforeUpdate 钩子函数 - 更新前执行
func (a *AdminUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if a.Password != "" {
		err = a.Encrypt()
	}
	return err
}
