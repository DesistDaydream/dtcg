package models

type SeriesGetResp struct {
	Data    []SeriesData `json:"data"`
	Message string       `json:"message"`
	Success bool         `json:"success"`
}

type SeriesData struct {
	SeriesID    int64        `json:"series_id"`
	SeriesIntro string       `json:"series_intro"`
	SeriesName  string       `json:"series_name"`
	SeriesOrder int64        `json:"series_order"`
	SeriesPack  []SeriesPack `json:"series_pack"`
}

type SeriesPack struct {
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
