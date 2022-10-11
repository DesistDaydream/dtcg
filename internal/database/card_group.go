package database

import "github.com/sirupsen/logrus"

type CardGroups struct {
	Count int64       `json:"count"`
	Data  []CardGroup `json:"data"`
}

type CardGroup struct {
	ID         int    `gorm:"primaryKey"`
	OfficialID int    `json:"official_id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	State      string `json:"state"`
	Position   string `json:"position"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
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
		Count: result.RowsAffected,
		Data:  cg,
	}, nil
}
