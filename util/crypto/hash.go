// 消息摘要算法
// 特征：加密过程不需要秘钥、加密的数据很能得到破解
// 作用：进行数字签名
// 常见算法：MD5、SHA
package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// Md5Hash
// link https://mp.weixin.qq.com/s/k-ToL356asWtS_PN30Z17w?
// MD5 生成的哈希值是128位的二进制数，也就是32位的十六进制数
// 作用：一般用来做签名
// 底层原理： 处理原文/设置初始值/循环加工/拼接结果
// MD5 hash 算法
func Md5Hash(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	ciphertext := md5Ctx.Sum(nil)
	// 将数据编码为字符串
	return hex.EncodeToString(ciphertext)
}

// Md5 Calculate the md5 hash of a string
// If the optional rawOutput is set to true, then the md5 digest is instead returned in raw binary format with a length of 16.
func Md5(str string, rawOutput ...bool) string {
	h := md5.New()
	h.Write([]byte(str))
	str = hex.EncodeToString(h.Sum(nil))
	if len(rawOutput) > 0 && rawOutput[0] {
		str = str[8:24]
	}
	return str
}

// Sha256Hash
// SHA256 hash 算法
func Sha256Hash(data []byte) string {
	sha256Ctx := sha256.New()
	sha256Ctx.Write(data)
	ciphertext := sha256Ctx.Sum(nil)
	// 将数据编码为字符串
	return hex.EncodeToString(ciphertext)
}

// Sha256MacHash
// MAC sha256 hash 算法
func Sha256MacHash(data, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	ciphertext := mac.Sum(nil)
	return hex.EncodeToString(ciphertext)
}
