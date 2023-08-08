package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"reflect"

	"github.com/sirupsen/logrus"
)

// 实现类似 Gin 的 Bind 效果，将 Request 中的 Query 从结构体转为 map
// 以便在生成发起请求时，使用 req.URL.Query().Add() 注意为请求中的 Request 添加 Query
// 这个功能好像只有在自己暴露 API，并且传入的参数需要当做发起其他请求的 Query 时才有用
func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tField := t.Field(i)
		vField := v.Field(i)
		// 注意！！！注意！！！注意！！！
		// 传入的结构体中，要带有 form Tag 的才可以被解析为 map
		// 使用 form 这个 Tag 的原因是 Gin 的转换 map 逻辑中，也是使用的 form 作为 Tag
		tFieldTag := string(tField.Tag.Get("form"))
		if len(tFieldTag) > 0 {
			data[tFieldTag] = vField.String()
		} else {
			data[tField.Name] = vField.String()
		}
	}

	return data
}

var JhsRsaPublicKey string = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBg
QCkLMhnY5tb9T0KMqq4It/yK7Mv
4jQt39RyrH9yqPcAg0lsFWKTXJdT0/c0P+yX
R1aF2xLOZhl3NA8eZWEF2YoCBJg6
h6QJ6dlMak8r2LDC89QJfq1ZlcA6qfiHzZk
fUbtGqXj3RbzfvKyGUdQHvXp9P/1C
ECZfetRusF4IncOklwIDAQAB
-----END PUBLIC KEY-----`

// 使用 RSA 公钥加密指定的字符串
func EncryptWithRsaPublicKey(needEncryptString string, publicKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	cipherByte, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(needEncryptString))
	if err != nil {
		return nil, err
	}

	return cipherByte, nil
}

type AecCrypto struct {
	block cipher.Block
}

func NewAesCrypto(key []byte) *AecCrypto {
	genKey := make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}

	block, err := aes.NewCipher(genKey)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	return &AecCrypto{
		block: block,
	}
}

// 加密
func (a *AecCrypto) AesEncryptECB(origData []byte) ([]byte, error) {
	length := (len(origData) + aes.BlockSize) / aes.BlockSize

	plain := make([]byte, length*aes.BlockSize)

	copy(plain, origData)

	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}

	encrypted := make([]byte, len(plain))

	// 分组分块加密
	for bs, be := 0, a.block.BlockSize(); bs <= len(origData); bs, be = bs+a.block.BlockSize(), be+a.block.BlockSize() {
		a.block.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

// 解密
func (a *AecCrypto) AesDecryptECB(encrypted []byte) ([]byte, error) {
	decrypted := make([]byte, len(encrypted))

	for bs, be := 0, a.block.BlockSize(); bs < len(encrypted); bs, be = bs+a.block.BlockSize(), be+a.block.BlockSize() {
		a.block.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}
