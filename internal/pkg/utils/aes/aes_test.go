package aes

import (
	"testing"
)

const (
	CRYPT_KEY_256 = "1~$c31kjtR^@@c2#9&iy"
	CRYPT_KEY_128 = "c31kjtR^@@c2#9&"
)

type testInfo struct {
	key      string //秘钥
	iv       string //iv
	origData string //原文
}

func TestAesCbc256(t *testing.T) {
	tests := []*testInfo{
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "704114497615264",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es0414497615272",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es704144222222297615406",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertysdddddduiopzxcvbnmgqo",
			origData: "df041449712619896",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwssaaertyuiopzxcvbnmgqo",
			origData: "fd041449768759744",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ds41449722659772",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ff4144977fe304674",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "vfr414497615388",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyu",
			origData: "bth14497615418",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "123456qo",
			origData: "4rb2auut300790842620672",
		},
	}

	for index, test := range tests {
		encrData := AesEncrypt([]byte(test.key), []byte(test.iv), []byte(test.origData))
		origData := AesDecrypt([]byte(test.key), []byte(test.iv), encrData)
		if string(origData) != test.origData {
			t.Error(index, " fail")
		}
	}
}

func TestAesCbc128(t *testing.T) {
	tests := []*testInfo{
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwerty",
			origData: "704114497615264",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890iopzxc",
			origData: "es0414497615272",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es704144222222297615406",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertysdddddduiopzxcvbnmgqo",
			origData: "df041449712619896",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwssaaertyuiopzxcvbnmgqo",
			origData: "fd041449768759744",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ds41449722659772",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ff4144977fe304674",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "vfr414497615388",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyu",
			origData: "bth14497615418",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "123456qo",
			origData: "4rb2auut300790842620672",
		},
	}

	for index, test := range tests {
		encrData := AesEncrypt([]byte(test.key), []byte(test.iv), []byte(test.origData))
		t.Log(encrData)
		origData := AesDecrypt([]byte(test.key), []byte(test.iv), encrData)
		if string(origData) != test.origData {
			t.Error(index, " fail")
		}
	}
}

func TestAesEncryptWithIV(t *testing.T) {
	type args struct {
		key      string
		origData string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"271828", "for test only"}, ""},
		{"test1", args{"920", "for test 2 only"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AesEncryptWithIV(tt.args.key, tt.args.origData)
			t.Log("密文 ", got)

			ori := AesDecryptWithIV(tt.args.key, got)
			t.Log("原文： ", ori)
		})
	}
}

func TestAesEncryptStruct(t *testing.T) {
	type MyName struct {
		Name string
		Age  int
	}

	type args struct {
		key string
		ori MyName
	}
	tests := []struct {
		name string
		args args
	}{

		{"test1", args{"27182", MyName{"linHC", 34}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesEncryptStruct(tt.args.key, tt.args.ori)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(got)
			var myName MyName
			err = AesDecryptStruct(tt.args.key, got, &myName)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(myName)
		})
	}
}
