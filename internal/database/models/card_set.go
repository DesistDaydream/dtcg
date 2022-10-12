package models

type CardSets struct {
	Count int64     `json:"count"`
	Data  []CardSet `json:"data"`
}

type CardSet struct {
	ID              int    `json:"id" gorm:"primaryKey"`
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
