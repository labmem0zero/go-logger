package formats_default

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labmem0zero/go-logger/formats"
)

type defaultPlain struct {
	f string
}

// DefaultPlain - format:
// `%timestamp environment: %v log_level: %v app_name: %v app_id: %v req_id: %v  func_name: %v additional: %v`
func DefaultPlain() formats.Format {
	return defaultPlain{f: "%v environment: %v log_level: %v app_name: %v app_id: %v req_id: %v func_name%v additional: %v"}
}

// String
// passing args: timestamp, environment, log_level, app_name, app_id, req_id, func_name args...
func (f defaultPlain) String(args ...interface{}) string {
	if len(args) < 7 {
		return fmt.Sprintln(args...)
	}
	var tail string
	if len(args) > 7 {
		tail = fmt.Sprintln(args[5:]...)
	} else {
		tail = "none"
	}
	return fmt.Sprintf(f.f, args[0], args[1], args[2], args[3], args[4], args[5], args[6], tail)
}

func f(t time.Time, appName string) {}

// Byte
// passing args: timestamp, environment, log_level, app_name, app_id, req_id, func_name, args...
func (f defaultPlain) Byte(args ...interface{}) []byte {
	s := struct {
		LogTime        string `json:"log_time"`
		Environment    string `json:"environment"`
		LogLevel       string `json:"log_level"`
		AppName        string `json:"app_name"`
		AppID          string `json:"app_id"`
		ReqID          string `json:"req_id"`
		FuncName       string `json:"func_name"`
		AdditionalData string `json:"additional_data"`
	}{}
	if len(args) < 6 {
		s.LogTime = time.Now().String()
		s.AdditionalData = fmt.Sprintln(args)
	} else {
		s.LogTime = args[0].(string)
		s.Environment = args[1].(string)
		s.LogLevel = args[2].(string)
		s.AppName = args[3].(string)
		s.AppID = args[4].(string)
		s.ReqID = args[5].(string)
		s.FuncName = args[6].(string)
	}
	if len(args) > 6 {
		s.AdditionalData = fmt.Sprintln(args[6:]...)
	}
	b, err := json.Marshal(s)
	if err != nil {
		return []byte(f.String())
	}
	return b
}
