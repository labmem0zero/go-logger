package formats_default

import (
	"fmt"
	"time"

	"github.com/labmem0zero/go-logger/formats"
)

type defaultHTML struct {
	f string
}

// DefaultHTML - format:
// `
// <b>Timestamp:</b><p>%v</p>
// <b>Environment</b><p>%v</p>
// <b>LogLevel</b><p>%v</p>
// <b>AppName</b><p>%v</p>
// <b>AppID</b><p>%v</p>
// <b>ReqID</b><p>%v</p>
// <b>FuncName: </b>%v
// <b>Additional</b><p>%v</p>
// `
func DefaultHTML() formats.Format {
	return defaultHTML{f: `<b>Timestamp: </b>%v
<b>Environment: </b>%v
<b>LogLevel: </b>%v
<b>AppName: </b>%v
<b>AppID: </b>%v
<b>ReqID: </b>%v
<b>FuncName: </b>%v
<b>Additional: </b>%v`}
}

// String(args ...interface{}) string
// passing args: timestamp, environment, log_level, app_name, app_id, req_id, func_name, args...
func (f defaultHTML) String(args ...interface{}) string {
	if len(args) < 7 {
		args = append([]interface{}{time.Now(), "unidentified", "unidentified", "unidentified", "unidentified", "unidentified", "unidentified"}, args)
		return fmt.Sprintf(f.f, args...)
	}
	var tail string
	if len(args) > 7 {
		tail = fmt.Sprintln(args[7:]...)
	} else {
		tail = "none"
	}
	return fmt.Sprintf(f.f, args[0], args[1], args[2], args[3], args[4], args[5], args[6], tail)
}

// Byte (args ...interface{}) []byte
// passing args: timestamp, environment, app_name, app_id, log_level, req_id, func_name, args...
func (f defaultHTML) Byte(args ...interface{}) []byte {
	return []byte(f.String(args...))
}
