package admin

// func BindUserAdapter(router *gate.APIApp, user domain.UserUseCase) {
// 	u := &userAdapter{
// 		user: user,
// 	}

// 	router.Post(v1+`user/login`, middlewareNologin(u.Login), gate.AuthTypeNoAuth|gate.ContentTypeJSON) // 用户登录
// 	router.Post(v1+`user/logout_login`, middlewareNologin(u.Logout), gate.AuthTypeNoAuth)              // 退出登录
// 	router.Post(v1+`permission/get_all`, mHadLogin(u.PermissionGetAll), gate.AuthTypeSessionID)        // 获取用户所有权限
// 	router.Post(v1+`user/set_password`, mHadLogin(u.ResetPassword), gate.AuthTypeSessionID)            // 重置密码
// 	router.Post(v1+`user/token/refresh_token`, middlewareNologin(u.RefreshToken), gate.AuthTypeNoAuth) // 刷新token
// 	router.Get(v1+`user/platform_list`, mHadLogin(u.PlatformList), gate.AuthTypeSessionID)             // 用户的所有权限列表
// 	router.Get(v1+`user/platform_link`, mHadLogin(u.PlatformLink), gate.AuthTypeSessionID)             // 获取项目跳转链接
// }

// type userAdapter struct {
// 	user domain.UserUseCase
// }

// // PlatformList 用户的所有项目权限列表
// func (api *userAdapter) PlatformList(req *gate.Request) error {
// 	userId := domain.GetUserID(req.Ctx)
// 	if userId < 1 {
// 		return domain.ErrQueryInvalid
// 	}

// 	urlStr := fmt.Sprintf("/api/v1/gwapi/public/user/platform_list?user_id=%d&source=%s&is_filter=1", userId, domain.ApiSourceExtend)
// 	val, err := utils.GwRawHttpGet(req.Ctx, config.Data().OpenCompany.InnerDomain, urlStr, domain.GetLanguage(req.Ctx))
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err)
// 		return err
// 	}

// 	domain.SetResponse(req, val)
// 	return nil
// }

// // PlatformLink 跳转项目
// func (api *userAdapter) PlatformLink(req *gate.Request) error {
// 	id, _ := strconv.Atoi(req.Args.Get("project_id"))
// 	typ, _ := strconv.Atoi(req.Args.Get("type"))
// 	redirect := req.Args.Get("redirect")
// 	userId := domain.GetUserID(req.Ctx)

// 	urlStr := fmt.Sprintf("/api/v1/gwapi/public/user/platform_link?user_id=%d&id=%d&type=%d&redirect=%s", userId, id, typ, redirect)
// 	val, err := utils.GwRawHttpGet(req.Ctx, config.Data().OpenCompany.InnerDomain, urlStr, domain.GetLanguage(req.Ctx))
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err)
// 		return err
// 	}

// 	domain.SetResponse(req, val)
// 	return nil
// }

// // RefreshToken
// func (api *userAdapter) RefreshToken(req *gate.Request) error {
// 	type Refresh struct {
// 		RefreshToken string `json:"refresh_token"`
// 	}
// 	var param Refresh
// 	err := rxkit.UnmarshalJSON(req.GetRawRequest().GetBody(), &param)
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err)
// 		return err
// 	}
// 	accessToken, refreshToken, err := api.user.RefreshToken(req.Ctx, param.RefreshToken)
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err)
// 		return err
// 	}
// 	req.Response().SetData(map[string]string{
// 		"access_token":  accessToken,
// 		"refresh_token": refreshToken,
// 	})
// 	return nil
// }

// // Login 用户登录
// func (api *userAdapter) Login(req *gate.Request) (err error) {
// 	var param domain.LoginParam
// 	err = json.Unmarshal(req.GetRawRequest().Body, &param)
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err, param)
// 		return err
// 	}

// 	cc, _ := strconv.Atoi(param.Cc)
// 	pc, _ := strconv.Atoi(param.Pc)
// 	userID, _ := strconv.Atoi(param.UserID)
// 	companyId, _ := strconv.Atoi(param.CompanyId)

// 	ts, _ := strconv.Atoi(param.Ts)
// 	if ts == 0 || int64(ts) > time.Now().Unix() || time.Now().Unix()-int64(ts) > 120 {
// 		if param.TraceID != "8e5b49683504eb6b2c38603d4d420cf6" {
// 			ulog.Warn(req.Ctx, domain.ErrSignExpired)
// 			return domain.ErrSignExpired
// 		}
// 	}
// 	if rxkit.GetCPSign(param.TraceID+"_"+param.UserID+"_"+param.CompanyId, param.Ts,
// 		config.Data().OpenCompany.SecretVal) != param.Sign {
// 		err := fmt.Errorf("签名错误")
// 		ulog.Warn(req.Ctx, err)
// 		return domain.ErrSignError
// 	}

// 	if userID < 1 || companyId < 1 {
// 		ulog.Warn(req.Ctx, "参数错误")
// 		return domain.ErrQueryInvalid
// 	}

// 	var strictParam = domain.LoginParamStrict{
// 		Sign:               param.Sign,
// 		Ts:                 int64(ts),
// 		TraceID:            param.TraceID,
// 		UserID:             int64(userID),
// 		IsModifiedPassword: param.IsModifiedPassword,
// 		Cc:                 int64(cc),
// 		Pc:                 int64(pc),
// 		CompanyId:          int64(companyId),
// 		Usertype:           param.Usertype,
// 	}

// 	user, companyInfo, err := api.user.Login(req.Ctx, &strictParam)
// 	if err != nil {
// 		ulog.Warn(req.Ctx, "登录失败, err: %v, params: %+v", err, strictParam)
// 		return err
// 	}
// 	_ = companyInfo
// 	val, _ := rxkit.MarshalJSON(user)
// 	var m = map[string]any{}
// 	rxkit.UnmarshalJSON(val, &m)

// 	req.Response().SetData(m)
// 	return nil
// }

// // Logout 退出登录
// func (api *userAdapter) Logout(req *gate.Request) (err error) {
// 	if config.GlobalTokenUse {
// 		var param domain.Token
// 		err = rxkit.UnmarshalJSON(req.GetRawRequest().GetBody(), &param)
// 		if err != nil || param.RefreshToken == "" {
// 			ulog.Warn(req.Ctx, err)
// 			return nil
// 		}
// 		err = api.user.Logout(req.Ctx, req, &param)
// 		if err != nil {
// 			ulog.Warn(req.Ctx, err)
// 			return err
// 		}
// 		return nil
// 	}

// 	req.GlobalSessions.Delete("userid")
// 	req.GlobalSessions.Delete("company_id")
// 	req.GlobalSessions.Delete("private_domain_id")
// 	req.GlobalSessions.Delete("username")
// 	req.GlobalSessions.Delete("name")
// 	req.GlobalSessions.Delete("email")
// 	req.GlobalSessions.Delete("company_name")
// 	req.GlobalSessions.Save()
// 	return nil
// }

// // PermissionGetAll
// func (api *userAdapter) PermissionGetAll(req *gate.Request) (err error) {
// 	data, err := api.user.GetPermissionAll(req.Ctx,
// 		domain.GetUserID(req.Ctx), domain.GetUsertype(req.Ctx))
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err)
// 		return err
// 	}

// 	req.Response().SetData(data)
// 	return nil
// }

// // ResetPassword 重置密码
// func (api *userAdapter) ResetPassword(req *gate.Request) (err error) {
// 	var param domain.ResetPassword
// 	err = json.Unmarshal(req.GetRawRequest().Body, &param)
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err, param)
// 		return err
// 	}

// 	if param.ID != int(req.RawRequest.GetUserID()) {
// 		ulog.Warn(req.Ctx, err, param)
// 		return domain.ErrQueryInvalid
// 	}

// 	err = api.user.ResetPassword(req.Ctx, domain.GetCompanyID(req.Ctx), &param)
// 	if err != nil {
// 		ulog.Warn(req.Ctx, err)
// 		return err
// 	}

// 	return nil
// }
