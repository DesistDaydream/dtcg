package models

// 卡包列表
type CardPackage struct {
	Msg  string            `json:"msg"`
	Code int               `json:"code"`
	List []CardPackageList `json:"list"`
}
type CardPackageList struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	State      string `json:"state"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
