package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

var publicKey string = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBg
QCkLMhnY5tb9T0KMqq4It/yK7Mv
4jQt39RyrH9yqPcAg0lsFWKTXJdT0/c0P+yX
R1aF2xLOZhl3NA8eZWEF2YoCBJg6
h6QJ6dlMak8r2LDC89QJfq1ZlcA6qfiHzZk
fUbtGqXj3RbzfvKyGUdQHvXp9P/1C
ECZfetRusF4IncOklwIDAQAB
-----END PUBLIC KEY-----`

// var key string = "QCBY{Ru4~Y7}c,7H"
var key string = "1234567890123456"

func TestGenKeyAndData(t *testing.T) {
	// msg := `{"game_key":"dgm","game_sub_key":"sc"}`
	msg := `{"game_key":"dgm","game_sub_key":"sc""card_version_id":"2688"}`

	a := NewAesCrypto([]byte(key))

	// 生成 key
	encryptedKey := encryptWithRsaPublicKey(key, publicKey)
	fmt.Println(encryptedKey)

	// 生成 data
	encrypted := a.AesEncryptECB([]byte(msg))
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
}

func TestDecryptData(t *testing.T) {
	a := NewAesCrypto([]byte(key))

	// 解密返回体
	dataByte, _ := base64.StdEncoding.DecodeString(Data)
	decrypted := a.AesDecryptECB(dataByte)

	// 将解密后中的 Unicode 解码
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(decrypted)), `\\u`, `\u`, -1))
	fmt.Println(string(str))
}
