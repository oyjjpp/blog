// Copyright (c) 2022. Homeland Interactive Technology Ltd. All rights reserved.

package zh

import (
	"golang.org/x/text/language"

	"git.jiaxianghudong.com/ruixue/rxkit/i18n"
)

var pack = i18n.Pack{
	// domain.ErrInternalSystem.Error():       "系统内部错误",
	// domain.ErrBodyInvalid.Error():          "请求参数体无效",
	// domain.ErrQueryInvalid.Error():         "请求参数无效",
	// domain.ErrHeaderInvalid.Error():        "请求头信息无效",
	// domain.ErrNoPermission.Error():         "没有权限",
	// domain.ErrNoRecords.Error():            "没有记录",
	// domain.ErrGetUserInfoFailed.Error():    "获取用户信息失败",
	// domain.ErrGetCompanyInfoFailed.Error(): "获取企业信息失败",
	// domain.ErrGetDataFailed.Error():        "获取数据失败",
	// domain.ErrSignExpired.Error():          "超过签名有效期",
	// domain.ErrSignError.Error():            "签名错误",
	// domain.ErrUserNotExist.Error():         "用户不存在",
	// domain.ErrOriginalPassword.Error():     "请输入正确的原密码",
	// domain.ErrPasswordNotAllowSame.Error(): "新密码和旧密码不能一致",
}

func init() {
	i18n.SetPackWithTag(language.Chinese, pack)
	i18n.SetPackWithTag(language.SimplifiedChinese, pack)
}
