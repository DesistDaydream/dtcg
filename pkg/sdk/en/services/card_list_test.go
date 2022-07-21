package services

import (
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/en/models"
	"github.com/sirupsen/logrus"
)

func TestGetCardList(t *testing.T) {
	type args struct {
		r *models.CardListReq
	}
	tests := []struct {
		name    string
		args    args
		want    *models.CardList
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				r: &models.CardListReq{
					CardSet:     "",
					GameTitleID: "2",
					Limit:       "2000",
					Offset:      "0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCardList(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			logrus.Infof("卡牌总数: %v", got.Success.Total)
		})
	}
}
