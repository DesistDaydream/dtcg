package database

import "github.com/sirupsen/logrus"

type CardGroupsFromDtcgDB struct {
	Count int64                 `json:"count"`
	Data  []CardGroupFromDtcgDB `json:"data"`
}

type CardGroupFromDtcgDB struct {
	ID              int    `gorm:"primaryKey"`
	SeriesID        int    `json:"series_id"`
	SeriesName      string `json:"series_name"`
	Language        string `json:"language"`
	PackCover       string `json:"pack_cover"`
	PackEnName      string `json:"pack_enName"`
	PackID          int    `json:"pack_id"`
	PackJapName     string `json:"pack_japName"`
	PackName        string `json:"pack_name"`
	PackPrefix      string `json:"pack_prefix"`
	PackReleaseDate string `json:"pack_releaseDate"`
	PackRemark      string `json:"pack_remark"`
}

func AddCardGroupFromDtcgDB(cardGroup *CardGroupFromDtcgDB) {
	result := db.FirstOrCreate(cardGroup, cardGroup)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡包
func ListCardGroupsFromDtcgDB() (*CardGroupsFromDtcgDB, error) {
	var cg []CardGroupFromDtcgDB
	result := db.Find(&cg)
	if result.Error != nil {
		return nil, result.Error
	}

	return &CardGroupsFromDtcgDB{
		Count: result.RowsAffected,
		Data:  cg,
	}, nil
}
