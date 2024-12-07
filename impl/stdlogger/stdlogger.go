package stdlogger

import (
	"fmt"
	"strings"
	"time"

	"github.com/labmem0zero/go-logger/formats"
	"github.com/labmem0zero/go-logger/formats/formats_default"
	"github.com/labmem0zero/go-logger/impl"
)

type StdLogger struct {
	format   formats.Format
	settings impl.LoggerSettings
}

func NewStdLogger(settings impl.LoggerSettings) (l StdLogger, err error) {
	l.format = formats_default.DefaultPlain()
	l.settings = settings
	return
}

func (l StdLogger) Write(level string, reqID string, v ...interface{}) {
	if l.settings.Levels != nil {
		if _, ok := l.settings.Levels[level]; !ok {
			return
		}
	}
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.0000"), l.settings.Environment, level, l.settings.AppName, l.settings.AppID, reqID}, v...)
	fmt.Println(strings.TrimSuffix(l.format.String(v...), "\n"))

}
