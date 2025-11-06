package utils

import "encoding/hex"

// Password 密码算法 t 为 md5 的加密次数 暂定为前端加密一次给后端 后端再加密一次
func Password(s []byte, t int) string {
	for i := 0; i < t; i++ {
		s = MD5(s)
	}
	return hex.EncodeToString(s)
}
