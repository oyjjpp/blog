package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// RsaPubEncrpty
// RSA通过公钥加密
// publicKey 公钥
// data 要加密的数据
func RsaPubEncrpty(publicKey, data []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("pubkey key error!")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// 私钥解密
// privateKey 私钥
// ciphertext 要解密的数据
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

// RSA私钥签名
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
