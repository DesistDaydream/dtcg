package models

type CardDesc struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Page Page   `json:"page"`
}
type Page struct {
	TotalCount int        `json:"totalCount"`
	PageSize   int        `json:"pageSize"`
	TotalPage  int        `json:"totalPage"`
	CurrPage   int        `json:"currPage"`
	List       []PageList `json:"list"`
}
type PageList struct {
	ID                   int    `json:"id"`
	CardGroup            string `json:"cardGroup"`
	Model                string `json:"model"`
	RareDegree           string `json:"rareDegree"`
	BelongsType          string `json:"belongsType"`
	CardLevel            string `json:"cardLevel"`
	Color                string `json:"color"`
	Form                 string `json:"form"`
	Attribute            string `json:"attribute"`
	Name                 string `json:"name"`
	Dp                   string `json:"dp"`
	Type                 string `json:"type"`
	EntryConsumeValue    string `json:"entryConsumeValue"`
	EnvolutionConsumeOne string `json:"envolutionConsumeOne"`
	EnvolutionConsumeTwo string `json:"envolutionConsumeTwo"`
	GetWay               string `json:"getWay"`
	Effect               string `json:"effect"`
	SafeEffect           string `json:"safeEffect"`
	EnvolutionEffect     string `json:"envolutionEffect"`
	ImageCover           string `json:"imageCover"`
	State                string `json:"state"`
	ParallCard           string `json:"parallCard"`
	KeyEffect            string `json:"keyEffect"`
	CreateTime           string `json:"createTime"`
	UpdateTime           string `json:"updateTime"`
}
