package xlog

func DebugLogOut(format *string, v ...interface{}) {
	if format != nil {
		zap.Debugf(*format, v...)
		return
	}
	zap.Debug(v...)
}
