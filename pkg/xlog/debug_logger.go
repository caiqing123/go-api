package xlog

import "api/di"

func DebugLogOut(format *string, v ...interface{}) {
	var zap = di.Zap()
	if format != nil {
		zap.Debugf(*format, v...)
		return
	}
	zap.Debug(v...)
}
