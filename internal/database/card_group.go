package database

import "github.com/sirupsen/logrus"

type CardGroups struct {
	Data []CardGroup
}

type CardGroup struct {
	ID         int `gorm:"primaryKey"`
	OfficialID int
	Name       string
	Image      string
	State      string
	Position   string
	CreateTime string
	UpdateTime string
}

func AddCardGroup(cardGroup *CardGroup) {
	// 根据第二个参数匹配记录，若没找到，则插入
	result := db.FirstOrCreate(cardGroup, cardGroup)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡包
func ListCardGroups() (*CardGroups, error) {
	var cg []CardGroup
	result := db.Find(&cg)
	if result.Error != nil {
		return nil, result.Error
	}

	return &CardGroups{
		Data: cg,
	}, nil
}
