package database

import "github.com/sirupsen/logrus"

type CardsDesc struct {
	Count int64      `json:"count"`
	Data  []CardDesc `json:"data"`
}

type CardDesc struct {
	ID                   int    `gorm:"primaryKey" json:"my_id"` // ID
	OfficialID           int    `json:"id"`
	CardGroup            string `json:"cardGroup"`              // 卡包
	Model                string `json:"model"`                  // 编号
	RareDegree           string `json:"rare_degree"`            // 稀有度
	BelongsType          string `json:"belongs_type"`           // 类别
	CardLevel            string `json:"card_level"`             // 等级
	Color                string `json:"color"`                  // 颜色
	Form                 string `json:"form"`                   // 形态
	Attribute            string `json:"attribute"`              // 属性
	Name                 string `json:"name"`                   // 名称
	Dp                   string `json:"dp"`                     // DP
	Type                 string `json:"type"`                   // 类型
	EntryConsumeValue    string `json:"entry_consume_value"`    // 登场费用
	EnvolutionConsumeOne string `json:"envolution_consume_one"` // 进化费用1
	EnvolutionConsumeTwo string `json:"envolution_consume_two"` // 进化费用2
	GetWay               string `json:"get_way"`                // 获得途径
	Effect               string `json:"effect"`                 // 效果
	SafeEffect           string `json:"safe_effect"`            // 安防效果
	EnvolutionEffect     string `json:"envolution_effect"`      // 进化源效果
	ImageCover           string `json:"image_cover"`            // 图片。这是一个卡图的 URL
	State                string `json:"state"`                  // 状态。0：显示，1：不显示
	ParallCard           string `json:"parall_card"`            // 是否是平卡。1 是平卡，0 是异画
	KeyEffect            string `json:"key_effect"`             // 效果关键字
}

func AddCardDesc(cardDesc *CardDesc) {
	// db.Create(cardDesc)
	result := db.FirstOrCreate(cardDesc, cardDesc)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡片描述
func ListCardDesc() (*CardsDesc, error) {
	var cd []CardDesc
	result := db.Find(&cd)
	if result.Error != nil {
		return nil, result.Error
	}

	return &CardsDesc{
		Count: result.RowsAffected,
		Data:  cd,
	}, nil
}
