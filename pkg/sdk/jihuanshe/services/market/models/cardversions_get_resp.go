package models

// 获取卡牌信息
type CardVersionGetResp struct {
	AvgPrice                  string        `json:"avg_price"`
	CardID                    int64         `json:"card_id"`
	CardNames                 []CardName    `json:"card_names"`
	CardVersions              []CardVersion `json:"card_versions"`
	DesireCount               int64         `json:"desire_count"`
	GameKey                   string        `json:"game_key"`
	GameSubKey                string        `json:"game_sub_key"`
	Grade                     interface{}   `json:"grade"`
	GradeDetail               []string      `json:"grade_detail"`
	GradeTotal                int64         `json:"grade_total"`
	GradeUserTotal            int64         `json:"grade_user_total"`
	ID                        int64         `json:"id"`
	ImageURL                  string        `json:"image_url"`
	IsDesire                  bool          `json:"is_desire"`
	IsGrade                   bool          `json:"is_grade"`
	IsLike                    bool          `json:"is_like"`
	LanguageText              string        `json:"language_text"`
	LikeCount                 int64         `json:"like_count"`
	MinPrice                  string        `json:"min_price"`
	NameCN                    string        `json:"name_cn"`
	NameOrigin                string        `json:"name_origin"`
	Number                    string        `json:"number"`
	Pack                      Pack          `json:"pack"`
	Ranking                   interface{}   `json:"ranking"`
	Rarity                    string        `json:"rarity"`
	UserCardVersionImageCount int64         `json:"user_card_version_image_count"`
}

type CardVersion struct {
	AvgPrice   string      `json:"avg_price"`
	CardID     int64       `json:"card_id"`
	CardNames  []CardName  `json:"card_names"`
	Grade      interface{} `json:"grade"`
	ID         int64       `json:"id"`
	ImageURL   string      `json:"image_url"`
	MinPrice   string      `json:"min_price"`
	NameCN     string      `json:"name_cn"`
	NameOrigin string      `json:"name_origin"`
	Number     string      `json:"number"`
	Pack       Pack        `json:"pack"`
	Rarity     string      `json:"rarity"`
}

type Pack struct {
	NameCN     string `json:"name_cn"`
	NameOrigin string `json:"name_origin"`
	PackID     int64  `json:"pack_id"`
	ReleasedAt string `json:"released_at"`
}
