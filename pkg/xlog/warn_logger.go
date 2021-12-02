package xlog

import "api/di"

func WarnLogOut(format *string, v ...interface{}) {
	var zap = di.Zap()
	if format != nil {
		zap.Warnf(*format, v...)
		return
	}
	zap.Warn(v...)
}
