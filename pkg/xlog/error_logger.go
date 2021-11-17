package xlog

func ErrorLogOut(format *string, v ...interface{}) {
	if format != nil {
		zap.Errorf(*format, v...)
		return
	}
	zap.Error(v...)
}
