package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type logger interface {
	Write(logLevel string, reqID string, i ...interface{})
}

type Logger struct {
	loggers []logger
}

func New(loggers ...logger) Logger {
	return Logger{loggers: loggers}
}

func (log Logger) Debug(reqID string, i ...interface{}) {
	i = append(i, []interface{}{getFuncName()}...)
	for _, l := range log.loggers {
		l.Write(debug, reqID, i...)
	}
}

func (log Logger) Info(reqID string, i ...interface{}) {
	i = append(i, []interface{}{getFuncName()}...)
	for _, l := range log.loggers {
		l.Write(info, reqID, i...)
	}
}

func (log Logger) Warn(reqID string, i ...interface{}) {
	i = append(i, []interface{}{getFuncName()}...)
	for _, l := range log.loggers {
		l.Write(warning, reqID, i...)
	}
}

func (log Logger) Error(reqID string, i ...interface{}) {
	i = append(i, []interface{}{getFuncName()}...)
	for _, l := range log.loggers {
		l.Write(err, reqID, i...)
	}
}

func (log Logger) Fatal(reqID string, i ...interface{}) {
	i = append(i, []interface{}{getFuncName()}...)
	for _, l := range log.loggers {
		l.Write(fatal, reqID, i...)
	}
	os.Exit(1)
}

func getFuncName() string {
	var buffer bytes.Buffer
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	frame, _ := runtime.CallersFrames(pc).Next()
	function := frame.Function
	line := frame.Line
	buffer.WriteString(function)
	buffer.WriteString(fmt.Sprintf(":%d", line))

	return filepath.Base(buffer.String())
}

const (
	debug   = "DEBUG"
	info    = "INFO"
	warning = "WARNING"
	err     = "ERROR"
	fatal   = "FATAL"
)
