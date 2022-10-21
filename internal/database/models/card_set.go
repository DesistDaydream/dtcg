package models

type CardSets struct {
	Count int64     `json:"count"`
	Data  []CardSet `json:"data"`
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
