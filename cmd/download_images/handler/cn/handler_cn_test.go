package cn

import (
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"
)

func TestImageHandlerGetImagesURL(t *testing.T) {
	type fields struct {
		Lang string
	}
	type args struct {
		c *models.FilterConditionReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test",
			fields:  fields{Lang: ""},
			args:    args{},
			want:    []string{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt.args.c = &models.FilterConditionReq{
			Page:             "",
			Limit:            "10",
			Name:             "",
			State:            "0",
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
		}

		t.Run(tt.name, func(t *testing.T) {
			i := &ImageHandler{
				Lang: tt.fields.Lang,
			}
			urls, err := i.GetImagesURL(tt.args.c)
			if err != nil {
				panic(err)
			}

			for _, url := range urls {
				logrus.Infoln(url)
			}

		})
	}
}

func TestImageHandlerGenFileName(t *testing.T) {
	type fields struct {
		Lang string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "test",
			fields: fields{Lang: "cn"},
			args:   args{url: "https://digimoncard-1258002530.file.myqcloud.com/DTCG/BTC2_BT3-034_D%C2%A0%E6%BA%90.png"},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ImageHandler{
				Lang: tt.fields.Lang,
			}
			fileName := i.GenFileName(tt.args.url)

			logrus.Infoln(fileName)
		})
	}
}
