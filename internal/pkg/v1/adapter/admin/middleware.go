package admin

// // 服务地址
// const v1 = "/v1/extension/"

// // mHadLogin 强制验证登录
// func mHadLogin(h gate.HandlerFunc) gate.HandlerFunc {
// 	return func(req *gate.Request) error {
// 		_initCtxMustLogin(req)

// 		if domain.GetCompanyID(req.Ctx) < 1 || domain.GetUserID(req.Ctx) < 1 {
// 			err := fmt.Errorf("userId 或者 companyId 小于1")
// 			ulog.Warn(req.Ctx, err)
// 			return err
// 		}

// 		err := h(req)
// 		if err != nil {
// 			ulog.Warnf(req.Ctx, "中间件err:%v", err.Error())
// 			feishuAlert(req, err)
// 		}
// 		return err
// 	}
// }

// // Sign 处理用户和权限
// func Sign(h gate.HandlerFunc) gate.HandlerFunc {
// 	return func(req *gate.Request) error {
// 		utils.InitGwTraceID(req, req.GetRawRequest().Header)
// 		ulog.Infof(req.Ctx, "path:%s, userid:%d, header:%+v, query:%s, body:%s",
// 			req.GetRawRequest().Path, req.GetRawRequest().GetUserID(),
// 			req.GetRawRequest().Header,
// 			req.GetRawRequest().Query, string(req.GetRawRequest().Body))
// 		// 使用cpid和cpkey验证权限
// 		if !getSign(req.Ctx, req.GetRawRequest().Header) {
// 			return domain.ErrUnauthorized
// 		}
// 		return h(req)
// 	}
// }

// // getSign 获取与cp通信的签名
// func getSign(cxt context.Context, header map[string]string) bool {
// 	traceID, ts, sign := "", "", ""
// 	var cpid uint32
// 	defer func() {
// 		ulog.Infof(cxt, "traceID:%s, ts:%s, sign:%s, cpid:%d", traceID, ts, sign, cpid)
// 	}()

// 	// 请求标识
// 	if _traceID, ok := header["ruixue-traceid"]; ok {
// 		traceID = _traceID
// 	}
// 	// 时间戳
// 	if _ts, ok := header["ruixue-cpts"]; ok {
// 		ts = _ts
// 	}
// 	// 签名
// 	if _sign, ok := header["ruixue-cpsign"]; ok {
// 		sign = _sign
// 	}
// 	if traceID == "" || ts == "" || sign == "" {
// 		return false
// 	}

// 	if _sign := rxkit.GetCPSign(traceID, ts, config.Data().OpenCompany.SecretVal); _sign != sign {
// 		ulog.Infof(cxt, "headerSign:%s, calcSign:%s", sign, _sign)
// 		return false
// 	}
// 	return true
// }

// // 无需登录接口的中间件
// func middlewareNologin(h gate.HandlerFunc) gate.HandlerFunc {
// 	return func(req *gate.Request) error {
// 		utils.InitGwTraceID(req, req.GetRawRequest().Header)

// 		req.Ctx = context.WithValue(req.Ctx, domain.Language, req.Language)

// 		ulog.Infof(req.Ctx, "middleware:nologin,Method:%s, path:%s, sessionId:%s, query:%s, body:%s",
// 			req.GetRawRequest().Method,
// 			req.GetRawRequest().Path,
// 			req.GetRawRequest().SessionID,
// 			req.GetRawRequest().Query,
// 			utils.FormatBodyToString(req.GetRawRequest().Body))

// 		ulog.Infof(req.Ctx, "header_detail:%+v", utils.FormatHeader(req.GetRawRequest().Header))

// 		err := h(req)
// 		if err != nil {
// 			ulog.Warnf(req.Ctx, "中间件err:%v", err.Error())
// 			if !errors.Is(err, domain.ErrUnlogin) {
// 				feishuAlert(req, err)
// 			}
// 		}
// 		return err
// 	}
// }

// func _initCtxMustLogin(req *gate.Request) {
// 	utils.InitGwTraceID(req, req.GetRawRequest().Header)

// 	req.Ctx = context.WithValue(req.Ctx, domain.Usertype, req.GlobalSessions.GetInt64("usertype"))
// 	req.Ctx = context.WithValue(req.Ctx, domain.UserID, req.GlobalSessions.GetInt64("userid"))
// 	req.Ctx = context.WithValue(req.Ctx, domain.Username, req.GlobalSessions.Get("contacts"))
// 	req.Ctx = context.WithValue(req.Ctx, domain.CompanyID, req.GlobalSessions.GetInt64("company_id"))
// 	req.Ctx = context.WithValue(req.Ctx, domain.CompanyName, req.GlobalSessions.Get("company_name"))
// 	req.Ctx = context.WithValue(req.Ctx, domain.Language, req.Language)
// 	xlog.Infof("head:%v", req.GetRawRequest().Header)
// 	ulog.Infof(req.Ctx, "Method:%s, name:%s, userId:(%d),userType:(%d) path:%s, sessionId:%s, query:%s, body:%s",
// 		req.GetRawRequest().Method,
// 		req.GlobalSessions.GetString("name"),
// 		req.GetRawRequest().GetUserID(),
// 		req.GlobalSessions.GetInt64("usertype"),
// 		req.GetRawRequest().Path,
// 		req.GetRawRequest().SessionID,
// 		req.GetRawRequest().Query,
// 		strings.ReplaceAll(string(req.GetRawRequest().Body), "\r\n", ""))
// }

// // 瑞雪 商业化标签 相关接口直接透传中间件
// func mRuixueDirect(h gate.HandlerFunc) gate.HandlerFunc {
// 	return func(req *gate.Request) error {
// 		_initCtxMustLogin(req)

// 		ulog.Infof(req.Ctx, "请求大数据接口, METHOD: %s, path:%s, body:%s", req.GetRawRequest().Method,
// 			req.GetRawRequest().Path,
// 			string(req.GetRawRequest().Body))

// 		if domain.GetCompanyID(req.Ctx) < 1 || domain.GetUserID(req.Ctx) < 1 || domain.GetPrivateDomainID(req.Ctx) < 1 {
// 			err := fmt.Errorf("userId 或者 companyId 或者 PrivateDomainID 小于1")
// 			ulog.Warn(req.Ctx, err)
// 			feishuAlert(req, err)
// 			return err
// 		}

// 		if GetProjectId(req) < 1 {
// 			ulog.Warn(req.Ctx, domain.ErrProjectIdInvalid)
// 			feishuAlert(req, domain.ErrProjectIdInvalid)
// 			return domain.ErrProjectIdInvalid
// 		}

// 		err := h(req)
// 		if err != nil {
// 			ulog.Warnf(req.Ctx, "中间件err:%v", err.Error())
// 			if !errors.Is(err, domain.ErrUnlogin) {
// 				feishuAlert(req, err)
// 			}
// 		}
// 		return err
// 	}
// }

// func GetProjectId(req *gate.Request) (resp int64) {
// 	defer func() {
// 		if resp == 0 {
// 			utils.SendFeishuMessage(req.Ctx, config.Data().Notice.ApiError, "从Header解析项目ID失败", []*utils.FeishuMessage{
// 				{
// 					Tag:  "text",
// 					Text: fmt.Sprintf("%s\nMETHOD:%s, path:%s\n", domain.GetDesc(req.Ctx), req.GetRawRequest().Method, req.GetRawRequest().Path),
// 				},
// 				{
// 					Tag:  "text",
// 					Text: fmt.Sprintf("header:%+v\n", req.GetRawRequest().Header),
// 				},
// 			})
// 		}
// 	}()
// 	// siyu-project-id  int
// 	for k, v := range req.GetRawRequest().Header {
// 		if strings.ToLower(k) == "siyu-project-id" {
// 			projectId, err := strconv.ParseInt(v, 10, 64)
// 			if err != nil {
// 				ulog.Warnf(req.Ctx, "解析项目id(%s)失败, err: %v, key: %s, val:%s", k, err, k, v)
// 				return 0
// 			}
// 			if projectId == 0 {
// 				ulog.Warnf(req.Ctx, "解析项目id等于0, (%s)失败, err: %v, key: %s, val:%s", k, err, k, v)
// 			}
// 			return projectId
// 		}
// 	}

// 	ulog.Warn(req.Ctx, "从header中未解析到项目id的key (siyu-project-id)")
// 	return 0
// }

// func feishuAlert(req *gate.Request, err error) {
// 	utils.SendFeishuMessage(req.Ctx, config.Data().Notice.ApiError,
// 		"中间件报错, METHOD: "+req.GetRawRequest().Method+", path:"+req.GetRawRequest().RawURL, []*utils.FeishuMessage{
// 			{
// 				Tag:  "text",
// 				Text: fmt.Sprintf("err:%+v\n", err),
// 			},
// 			{
// 				Tag:  "text",
// 				Text: fmt.Sprintf("query: %s\n", req.GetRawRequest().Query),
// 			},
// 			{
// 				Tag:  "text",
// 				Text: fmt.Sprintf("body: \n%s\n", string(req.GetRawRequest().Body)),
// 			},
// 			{
// 				Tag: "text",
// 				Text: fmt.Sprintf("%s\n",
// 					domain.GetDesc(req.Ctx),
// 				),
// 			},
// 		})
// }

// // PrivateDomainSign 私域验签
// func PrivateDomainSign(h gate.HandlerFunc) gate.HandlerFunc {
// 	return func(req *gate.Request) error {

// 		req.Ctx = context.WithValue(req.Ctx, domain.CompanyID, int64(config.Data().OpenCompany.CompanyID))
// 		utils.InitGwTraceID(req, req.GetRawRequest().Header)

// 		ulog.Infof(req.Ctx, "path:%s,  query:%s, body:%s",
// 			req.GetRawRequest().Path,
// 			req.GetRawRequest().Query,
// 			string(req.GetRawRequest().Body))

// 		ulog.Infof(req.Ctx, "header_detail:%+v", utils.FormatHeader(req.GetRawRequest().Header))

// 		if !privateDomainGetSign(req.Ctx, req.GetRawRequest().Header) {
// 			ulog.Warnf(req.Ctx, "验签失败, %v", domain.ErrUnauthorized)
// 			// dev 放开
// 			if config.Data().Mode != "test" {
// 				feishuAlert(req, domain.ErrUnauthorized)
// 				return domain.ErrUnauthorized
// 			}
// 		}
// 		err := h(req)
// 		if err != nil {
// 			feishuAlert(req, err)
// 		}
// 		return err
// 	}
// }

// // 私域 验签
// func privateDomainGetSign(cxt context.Context, header map[string]string) bool {
// 	traceID, ts, sign := "", "", ""
// 	defer func() {
// 		ulog.Infof(cxt, "traceID:%s, ts:%s, sign:%s, SecretVal:%s", traceID, ts, sign, config.Data().OpenCompany.SecretVal)
// 	}()

// 	// 请求标识
// 	if _traceID, ok := header["ruixue-traceid"]; ok {
// 		traceID = _traceID
// 	}
// 	// 时间戳
// 	if _ts, ok := header["ruixue-cpts"]; ok {
// 		ts = _ts
// 	}
// 	// 签名
// 	if _sign, ok := header["ruixue-cpsign"]; ok {
// 		sign = _sign
// 	}
// 	if traceID == "" || ts == "" || sign == "" {
// 		return false
// 	}

// 	if _sign := rxkit.GetCPSign(traceID, ts, config.Data().OpenCompany.SecretVal); _sign != sign {
// 		ulog.Infof(cxt, "headerSign:%s, calcSign:%s", sign, _sign)
// 		return false
// 	}
// 	return true
// }
