package models

import (
	"fmt"
	"time"
)

// 数据库模型。用户信息
type Users struct {
	Count       int64  `json:"count"`
	PageSize    int    `json:"page_size"`
	PageCurrent int    `json:"page_current"`
	PageTotal   int    `json:"page_total"`
	Data        []User `json:"data"`
}

type User struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	MoecardUsername string    `json:"moecard_username"`
	MoecardPassword string    `json:"moecard_password"`
	MoecardToken    string    `json:"moecard_token"`
	JhsUsername     string    `json:"jhs_username"`
	JhsToken        string    `json:"jhs_token"`
	JhsRsaPublicKey string    `json:"rsa_public_key"`
	CreatedAt       time.Time `json:"create_at"`
	UpdatedAt       time.Time `json:"update_at"`
}

func (u *User) ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is empty")
	}
	if u.Password != password {
		return fmt.Errorf("password is wrong")
	}
	return nil
}
