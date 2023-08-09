package models

type CardVersionsPriceHistoryReqBody struct {
	CardVersionID string `json:"card_version_id"`
	GameKey       string `json:"game_key"`
	GameSubKey    string `json:"game_sub_key"`
	// 感觉应该还有一个字段，用来表示要求返回什么时间间隔的数据，默认是返回 14 天的价格
}
