package models

// 卡包列表
type CardGroupsResponse struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	List []CardGroup `json:"list"`
}
type CardGroup struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	State      string `json:"state"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
