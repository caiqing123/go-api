package xlog

import (
	"api/di"
)

func InfoLogOut(format *string, v ...interface{}) {
	var zap = di.Zap()
	if format != nil {
		zap.Infof(*format, v...)
		return
	}
	zap.Info(v...)
}
