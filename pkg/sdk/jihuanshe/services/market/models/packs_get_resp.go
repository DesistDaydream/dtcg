package models

// 卡包信息
type PacksGetResp struct {
	CardVersionCount int64             `json:"card_version_count"`
	CardVersions     []PackCardVersion `json:"card_versions"`
	DesireCount      int64             `json:"desire_count"`
	GameKey          string            `json:"game_key"`
	GameSubKey       string            `json:"game_sub_key"`
	Grade            string            `json:"grade"`
	GradeDetail      []GradeDetail     `json:"grade_detail"`
	GradeTotal       int64             `json:"grade_total"`
	GradeUserTotal   int64             `json:"grade_user_total"`
	ID               int64             `json:"id"`
	ImageURL         string            `json:"image_url"`
	IsDesire         bool              `json:"is_desire"`
	IsGrade          bool              `json:"is_grade"`
	IsLike           bool              `json:"is_like"`
	LanguageText     string            `json:"language_text"`
	LikeCount        int64             `json:"like_count"`
	NameCN           string            `json:"name_cn"`
	NameOrigin       interface{}       `json:"name_origin"`
	Number           string            `json:"number"`
	Ranking          Ranking           `json:"ranking"`
	ReleasedAt       string            `json:"released_at"`
	Specs            interface{}       `json:"specs"`
}

type PackCardVersion struct {
	CardID     int        `json:"card_id"`
	CardNames  []CardName `json:"card_names"`
	GameKey    string     `json:"game_key"`
	GameSubKey string     `json:"game_sub_key"`
	Grade      *string    `json:"grade"`
	ID         int        `json:"id"`
	ImageURL   string     `json:"image_url"`
	MinPrice   string     `json:"min_price"`
	NameCN     string     `json:"name_cn"`
	NameOrigin string     `json:"name_origin"`
	Number     string     `json:"number"`
	Rarity     string     `json:"rarity"`
}

type GradeDetail struct {
	Grade     int64 `json:"grade"`
	UserCount int64 `json:"user_count"`
}

type Ranking struct {
	AllGame      int64  `json:"all_game"`
	Rank         int64  `json:"rank"`
	RankTypeID   int64  `json:"rank_type_id"`
	RankTypeName string `json:"rank_type_name"`
}
