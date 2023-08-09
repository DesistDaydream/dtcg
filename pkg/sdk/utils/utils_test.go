package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
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

// var key string = "QCBY{Ru4~Y7}c,7H"
var key string = "1234567890123456"

func TestGenKeyAndData(t *testing.T) {
	reqBody := `{"game_key":"dgm","game_sub_key":"sc"}`
	// reqBody := `{"game_key":"dgm","game_sub_key":"sc","card_version_id":"2688"}`
	// reqBody := `{"categoryId":"4793","rarity":"","sorting":"number","sorting_price_type":"product","game_key":"dgm","game_sub_key":"sc","page":"1"}`

	a := NewAesCrypto([]byte(key))

	// 生成 key
	encryptedKey, _ := EncryptWithRsaPublicKey(key, JhsRsaPublicKey)
	fmt.Println(base64.StdEncoding.EncodeToString(encryptedKey))

	// 生成 data
	encryptedData, _ := a.AesEncryptECB([]byte(reqBody))

	fmt.Println(url.QueryEscape(base64.StdEncoding.EncodeToString(encryptedData)))
}

func TestDecryptData(t *testing.T) {
	a := NewAesCrypto([]byte(key))

	// 解密返回体
	dataByte, _ := base64.StdEncoding.DecodeString(Data)
	decryptedData, _ := a.AesDecryptECB(dataByte)
	fmt.Println(string(decryptedData))

	// 将解密后中的 Unicode 解码
	// 在单元测试里，没有声明结构体并使用 json 库解码到结构体中，导致响应字符串中有很多 Unicode
	newData, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(decryptedData)), `\\u`, `\u`, -1))
	fmt.Println(string(newData))
}

var Data string = ""
