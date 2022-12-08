package models

// 数据库模型。卡牌集合信息
type CardSets struct {
	Count       int64     `json:"count"`
	PageSize    int       `json:"page_size"`
	PageCurrent int       `json:"page_current"`
	PageTotal   int       `json:"page_total"`
	Data        []CardSet `json:"data"`
}

type CardSet struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	SeriesID       int    `json:"series_id"`
	SeriesName     string `json:"series_name"`
	Language       string `json:"language"`
	SetCover       string `json:"set_cover"`
	SetEnName      string `json:"set_enName"`
	SetID          int    `json:"set_id"`
	SetJapName     string `json:"set_japName"`
	SetName        string `json:"set_name"`
	SetPrefix      string `json:"set_prefix"`
	SetReleaseDate string `json:"set_releaseDate"`
	SetRemark      string `json:"set_remark"`
}
