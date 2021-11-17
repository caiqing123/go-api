package xlog

func InfoLogOut(format *string, v ...interface{}) {
	if format != nil {
		zap.Infof(*format, v...)
		return
	}
	zap.Info(v...)
}
