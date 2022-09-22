package products

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
)

var token string = ""
var cardVersionID string = "2544"

func TestProductsClientGet(t *testing.T) {
	client := NewProductsClient(core.NewClient(token))

	got, err := client.Get(cardVersionID)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(got)
}

func TestProductsClientList(t *testing.T) {
	currentPage := 1
	for {
		client := NewProductsClient(core.NewClient(token))
		got, err := client.List(fmt.Sprint(currentPage))
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println(got.CurrentPage)

		if got.CurrentPage == got.LastPage {
			break
		}

		currentPage++
	}
}

func TestStructToMapStr(t *testing.T) {
	obj := models.ProductsGetRequestQuery{
		GameKey:       "dgm",
		SellerUserID:  "609077",
		CardVersionID: cardVersionID,
		Token:         token,
	}

	got := StructToMapStr(&obj)

	fmt.Println(len(got))

	gotByte, _ := json.Marshal(got)
	fmt.Println(string(gotByte))
	for k, v := range got {
		fmt.Println(k, v)
	}

}
