package ulog

// const (
// 	ApiLogTraceID = "api-log-trace-id"
// )

// var logger = xlog.GetLogger().WithOptions(zap.AddCallerSkip(1)).Sugar()

// func GetApiTraceId(ctx context.Context) string {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	return traceId
// }

// func GetLogger(ctx context.Context) *zap.SugaredLogger {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	if traceId == "" {
// 		traceId = uuid.NewString()
// 	}
// 	return xlog.GetSugaredLogger().With(ApiLogTraceID, traceId)
// }

// func Debugf(ctx context.Context, template string, args ...interface{}) {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	if traceId == "" {
// 		logger.Debugf(template, args...)
// 		return
// 	}
// 	logger.With(ApiLogTraceID, traceId).Debugf(template, args...)
// }

// func Infof(ctx context.Context, template string, args ...interface{}) {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	if traceId == "" {
// 		logger.Infof(template, args...)
// 		return
// 	}
// 	logger.With(ApiLogTraceID, traceId).Infof(template, args...)
// }

// func Warnf(ctx context.Context, template string, args ...interface{}) {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	if traceId == "" {
// 		logger.Warnf(template, args...)
// 		return
// 	}
// 	logger.With(ApiLogTraceID, traceId).Warnf(template, args...)
// }
// func Errorf(ctx context.Context, template string, args ...interface{}) {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	if traceId == "" {
// 		logger.Errorf(template, args...)
// 		return
// 	}
// 	logger.With(ApiLogTraceID, traceId).Errorf(template, args...)
// }

// func Warn(ctx context.Context, args ...interface{}) {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	if traceId == "" {
// 		logger.Warn(args...)
// 		return
// 	}
// 	logger.With(ApiLogTraceID, traceId).Warn(args...)
// }

// func Info(ctx context.Context, args ...interface{}) {
// 	traceId, _ := ctx.Value(ApiLogTraceID).(string)
// 	if traceId == "" {
// 		logger.Info(args...)
// 		return
// 	}
// 	logger.With(ApiLogTraceID, traceId).Info(args...)
// }
