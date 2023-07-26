package models

type LoginReqBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}
