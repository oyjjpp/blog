package serialize

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/models/serialize/userinfo"
	"google.golang.org/protobuf/proto"
)

func Protobuf(ctx *gin.Context) {
	ctx.ProtoBuf(http.StatusOK, productData())
}

// productData
// 生产数据
func productData() *userinfo.ContactBook {
	per1 := &userinfo.Person{
		Id:   1,
		Name: "赵",
		Phones: []*userinfo.Phone{
			{Type: userinfo.PhoneType_HOME, Number: "18243089001"},
			{Type: userinfo.PhoneType_WORK, Number: "18243089002"},
		},
	}

	per2 := &userinfo.Person{
		Id:   2,
		Name: "钱",
		Phones: []*userinfo.Phone{
			{Type: userinfo.PhoneType_HOME, Number: "18243089003"},
			{Type: userinfo.PhoneType_WORK, Number: "18243089004"},
		},
	}

	book := &userinfo.ContactBook{}
	book.Persons = append(book.Persons, per1)
	book.Persons = append(book.Persons, per2)
	return book
}

func write() {
	book := productData()
	// 编码数据
	data, _ := proto.Marshal(book)
	// 把数据写入文件
	_ = ioutil.WriteFile("./protobuf_test.log ", data, os.ModePerm)
}
