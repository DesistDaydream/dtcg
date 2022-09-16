package models

type CacheListResp struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	List []CacheList `json:"list"`
}
type CacheList struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
