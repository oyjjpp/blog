package user

import (
	"errors"
	"net/http"

	"github.com/CodeLineage/tool"
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/global"
	"github.com/oyjjpp/blog/models"
	"github.com/oyjjpp/blog/util"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
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

// 批量添加
func BatchCreate(ctx *gin.Context) {
	// var users models.T
	city := []string{"哈尔宾", "长春", "沈阳", "大连", "石家庄", "杭州","济南","太原","郑州","广州","海南","西安","大同"}

	age := []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}

	name := []string{"龙与", "雪迎", "赵启慧", "妮子", "如男", "高飞", "果类", "迎非"}
	
	users := make([]models.T,0, 25000)
	for i := 5000; i < 25000; i++ {
        cityIndex := tool.Rand(0, 12)
        ageIndex := tool.Rand(0, 9)
        nameIndex := tool.Rand(0, 7)
        
		userInfo := models.T{
			Id:   i + 1,
			City: city[cityIndex],
			Age:  age[ageIndex],
			Name: name[nameIndex],
			Addr: "",
		}
        users = append(users, userInfo)
	}
	data := global.MysqlDB.Table("t").Create(&users)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500204,
			"msg":  data.Error.Error(),
			"body": map[string]interface{}{},
		})
	} else if data.Error == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  "sucess",
			"body": data,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  data.Error.Error(),
			"body": map[string]interface{}{},
		})
	}
}
