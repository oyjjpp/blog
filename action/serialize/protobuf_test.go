package serialize

import (
	"io/ioutil"
	"testing"

	"github.com/oyjjpp/blog/models/serialize/userinfo"
	"google.golang.org/protobuf/proto"
)

func TestReadProtobuf(t *testing.T) {
	data, err := ioutil.ReadFile("./protobuf_test.log")
	if err != nil {
		t.Fatal(err.Error())
	}
	book := &userinfo.ContactBook{}
	if err := proto.Unmarshal(data, book); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(book)
}
