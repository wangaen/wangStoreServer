package service

import (
	"errors"
	"fmt"
	"wangStoreServer/app/admin/models"
	"wangStoreServer/config"
)

// GetAdminUserInfoSer 获取admin用户信息
func GetAdminUserInfoSer(admin *models.AdminUser, userId int, username string) bool {
	if userId != 0 {
		admin.UserId = userId
	}
	if username != "" {
		admin.Username = username
	}
	result := config.DB.Debug().Limit(1).Where(admin).Find(&admin)
	if result.RowsAffected == 1 {
		return true
	}
	return false
}

// CreateAdminUserSer 创建amin用户
func CreateAdminUserSer(admin *models.AdminUser) error {
	result := config.DB.Debug().Create(admin)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

// UpdateAdminUserSer 修改amin用户
func UpdateAdminUserSer(admin *models.AdminUser) error {
	result := config.DB.Debug().Model(&admin).Select("*").Omit(
		"UserId",
		"Username",
		"Password",
		"CreateBy",
		"UpdateBy",
	).Updates(admin)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

// DeleteAdminUserSer 删除amin用户
func DeleteAdminUserSer(admin *models.AdminUser, userId int) bool {
	result := config.DB.Debug().Delete(admin, userId)
	fmt.Printf("%#v", result)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// GetAdminUserListSer 获取amin用户列表
func GetAdminUserListSer(admins *[]models.AdminUser, pageSize int64, offset int64, username string, email string, phone string, status string) (int64, error) {
	var count int64
	result := config.DB.Debug().Limit(int(pageSize)).Offset(int(offset)).
		Where(
			"username LIKE ? AND email LIKE ? AND phone LIKE ? AND status LIKE ? ",
			"%"+username+"%",
			"%"+email+"%",
			"%"+phone+"%",
			"%"+status+"%",
		).Find(admins).Count(&count)
	if result.Error != nil {
		return 0, errors.New(result.Error.Error())
	}
	return count, nil
}
