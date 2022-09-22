package products

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
)

var token string = ""
var cardVersionID string = "2544"

func getToken() {
	file, err := os.ReadFile("token.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	token = string(file)
}

func TestProductsClientGet(t *testing.T) {
	getToken()
	client := NewProductsClient(core.NewClient(token))

	got, err := client.Get(cardVersionID)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(got)
}

func TestProductsClientList(t *testing.T) {
	getToken()
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

type NeedAddCards struct {
	NeedAddCards []NeedAddCard
}

type NeedAddCard struct {
	CardVersionID string
	Price         string
	Quantity      string
}

func TestProductsClientAdd(t *testing.T) {
	getToken()
	client := NewProductsClient(core.NewClient(token))

	resp, err := client.Add("1871", "199", "4")
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(resp)
}
