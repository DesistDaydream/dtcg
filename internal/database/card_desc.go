package database

type CardDesc struct {
	// gorm.Model 是一个包含了ID, CreatedAt, UpdatedAt, DeletedAt四个字段的结构体。
	// gorm.Model
	ID                   int    `gorm:"primaryKey" json:"id"` // ID
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
	GetWay               string `json:"getWay"`               //
	Effect               string `json:"effect"`               // 效果
	SafeEffect           string `json:"safeEffect"`           // 安防效果
	EnvolutionEffect     string `json:"envolutionEffect"`     // 进化源效果
	ImageCover           string `json:"imageCover"`           // 图片。这是一个卡图的 URL
	State                string `json:"state"`                // 状态。0：显示，1：不显示
	ParallCard           string `json:"parallCard"`           // 是否是平卡。1 是平卡，0 是异画
	KeyEffect            string `json:"keyEffect" copier:"-"` // 效果关键字
}

func AddCard(cardDesc *CardDesc) {
	db.Create(cardDesc)
}
