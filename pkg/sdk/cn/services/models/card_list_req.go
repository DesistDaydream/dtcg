package models

type FilterConditionReq struct {
	// 页数
	Page string `json:"page"`
	// 每页显示的卡片数量
	Limit string `json:"limit"`
	// 卡名
	Name string `json:"name"`
	// 卡片状态。0为显示，1为隐藏
	State string `json:"state"`
	// 所属卡盒
	CardGroup string `json:"cardGroup"`
	// 稀有度
	RareDegree string `json:"rareDegree"`
	// 类别。包括 数码蛋、数码宝贝、驯兽师、选项
	BelongsType string `json:"belongsType"`
	// 等级。包括 Lv.2、Lv.3...Lv.7
	CardLevel string `json:"cardLevel"`
	// 形态。幼年期、成长期、成熟期、完全体、究极体
	Form string `json:"form"`
	// 属性。病毒种、数据种、疫苗种、自由、不明
	Attribute string `json:"attribute"`
	// 类型。龙型、天使型、等等...... 等等
	Type string `json:"type"`
	// 颜色。红、蓝、黄、绿、白、黑、紫、混合色
	Color string `json:"color"`
	// 是否有进化源效果.0有,1没有
	EnvolutionEffect string `json:"envolutionEffect"`
	// 是否具有安防效果。0有，1没有
	SafeEffect string `json:"safeEffect"`
	// 是否是异画。0是，1不是。即
	ParallCard string `json:"parallCard"`
	// 关键词效果。抽1张卡、干扰、贯通...... 等等
	KeyEffect string `json:"keyEffect"`
}
