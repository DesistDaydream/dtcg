package models

import "time"

type UsersGetResp struct {
	Count       int64      `json:"count"`
	PageSize    int        `json:"page_size"`
	PageCurrent int        `json:"page_current"`
	PageTotal   int        `json:"page_total"`
	Data        []UserData `json:"data"`
}

type UserData struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username"`
	MoecardToken string    `json:"moecard_token"`
	JhsToken     string    `json:"jhs_token"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"update_at"`
}
