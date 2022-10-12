package database

import "github.com/sirupsen/logrus"

type CardGroupsOfficial struct {
	Count int64               `json:"count"`
	Data  []CardGroupOfficial `json:"data"`
}

type CardGroupOfficial struct {
	ID         int    `gorm:"primaryKey"`
	OfficialID int    `json:"official_id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	State      string `json:"state"`
	Position   string `json:"position"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func AddCardGroupFromOfficial(cardGroup *CardGroupOfficial) {
	// 根据第二个参数匹配记录，若没找到，则插入
	result := db.FirstOrCreate(cardGroup, cardGroup)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡包
func ListCardGroupsFromOfficial() (*CardGroupsOfficial, error) {
	var cg []CardGroupOfficial
	result := db.Find(&cg)
	if result.Error != nil {
		return nil, result.Error
	}

	return &CardGroupsOfficial{
		Count: result.RowsAffected,
		Data:  cg,
	}, nil
}
