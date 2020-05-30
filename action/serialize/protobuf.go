package serialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/models/serialize/userinfo"
)

func Protobuf(ctx *gin.Context) {
	ctx.ProtoBuf(http.StatusOK, write())
}

func write() *userinfo.ContactBook {
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
