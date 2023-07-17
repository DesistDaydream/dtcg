package models

type WithdrawResp struct {
	CurrentPage int            `json:"current_page"`
	Data        []WithdrawData `json:"data"`
	From        int64          `json:"from"`
	LastPage    int            `json:"last_page"`
	NextPageURL string         `json:"next_page_url"`
	Path        string         `json:"path"`
	PerPage     int64          `json:"per_page"`
	PrevPageURL string         `json:"prev_page_url"`
	To          int64          `json:"to"`
	Total       int64          `json:"total"`
}

type WithdrawData struct {
	CreatedAt       string `json:"created_at"`
	Ecommerce       int64  `json:"ecommerce"`
	EcommerceStatus string `json:"ecommerce_status"`
	Money           string `json:"money"`
	Remark          string `json:"remark"`
	Status          int64  `json:"status"`
	StatusText      string `json:"status_text"`
	WithdrawLogID   int64  `json:"withdraw_log_id"`
}
