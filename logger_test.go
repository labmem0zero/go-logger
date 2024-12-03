package logger

import (
	"log"
	"testing"

	"github.com/labmem0zero/go-logger/impl"
	"github.com/labmem0zero/go-logger/impl/flogger"
	"github.com/labmem0zero/go-logger/impl/tlogger"
)

func TestNewLogger(t *testing.T) {
	fl, err := flogger.NewFileLogger("log.log", impl.LoggerSettings{
		AppName:     "AppNameTest",
		AppID:       "AppIDTest",
		Environment: "EnvTest",
	})
	if err != nil {
		log.Println(err)
		return
	}
	tl, err := tlogger.NewLogger("", 0, impl.LoggerSettings{
		AppName:     "AppNameTest",
		AppID:       "AppIDTest",
		Environment: "EnvTest",
		Levels: map[string]struct{}{
			levelErr: {},
		},
	})
	l := New(fl, tl)
	l.Debug("reqID", "Hello, world!")
}
