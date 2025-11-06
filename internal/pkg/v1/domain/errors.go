// Copyright (c) 2022. Homeland Interactive Technology Ltd. All rights reserved.

package domain

// func init() {
// 	consts.SetServiceCode(consts.SCodeProduction)
// }

// // 基础设施错误码
// const (
// 	CodeBaseError = consts.ErrCodeTypeBase + consts.SCodeProduction + (1 + iota)
// 	CodeDBError
// 	CodeRedisError
// )

// // 业务内部错误码
// const (
// 	CodeInternalError = consts.ErrCodeTypeService + consts.SCodeProduction + (1 + iota)
// 	CodeUserInvalid
// )

// // RPC 错误码
// const (
// 	CodeRPCRequesError = consts.ErrCodeTypeRPCRequest + consts.SCodeProduction + (1 + iota)
// )

// // RPC 参数错误
// const (
// 	CodeRPCArgument = consts.ErrCodeTypeRPCArgument + consts.SCodeProduction + (1 + iota)
// )

// // 用户级错误码
// const (
// 	CodeArgumentInvalid = consts.ErrCodeTypeUserArgument + consts.SCodeProduction + (1 + iota)
// 	CodeQueryInvalid
// 	CodeBodyInvalid
// 	CodeHeaderInvalid
// 	CodeNoPermission
// 	CodeNoRecords
// 	CodeUnauthorized
// 	CodeGetProjectInfoFailed
// 	CodeGetUserInfoFailed
// 	CodeGetCompanyInfoFailed
// 	CodeGetDataFailed
// 	CodeProjectIdInvalid
// 	CodeSignExpired
// 	CodeSignError
// 	CodeResetPassword
// 	CodeUserNotExist
// 	CodePasswordNotAllowSame
// )

// // 错误码  305204
// var (
// 	ErrUnlogin          = eerr.DefineCode(325001, "unlogin")
// 	ErrInternalSystem   = eerr.DefineCode(CodeInternalError, "internal error")
// 	ErrRPCUserIDInvalid = eerr.DefineCode(CodeRPCArgument, "invalid userid")
// 	// 用户级错误, 需要翻译
// 	ErrQueryInvalid         = eerr.DefineCode(CodeQueryInvalid, "request query invalid")
// 	ErrBodyInvalid          = eerr.DefineCode(CodeBodyInvalid, "request body invalid")
// 	ErrHeaderInvalid        = eerr.DefineCode(CodeHeaderInvalid, "request header invalid")
// 	ErrNoPermission         = eerr.DefineCode(CodeNoPermission, "no permission")
// 	ErrNoRecords            = eerr.DefineCode(CodeNoRecords, "no records")
// 	ErrUnauthorized         = eerr.DefineCode(CodeUnauthorized, "unauthorized")
// 	ErrGetProjectInfoFailed = eerr.DefineCode(CodeGetProjectInfoFailed, "get project info failed")  // 获取项目信息失败
// 	ErrGetUserInfoFailed    = eerr.DefineCode(CodeGetUserInfoFailed, "get user info failed")        // 获取用户信息失败
// 	ErrGetCompanyInfoFailed = eerr.DefineCode(CodeGetCompanyInfoFailed, "get company info failed")  // 获取企业信息失败
// 	ErrGetDataFailed        = eerr.DefineCode(CodeGetDataFailed, "get data failed")                 // 获取数据失败
// 	ErrProjectIdInvalid     = eerr.DefineCode(CodeProjectIdInvalid, "project id is invalid")        // 项目id必传
// 	ErrSignExpired          = eerr.DefineCode(CodeSignExpired, "Expiration date of signature")      // 超过签名有效期
// 	ErrSignError            = eerr.DefineCode(CodeSignError, "sign error")                          // 签名错误
// 	ErrUserNotExist         = eerr.DefineCode(CodeUserNotExist, "user not exists")                  // 用户不存在
// 	ErrOriginalPassword     = eerr.DefineCode(CodeArgumentInvalid, "abnormal original password")    // 原密码异常
// 	ErrPasswordNotAllowSame = eerr.DefineCode(CodePasswordNotAllowSame, "password not allow samed") // 密码不允许相同
// 	ErrUpdatePassword       = eerr.DefineCode(CodeArgumentInvalid, "update password invalid")       // 密码为空的提示
// )

// // WrapCode 包装错误信息
// func WrapCode(cause error, errCode int, msg ...string) error {
// 	if eerr.HasType(eerr.Code(cause)) {
// 		return cause
// 	}
// 	xlog.Error(cause)
// 	return eerr.WrapCode(cause, errCode, msg...)
// }

// // T 翻译错误
// func T(ctx context.Context, err error, language ...string) string {
// 	_language := i18n.DefaultLocaleName
// 	if len(language) > 0 {
// 		_language = language[0]
// 	} else if ctx.Value(Language) != nil {
// 		if tmp, ok := ctx.Value(Language).(string); ok {
// 			_language = tmp
// 		}
// 	}
// 	l := i18n.GetLocale(_language)
// 	if l == nil {
// 		return err.Error()
// 	}
// 	msg := l.T(eerr.Message(err))
// 	if msg == "" {
// 		return err.Error()
// 	}
// 	return msg
// }
