package algorithm

import (
	"encoding/base64"
	"oyjblog/define"
	"oyjblog/util"

	"github.com/astaxie/beego"
)

// Operations about object
type EncryptionController struct {
	beego.Controller
}

//AES 加解密算法
func (encry *EncryptionController) AesEncry() {
	//加密的key
	key := encry.GetString("key")
	var ciptherKey []byte
	if key == "" {
		ciptherKey = []byte("0123456789abcdef")
	} else {
		ciptherKey = []byte(key)
	}

	//加密的value
	data := encry.GetString("data")
	if data == "" {
		data = "hello world"
	}

	result := make(map[string]interface{})

	code := 0
	msg := "success"
	body := make(map[string]interface{})

	//简单实现do while
	doIndex := 0
	for doIndex < 1 {
		doIndex++
		body["ciptherKeyLen"] = len(ciptherKey)
		body["data"] = data
		rs, err := util.AesEncrpty([]byte(data), ciptherKey)
		if err != nil {
			code = 100001
			msg = "加密错误"
			break
		}
		body["ciptherRs"] = rs
		body["ciptherRsLen"] = len(rs)
		base64Rs := base64.StdEncoding.EncodeToString(rs)
		body["base64Rs"] = base64Rs
		body["base64RsLen"] = len(base64Rs)
		origin, err := util.AesDecrypt(rs, ciptherKey)
		if err != nil {
			code = 100002
			msg = "解密错误"
			break
		}
		body["origin"] = string(origin)
	}

	result["code"] = code
	result["msg"] = msg
	result["body"] = body

	encry.Data["json"] = result
	encry.ServeJSON()
}

//DES 加解密算法
func (encry *EncryptionController) DesEncry() {
	//加密的key
	key := encry.GetString("key")
	var ciptherKey []byte
	if key == "" {
		ciptherKey = []byte("12345678")
	} else {
		ciptherKey = []byte(key)
	}

	//加密的value
	data := encry.GetString("data")
	if data == "" {
		data = "hello world"
	}

	result := make(map[string]interface{})

	code := 0
	msg := "success"
	body := make(map[string]interface{})

	//简单实现do while
	doIndex := 0
	for doIndex < 1 {
		doIndex++
		body["ciptherKeyLen"] = len(ciptherKey)
		body["data"] = data
		rs, err := util.DesEncrpty([]byte(data), ciptherKey)
		if err != nil {
			code = 100001
			msg = err.Error()
			break
		}
		body["ciptherRs"] = rs
		body["ciptherRsLen"] = len(rs)
		base64Rs := base64.StdEncoding.EncodeToString(rs)
		body["base64Rs"] = base64Rs
		body["base64RsLen"] = len(base64Rs)

		origin, err := util.DesDecrypt(rs, ciptherKey)
		//origin, err := util.DecryptDES_CBC(fmt.Sprintf("%X", rs), ciptherKey)
		if err != nil {
			code = 100002
			msg = err.Error()
			break
		}
		body["origin"] = string(origin)

	}

	result["code"] = code
	result["msg"] = msg
	result["body"] = body

	encry.Data["json"] = result
	encry.ServeJSON()
}

//RSA加密
func (encry *EncryptionController) RsaEncry() {
	privateKey := define.PRIVATE_KEY
	publicKey := define.PUBLIC_KEY

	//加密的value
	data := encry.GetString("data")
	if data == "" {
		data = "hello world"
	}

	result := make(map[string]interface{})

	code := 0
	msg := "success"
	body := make(map[string]interface{})

	//简单实现do while
	doIndex := 0
	for doIndex < 1 {
		doIndex++
		body["data"] = data
		rs, err := util.RsaPubEncrpty([]byte(publicKey), []byte(data))
		if err != nil {
			code = 100001
			msg = err.Error()
			break
		}
		body["ciptherRs"] = rs
		body["ciptherRsLen"] = len(rs)
		base64Rs := base64.StdEncoding.EncodeToString(rs)
		body["base64Rs"] = base64Rs
		body["base64RsLen"] = len(base64Rs)

		origin, err := util.RsaPriDecrypt([]byte(privateKey), []byte(rs))
		if err != nil {
			code = 100002
			msg = err.Error()
			break
		}
		body["origin"] = string(origin)

	}

	result["code"] = code
	result["msg"] = msg
	result["body"] = body

	encry.Data["json"] = result
	encry.ServeJSON()
}

//RSA签名
func (encry *EncryptionController) RsaSign() {
	privateKey := define.PRIVATE_KEY
	publicKey := define.PUBLIC_KEY

	//加密的value
	data := encry.GetString("data")
	if data == "" {
		data = "hello world"
	}

	result := make(map[string]interface{})

	code := 0
	msg := "success"
	body := make(map[string]interface{})

	//简单实现do while
	doIndex := 0
	for doIndex < 1 {
		doIndex++
		body["data"] = data
		rs, err := util.RsaPriEncrpty([]byte(privateKey), []byte(data))
		if err != nil {
			code = 100001
			msg = err.Error()
			break
		}
		body["ciptherRs"] = rs
		body["ciptherRsLen"] = len(rs)
		base64Rs := base64.StdEncoding.EncodeToString(rs)
		body["base64Rs"] = base64Rs
		body["base64RsLen"] = len(base64Rs)

		//验证签名
		err = util.RsaVerifySign([]byte(data), []byte(publicKey), base64Rs)
		if err != nil {
			code = 100002
			msg = err.Error()
			break
		}
	}

	result["code"] = code
	result["msg"] = msg
	result["body"] = body

	encry.Data["json"] = result
	encry.ServeJSON()
}
