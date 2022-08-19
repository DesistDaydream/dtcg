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
	ID                   int    `json:"id"`                   //
	CardGroup            string `json:"cardGroup"`            // 卡包
	Model                string `json:"model"`                // 编号
	RareDegree           string `json:"rareDegree"`           // 稀有度
	BelongsType          string `json:"belongsType"`          // 类别
	CardLevel            string `json:"cardLevel"`            // 等级
	Color                string `json:"color"`                // 颜色
	Form                 string `json:"form"`                 // 形态
	Attribute            string `json:"attribute"`            // 属性
	Name                 string `json:"name"`                 // 名称
	Dp                   string `json:"dp"`                   // DP
	Type                 string `json:"type"`                 // 类型
	EntryConsumeValue    string `json:"entryConsumeValue"`    // 登场费用
	EnvolutionConsumeOne string `json:"envolutionConsumeOne"` // 进化费用1
	EnvolutionConsumeTwo string `json:"envolutionConsumeTwo"` // 进化费用2
	GetWay               string `json:"getWay"`
	Effect               string `json:"effect"`
	SafeEffect           string `json:"safeEffect"`
	EnvolutionEffect     string `json:"envolutionEffect"`
	ImageCover           string `json:"imageCover"`
	State                string `json:"state"`      // 状态。0：显示，1：不显示
	ParallCard           string `json:"parallCard"` // 是否是平卡。1 是平卡，0 是异画
	KeyEffect            string `json:"keyEffect"`
	CreateTime           string `json:"createTime"`
	UpdateTime           string `json:"updateTime"`
}
