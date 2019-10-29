package util

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
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
	data = PKCS7Padding(data, blockSize)
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
	data = PKCS7UnPadding(data)
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
	if length == 0 {
		return []byte{}
	}
	unpadding := int(data[length-1])
	if unpadding < 0 {
		return data
	}
	return data[:(length - unpadding)]
}

//RSA通过公钥加密
//publicKey 公钥
//data 要加密的数据
func RsaPubEncrpty(publicKey, data []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("pubkey key error!")
	}

	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	pub := pubInterface.(*rsa.PublicKey)

	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

//私钥解密
//privateKey 私钥
//ciphertext 要解密的数据
func RsaPriDecrypt(privateKey, ciphertext []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	//解析PKCS1格式私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	//解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//RSA私钥签名
func RsaPriEncrpty(privateKey, data []byte) ([]byte, error) {
	/*
		h := sha256.New()
		h.Write(data)
		hashed := h.Sum(nil)
	*/
	hashed := sha1.Sum(data)

	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	//解析PKCS1格式私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//加密
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hashed[:])
	//return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
}

//RSA通过公钥验证签名
func RsaVerifySign(data, pubkey []byte, sign string) error {
	hashed := sha1.Sum(data)
	sign2, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(pubkey)

	if block == nil {
		return errors.New("pubkey key error!")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	pub := pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed[:], []byte(sign2))
}
