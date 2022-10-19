package database

import "github.com/sirupsen/logrus"

type CardGroupsOfficial struct {
	Count int64                   `json:"count"`
	Data  []CardGroupFromOfficial `json:"data"`
}

type CardGroupFromOfficial struct {
	ID         int    `gorm:"primaryKey"`
	OfficialID int    `json:"official_id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	State      string `json:"state"`
	Position   int    `json:"position"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func AddCardGroupFromOfficial(cardGroup *CardGroupFromOfficial) {
	// 根据第二个参数匹配记录，若没找到，则插入
	result := db.FirstOrCreate(cardGroup, cardGroup)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡包
func ListCardGroupsFromOfficial() (*CardGroupsOfficial, error) {
	var cg []CardGroupFromOfficial
	result := db.Find(&cg)
	if result.Error != nil {
		return nil, result.Error
	}

	return &CardGroupsOfficial{
		Count: result.RowsAffected,
		Data:  cg,
	}, nil
}
