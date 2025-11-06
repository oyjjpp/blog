package domain

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type ctxKey string

const (
	// 接口调用来源
	ApiSourceRuixue      string = "ruixue"
	ApiSourcePrivate     string = "private_domain"
	ApiSourceOpenCompany string = "open_company"
	ApiSourceExtend      string = "extension"
)

// 会话数据
const (
	Language         ctxKey = "language" //  Language 客户端语言
	UserID           ctxKey = "userid"
	Usertype         ctxKey = "usertype"
	Username         ctxKey = "username"
	UserEmail        ctxKey = "email"
	CompanyID        ctxKey = "company_id"
	CompanyName      ctxKey = "company_name"
	AccountType      ctxKey = "account_type"
	ProjectID        ctxKey = "project_id"
	PrivateDomainID  ctxKey = "private_domain_id"
	ProjectIdByLogin ctxKey = "project_id_by_login" // 取值前需要赋值，
)

// 时间格式
const (
	DateFormatLayout        string = "2006-01-02"
	DateTimeFormatLayout    string = "2006-01-02 15:04:05"
	ZeroDateTime            string = "0000-00-00 00:00:00"
	DateTimeLogFormatLayout string = "2006-01-02 15:04:05.000"
	SQLDefaultWhere         string = " AND `deleted_at`='0000-00-00 00:00:00'"
	DefaultPassport         string = "Rx@1234567"
)

const (
	UsertypeGw                 = 0
	UsertypeExtendJigou        = 4
	UsertypeExtendTuiguangyuan = 5
)

// RespCommon 返回通用结构
type RespCommon struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// GetCompanyID 获取操作者账号的公司id
func GetCompanyID(ctx context.Context) int64 {
	companyID := ctx.Value(CompanyID)
	if companyID == nil || companyID.(int64) == 0 {
		return 0
	}
	return companyID.(int64)
}

func GetCompanyName(ctx context.Context) string {
	companyName := ctx.Value(CompanyName)
	if companyName == nil || companyName.(string) == "" {
		return ""
	}
	return companyName.(string)
}

func GetPrivateDomainID(ctx context.Context) int64 {
	privateDomainId := ctx.Value(PrivateDomainID)
	if privateDomainId == nil || privateDomainId.(int64) == 0 {
		return 0
	}
	return privateDomainId.(int64)
}

// GetUserID 获取操作者id
func GetUserID(ctx context.Context) int64 {
	userID := ctx.Value(UserID)
	if userID == nil || userID.(int64) == 0 {
		return 0
	}
	return userID.(int64)
}

func GetUsertype(ctx context.Context) int64 {
	userID := ctx.Value(Usertype)
	if userID == nil || userID.(int64) == 0 {
		return 0
	}
	return userID.(int64)
}

func GetLanguage(ctx context.Context) string {
	val, _ := ctx.Value(Language).(string)
	return val
}

func GetUsername(ctx context.Context) string {
	val, _ := ctx.Value(Username).(string)
	return val
}

func GetDesc(ctx context.Context) string {
	return fmt.Sprintf("userId:%v (%v)(%v), companyId:%v (%v), language:%v",
		GetUserID(ctx),
		GetUsertype(ctx),
		GetUsername(ctx),
		GetCompanyID(ctx),
		GetCompanyName(ctx),
		GetLanguage(ctx),
	)
}

// NowTime 当前时间
func NowTime() *time.Time {
	t := time.Now()
	return &t
}

// NowTimeString 当前时间字符串
func NowTimeString() string {
	return NowTime().Format(DateTimeFormatLayout)
}

// NowDateString 当前日期字符串
func NowDateString() string {
	return NowTime().Format(DateFormatLayout)
}

type CommonArgs struct {
	UserId                int64
	PrivateDomain         int64
	CompanyId             int64
	ProjectId             int64
	PortraitId            int64
	ActiveReachStrategyId int64
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserSource int64

const (
	UserSourceOpenCompany UserSource = 1
	UserSourceExtension   UserSource = 2
)

func GetUserSource(userId int64) UserSource {
	if userId > 10000000 {
		return UserSourceExtension
	}
	return UserSourceOpenCompany
}

const (
	// 删除标志 0 未删除 1 已删除
	DelFlagNot int = iota
	DelFlagSuccess
)

func SqlLike(s string) string {
	s = strings.ReplaceAll(s, "%", `\%`)
	s = strings.ReplaceAll(s, "_", `\_`)
	return s
}

func LogTime(ctx context.Context, key string) func() {
	// ts := time.Now()
	return func() {
		// ulog.Infof(ctx, "执行耗时: %s ms, key:%s ", time.Since(ts).Milliseconds(), key)
	}
}

func CopyCtx(ctx context.Context) context.Context {
	newCtx := context.Background()
	// newCtx = context.WithValue(newCtx, ulog.ApiLogTraceID, ulog.GetApiTraceId(ctx))
	newCtx = context.WithValue(newCtx, Usertype, GetUsertype(ctx))
	newCtx = context.WithValue(newCtx, UserID, GetUserID(ctx))
	newCtx = context.WithValue(newCtx, Username, GetUsername(ctx))
	newCtx = context.WithValue(newCtx, CompanyID, GetCompanyID(ctx))
	newCtx = context.WithValue(newCtx, CompanyName, GetCompanyName(ctx))
	newCtx = context.WithValue(newCtx, Language, GetLanguage(ctx))
	return newCtx
}
