package services

import (
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
)

func TestGetCardsDesc(t *testing.T) {
	type args struct {
		r *models.FilterConditionReq
	}
	tests := []struct {
		name    string
		args    args
		want    *models.CardDesc
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				r: &models.FilterConditionReq{
					Page:             "1",
					Limit:            "10",
					Name:             "",
					State:            "1",
					CardGroup:        "",
					RareDegree:       "",
					BelongsType:      "",
					CardLevel:        "",
					Form:             "",
					Attribute:        "",
					Type:             "",
					Color:            "",
					EnvolutionEffect: "",
					SafeEffect:       "",
					ParallCard:       "",
					KeyEffect:        "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCardDescs(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardsDesc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// fmt.Println(got)

			for _, c := range got.Page.List {
				logrus.WithFields(logrus.Fields{
					"name":       c.Name,
					"state":      c.State,
					"cardGroup":  c.CardGroup,
					"model":      c.Model,
					"rareDegree": c.RareDegree,
				}).Infoln("卡片详情")
			}

		})
	}
}
