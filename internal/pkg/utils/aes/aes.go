package aes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

var gIV []byte

func init() {
	gIV = []byte("123456789abcdefghijklmnpqrst")
}

type IAesCipher interface {
	Encrypt(origData []byte) []byte
	Decrypt(origData []byte) []byte
}

func NewAesCipher(key, iv []byte) IAesCipher {
	if len(key) > BLOCK_SIZE_16 {
		return NewAesCipher256(key, iv)
	}
	return NewAesCipher128(key, iv)
}

// 加密
func AesEncrypt(key, iv []byte, origData []byte) []byte {
	aesCipher := NewAesCipher(key, iv)
	return aesCipher.Encrypt(origData)
}

// 解密
func AesDecrypt(key, iv []byte, encrData []byte) []byte {
	aesCipher := NewAesCipher(key, iv)
	return aesCipher.Decrypt(encrData)
}

// 加密
func AesEncryptWithIV(key, origData string) string {
	encrypt := AesEncrypt([]byte(key), gIV, []byte(origData))
	encodeString := base64.StdEncoding.EncodeToString(encrypt)
	return encodeString
}

// 解密
func AesDecryptWithIV(key, encryptData string) string {
	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return ""
	}
	return string(AesDecrypt([]byte(key), gIV, decodeBytes))
}

// AesEncryptStruct 加密
func AesEncryptStruct(key string, ori interface{}) (string, error) {
	if key == "" {
		return "", fmt.Errorf(" encrypt err,key is nil")
	}
	b, err := json.Marshal(ori)
	if err != nil {
		return "", err
	}
	// log.Debugf("AesEncryptStruct = key: %v ,iv: %v", key, gIV)
	encrypt := AesEncrypt([]byte(key), gIV, b)
	encodeString := base64.StdEncoding.EncodeToString(encrypt)
	return encodeString, nil
}

// AesDecryptStruct 解密
func AesDecryptStruct(key, encryptData string, ori interface{}) error {
	if key == "" {
		return fmt.Errorf(" decrypt err,key is nil")
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return err
	}
	resJson := AesDecrypt([]byte(key), gIV, decodeBytes)
	err = json.Unmarshal(resJson, ori)
	if err != nil {
		return err
	}
	return nil
}
