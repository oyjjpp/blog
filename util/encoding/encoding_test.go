package encoding

import (
	"testing"

	"github.com/oyjjpp/blog/util/constant"
)

func TestPEMDecode(t *testing.T) {
	block, _ := PEMDecode(constant.PRIVATE_KEY)
	if block == nil {
		t.Error("公钥异常")
	}
	t.Log(block.Type)
	t.Log(block.Headers)
	t.Log(string(block.Bytes))
}
