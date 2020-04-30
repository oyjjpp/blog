package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

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

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize //需要padding的数目
	//只要少于256就能放到一个byte中，默认的blockSize=16(即采用16*8=128, AES-128长的密钥)
	//最少填充1个byte，如果原文刚好是blocksize的整数倍，则再填充一个blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding) //生成填充的文本
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func EcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func EcbEncrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
