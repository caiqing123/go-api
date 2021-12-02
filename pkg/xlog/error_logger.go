package xlog

import (
	"api/di"
)

func ErrorLogOut(format *string, v ...interface{}) {
	var zap = di.Zap()
	if format != nil {
		zap.Errorf(*format, v...)
		return
	}
	zap.Error(v...)
}
