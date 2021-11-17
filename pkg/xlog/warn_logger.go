package xlog

func WarnLogOut(format *string, v ...interface{}) {
	if format != nil {
		zap.Warnf(*format, v...)
		return
	}
	zap.Warn(v...)
}
