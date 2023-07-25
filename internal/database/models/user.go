package models

import "time"

// 数据库模型。用户信息
type Users struct {
	Count       int64  `json:"count"`
	PageSize    int    `json:"page_size"`
	PageCurrent int    `json:"page_current"`
	PageTotal   int    `json:"page_total"`
	Data        []User `json:"data"`
}

type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	MoecardToken string    `json:"moecard_token"`
	JhsToken     string    `json:"jhs_token"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"update_at"`
}
