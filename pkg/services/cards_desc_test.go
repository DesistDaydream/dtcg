package services

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/models"
)

func TestGetCardsDesc(t *testing.T) {
	type args struct {
		c *FilterCondition
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
				c: &FilterCondition{
					Page:             "1",
					Limit:            "40",
					Name:             "",
					State:            "0",
					CardGroup:        "BTC-01",
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
			got, err := GetCardDescs(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardsDesc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, c := range got.Page.List {
				fmt.Println(c.Name)
			}

		})
	}
}
