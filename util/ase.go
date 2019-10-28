package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func hash_hmac() {
	secret := "2QmS3P5UwxgVH8xKIMAxnvVvzJojmneWU1CxJNjztyXiJystqEhh2S2K2EmLBLO9lmm8fiKOuq85lyTros1CP0j2GMpFX9VuO5mDhURMmFdmnh9fBTA6HSJ5rs30Xkg3"
	data := `/api/huawei/reader/getUserAsset{"timestamp":"20191025044022","data":{"userId":"70852000004110267","getAsset":1}}`
	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)
	encodeString := base64.StdEncoding.EncodeToString([]byte(sha))
	fmt.Println(encodeString)
}

//DES 加密
//data []byte 原始数据
//key []byte 密钥 长度必须为8
func DesEncrpty(data, key []byte) ([]byte, error) {
	//根据密钥生成block
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	data = PKCS5Padding(data, blockSize)
	//根据block和初始向量生成加密算法，IV长度与block.size需要保持一致
	blockMode := cipher.NewCBCEncrypter(block, key)

	crypted := make([]byte, len(data))
	//data 加密
	blockMode.CryptBlocks(crypted, data)
	return crypted, nil
}

//DES 解密
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	//对密文解密
	data := make([]byte, len(crypted))
	blockMode.CryptBlocks(data, crypted)
	//反扩充，获取原始明文
	data = PKCS5UnPadding(data)
	return data, nil
}

//AES加密 使用CBC模式 PKCS7填充方式
// data 加密数据
// key 为密钥 16: AES-128 24 : AES-192 32 AES-256
func AesEncrpty(data, key []byte) ([]byte, error) {
	//创建一个cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//加密字节块的大小
	blockSize := block.BlockSize()
	//通过PKCS7方式填充明文数据
	data = PKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	//加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
	blockMode.CryptBlocks(crypted, data)
	return crypted, nil
}

//ASE 解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	data := make([]byte, len(crypted))
	blockMode.CryptBlocks(data, crypted)
	data = PKCS7UnPadding(data)
	return data, nil
}

//通过PKCS7方式填充明文数据
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//通过PKCS7方式将明文数据的填充值去掉
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

//通过PKCS5方式填充明文数据
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	println(blockSize, len(ciphertext), padding)
	println("padding=" + string(padding) + "test")
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	println(string(padtext))
	return append(ciphertext, padtext...)
}

//通过PKCS5方式将明文数据的填充值去掉
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return []byte{}
	}
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if unpadding < 0 {
		return origData
	}
	println("unpadding=" + string(unpadding))

	return origData[:(length - unpadding)]
}
