package crypto

import (
	"testing"

	"github.com/oyjjpp/blog/util/constant"
)

const (
	CRYTO = "golang"
)

func TestMd5(t *testing.T) {
	rs := Md5Hash([]byte("ouyangjun"))
	t.Log(rs)
	t.Log(len(rs))
	res := Md5("ouyangjun", true)
	t.Log(res)
	t.Log(len(res))
}

func TestSha(t *testing.T) {
	rs := Sha256Hash([]byte("ouyangjun"))
	t.Log(rs)
	t.Log(len(rs))
	res := Sha256MacHash([]byte("ouyangjun"), []byte(CRYTO))
	t.Log(res)
	t.Log(len(res))
}

