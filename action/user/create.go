package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/global"
	"github.com/oyjjpp/blog/models"
	"github.com/oyjjpp/blog/util"
	uuid "github.com/satori/go.uuid"
)

// User register structure
type RegisterStruct struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// Create
func Create(ctx *gin.Context) {
	var R RegisterStruct
	_ = ctx.ShouldBindJSON(&R)

	user := &models.SysUser{
		Username:    R.Username,
		NickName:    R.NickName,
		Password:    R.Password,
		HeaderImg:   R.HeaderImg,
		AuthorityId: R.AuthorityId,
	}
	register(user)
	global.MysqlDB.Create(&user)
}

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser
func register(u *models.SysUser) (userInter *models.SysUser, err error) {
	var user models.SysUser
	//判断用户名是否注册
	data := global.MysqlDB.Where("username = ?", u.Username).First(&user)
	//notRegister为false表明读取到了 不能注册
	if data != nil {
		return userInter, errors.New("用户名已注册")
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = util.Md5(u.Password)
		u.UUID = uuid.NewV4()
		err = global.MysqlDB.Create(&u).Error
	}
	return u, err
}
