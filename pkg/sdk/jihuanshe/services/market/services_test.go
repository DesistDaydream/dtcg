package market

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/sirupsen/logrus"
)

func TestMarketClient_GetProductSellers(t *testing.T) {
	client := NewMarketClient(core.NewClient(""))
	got, err := client.GetProductSellers("2676", "1")
	if err != nil {
		logrus.Errorln(err)
	}

	for _, data := range got.Data {
		fmt.Println(data.CardVersionImage)
	}
}
