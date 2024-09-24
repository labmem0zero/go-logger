package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type logger interface {
	Write(logLevel string, i ...interface{})
}

type Logger struct {
	loggers []logger
}

func New(loggers ...logger) Logger {
	return Logger{loggers: loggers}
}

func (log Logger) Debug(i ...interface{}) {
	i = append([]interface{}{getFuncName()}, i...)
	for _, l := range log.loggers {
		l.Write(debug, i...)
	}
}

func (log Logger) Info(i ...interface{}) {
	i = append([]interface{}{getFuncName()}, i...)
	for _, l := range log.loggers {
		l.Write(info, i...)
	}
}

func (log Logger) Warn(i ...interface{}) {
	i = append([]interface{}{getFuncName()}, i...)
	for _, l := range log.loggers {
		l.Write(warning, i...)
	}
}

func (log Logger) Error(i ...interface{}) {
	i = append([]interface{}{getFuncName()}, i...)
	for _, l := range log.loggers {
		l.Write(err, i...)
	}
}

func (log Logger) Fatal(i ...interface{}) {
	i = append([]interface{}{getFuncName()}, i...)
	for _, l := range log.loggers {
		l.Write(fatal, i...)
	}
	os.Exit(1)
}

func getFuncName() string {
	var buffer bytes.Buffer
	pc := make([]uintptr, 10)
	runtime.Callers(4, pc)
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
