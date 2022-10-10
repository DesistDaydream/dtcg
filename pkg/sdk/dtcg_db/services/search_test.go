package services

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/sirupsen/logrus"
)

func TestSearchClient_CardDeckSearch(t *testing.T) {
	client := NewSearchClient(core.NewClient(""))
	got, err := client.PostCardSearch(50)
	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println(got)
}
