package filelogger

import (
	"fmt"
	"os"
	"time"
)

type FLogger struct {
	out    *os.File
	offset *int64
}

func NewFileLogger(filename string) (l FLogger, err error) {
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
	return FLogger{out: f, offset: &offset}, err
}

func (l FLogger) Write(level string, v ...interface{}) {
	fmt.Println("Starting: ", *l.offset)
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.0000"), level}, v...)
	out := []byte(fmt.Sprintln(v...))
	if b, err := l.out.WriteAt(out, *l.offset); err != nil {
		fmt.Println(err)
	} else {
		*l.offset += int64(b)
		fmt.Println("After: ", *l.offset)
	}
}
