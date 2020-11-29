package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5
func Md5(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return hex.EncodeToString(h.Sum(nil))
}
