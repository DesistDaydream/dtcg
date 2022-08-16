package handler

import (
	"fmt"
	"testing"
)

func TestGenerateDir(t *testing.T) {
	type args struct {
		dirPrefix       string
		lang            string
		cardPackageName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				dirPrefix:       "/mnt/d/Projects/images",
				lang:            "zh",
				cardPackageName: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := GenerateDir(tt.args.dirPrefix, tt.args.lang, tt.args.cardPackageName)
			fmt.Println(path)
		})
	}
}
