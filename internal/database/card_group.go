package database

type CardGroup struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	State      string `json:"state"`
	Position   string `json:"position"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func AddCardGroup(cardGroup *CardGroup) {
	db.Create(cardGroup)
}

// 获取所有卡包
func ListCardGroups() ([]CardGroup, error) {
	var cg []CardGroup
	result := db.Find(&cg)
	if result.Error != nil {
		return nil, result.Error
	}
	return cg, nil
}
