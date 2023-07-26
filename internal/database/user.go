package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

// 获取用户信息
func GetUser(userID string) (*models.User, error) {
	var user models.User
	result := DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByName(username string) (*models.User, error) {
	var user models.User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// 更新用户信息
func UpdateUser(user *models.User, condition map[string]interface{}) error {
	result := DB.Model(user).Where("id = ?", user.ID).Updates(condition)

	if result.Error != nil {
		return result.Error
	}

	logrus.Debugf("已更新 %v 条数据", result.RowsAffected)

	return nil
}

// 创建管理员账户
func CraetAdminUser() error {
	var count int64

	if err := DB.Table("users").Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		user := models.User{
			ID:           1,
			Username:     "DesistDaydream",
			Password:     "DesistDaydream",
			MoecardToken: "",
			JhsToken:     "",
		}

		result := DB.FirstOrCreate(user, models.User{ID: user.ID})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
