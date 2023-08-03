package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
)

// 测试将结构体转为 map
func TestStructToMapStr(t *testing.T) {
	obj := models.ProductsGetReqQuery{
		GameKey:       "dgm",
		SellerUserID:  "123456",
		CardVersionID: "98765432",
	}

	got := StructToMapStr(&obj)

	fmt.Println(len(got))

	gotByte, _ := json.Marshal(got)
	fmt.Println(string(gotByte))
	for k, v := range got {
		fmt.Println(k, v)
	}
}
