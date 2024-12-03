package flogger

import (
	"fmt"
	"os"
	"time"

	"github.com/labmem0zero/go-logger/formats"
	"github.com/labmem0zero/go-logger/formats/formats_default"
	"github.com/labmem0zero/go-logger/impl"
)

type FLogger struct {
	out      *os.File
	offset   *int64
	format   formats.Format
	settings impl.LoggerSettings
}

func NewFileLogger(filename string, settings impl.LoggerSettings) (l FLogger, err error) {
	var f *os.File
	var offset int64
	if f, err = os.OpenFile(filename, os.O_RDWR, 0644); os.IsNotExist(err) {
		f, err = os.Create(filename)
		offset = 0
	} else {
		var stat os.FileInfo
		stat, err = f.Stat()
		if err != nil {
			return
		}
		offset = stat.Size()
	}
	if err != nil {
		return
	}
	l.out = f
	l.offset = &offset
	l.format = formats_default.DefaultPlain()
	l.settings = settings
	return
}

func (l FLogger) Write(level string, reqID string, v ...interface{}) {
	if l.settings.Levels != nil {
		if _, ok := l.settings.Levels[level]; !ok {
			return
		}
	}
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.0000"), l.settings.Environment, level, l.settings.AppName, l.settings.AppID, reqID}, v...)
	out := []byte(l.format.String(v...))
	if b, err := l.out.WriteAt(out, *l.offset); err != nil {
		fmt.Println(err)
	} else {
		*l.offset += int64(b)
	}
}
