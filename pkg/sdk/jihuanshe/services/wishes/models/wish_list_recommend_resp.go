package models

// 获取推荐列表的响应体
type WishListRecommendResp struct {
	CurrentPage int64                   `json:"current_page"`
	Data        []WishListRecommendData `json:"data"`
	From        int64                   `json:"from"`
	LastPage    int64                   `json:"last_page"`
	NextPageURL string                  `json:"next_page_url"`
	Path        string                  `json:"path"`
	PerPage     int64                   `json:"per_page"`
	PrevPageURL string                  `json:"prev_page_url"`
	To          int64                   `json:"to"`
	Total       int64                   `json:"total"`
}

type WishListRecommendData struct {
	CardVersionImage string `json:"card_version_image"`
	Desc             string `json:"desc"`
	Name             string `json:"name"`
	Quantity         int64  `json:"quantity"`
	WishListID       int64  `json:"wish_list_id"`
}
