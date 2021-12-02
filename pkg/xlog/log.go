package xlog

import (
	"fmt"
	"runtime"
)

func Info(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf(" file:	%s:%d", file, line)
	InfoLogOut(nil, append(args, files)...)
}

func Infof(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	format = fmt.Sprintf("%s file: %s:%d", format, file, line)
	InfoLogOut(&format, args...)
}

func Debug(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf(" file: %s:%d", file, line)
	DebugLogOut(nil, append(args, files)...)
}

func Debugf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	format = fmt.Sprintf("%s file: %s:%d", format, file, line)
	DebugLogOut(&format, args...)
}

func Warn(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf(" file: %s:%d", file, line)
	WarnLogOut(nil, append(args, files)...)
}

func Warnf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	format = fmt.Sprintf("%s file: %s:%d", format, file, line)
	WarnLogOut(&format, args...)
}

func Error(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf(" file: %s:%d", file, line)
	ErrorLogOut(nil, append(args, files)...)
}

func Errorf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	format = fmt.Sprintf("%s file: %s:%d", format, file, line)
	ErrorLogOut(&format, args...)
}
