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

// 更新用户信息
func UpdateUser(user *models.User, condition map[string]interface{}) {
	result := DB.Model(user).Where("id = ?", user.ID).Updates(condition)

	if result.Error != nil {
		logrus.Errorf("更新 %v 信息异常: %v", user.Username, result.Error)
	}

	logrus.Debugf("已更新 %v 条数据", result.RowsAffected)
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
			logrus.Errorf("创建管理员账户失败，原因: %v", result.Error)
		}
	}

	return nil
}
