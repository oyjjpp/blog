package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(b []byte) []byte {
	h := md5.New()
	h.Write(b)
	return h.Sum(nil)
}

func CalculateUpperMD5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(hashBytes))

}
