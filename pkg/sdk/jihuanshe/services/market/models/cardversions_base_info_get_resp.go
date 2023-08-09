package models

// 卡牌基础信息
type CardVersionsBaseInfoResp struct {
	CardID                int        `json:"card_id"`
	CardNames             []CardName `json:"card_names"`
	DescEvoCN             string     `json:"desc_evo_cn"`
	DescMainCN            string     `json:"desc_main_cn"`
	DescSecurityCN        string     `json:"desc_security_cn"`
	EffectByHTML          string     `json:"effect_by_html"`
	Grade                 string     `json:"grade"`
	ID                    int64      `json:"id"`
	ImageURL              string     `json:"image_url"`
	LanguageText          string     `json:"language_text"`
	NameCN                string     `json:"name_cn"`
	NameOrigin            string     `json:"name_origin"`
	Number                string     `json:"number"`
	OtherCardVersionCount int64      `json:"other_card_version_count"`
	PackID                int64      `json:"pack_id"`
	PackNameCN            string     `json:"pack_name_cn"`
	PackReleasedAt        string     `json:"pack_released_at"`
	Ranking               string     `json:"ranking"`
	Rarity                string     `json:"rarity"`
}
