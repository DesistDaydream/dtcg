package services

import (
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
)

func TestGetCardPackage(t *testing.T) {
	tests := []struct {
		name    string
		want    *models.CardPackage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCardPackage()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, p := range got.List {
				logrus.WithFields(logrus.Fields{
					"名称": p.Name,
					"ID": p.ID,
					"状态": p.State,
				}).Info("卡包信息")
			}
		})
	}
}
