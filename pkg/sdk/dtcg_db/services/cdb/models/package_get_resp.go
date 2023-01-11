package models

type PackageGetResp struct {
	Data    PackageGetData `json:"data"`
	Message string         `json:"message"`
	Success bool           `json:"success"`
}

type PackageGetData struct {
	Cards           []Card `json:"cards"`
	Language        string `json:"language"`
	PackCover       string `json:"pack_cover"`
	PackEnName      string `json:"pack_enName"`
	PackID          int64  `json:"pack_id"`
	PackJapName     string `json:"pack_japName"`
	PackName        string `json:"pack_name"`
	PackPrefix      string `json:"pack_prefix"`
	PackReleaseDate string `json:"pack_releaseDate"`
	PackRemark      string `json:"pack_remark"`
}
