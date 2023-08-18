package card

import (
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/models"
	"github.com/sirupsen/logrus"
)

func TestGetCardFilterInfo(t *testing.T) {
	type args struct {
		r *models.CardFilterInfoReq
	}
	tests := []struct {
		name    string
		args    args
		want    *models.CardMetadataGetResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				r: &models.CardFilterInfoReq{
					GameTitleID:  "2",
					LanguageCode: "EN",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCardFilterInfo(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardFilterInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, cardSet := range got.Success.CardSetList {
				logrus.WithFields(logrus.Fields{
					"名称": cardSet.Name,
					"编号": cardSet.ID,
				}).Infof("卡包信息")
			}

			for _, config := range got.Success.CardSearchConfigs {
				logrus.WithFields(logrus.Fields{
					"名称":  config.ConfigName,
					"可用值": config.Choices,
				}).Infof("搜索配置")
			}
		})
	}
}
