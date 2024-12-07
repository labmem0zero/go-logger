package logger

import (
	"log"
	"testing"

	"github.com/labmem0zero/go-logger/impl"
	"github.com/labmem0zero/go-logger/impl/stdlogger"
)

func TestNewLogger(t *testing.T) {
	//fl, err := flogger.NewFileLogger("log.log", impl.LoggerSettings{
	//	AppName:     "AppNameTest",
	//	AppID:       "AppIDTest",
	//	Environment: "EnvTest",
	//})
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//tl, err := tlogger.NewLogger("", 0, impl.LoggerSettings{
	//	AppName:     "AppNameTest",
	//	AppID:       "AppIDTest",
	//	Environment: "EnvTest",
	//	Levels: map[string]struct{}{
	//		LevelErr: {},
	//	},
	//})
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	sl, err := stdlogger.NewStdLogger(impl.LoggerSettings{
		AppName:     "AppNameTest",
		AppID:       "AppIDTest",
		Environment: "EnvTest"})
	if err != nil {
		log.Println(err)
		return
	}
	l := New(sl)
	l.Debug("reqID", "Hello, world!")
	l.Debug("reqID", "Goodbye, world!")
}
