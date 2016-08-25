package xerror

import (
	"fmt"
	"runtime"
)

type frame struct {
	file     string
	function string
	line     int
	msg      string
}

type xerror struct {
	frames []frame
}

func (err *xerror) Error() string {
	var buf string

	for _, f := range err.frames {
		buf += fmt.Sprintf("%s|%s:%d|%s\n", f.file, f.function, f.line, f.msg)
	}

	return buf
}

func caller() (function string, file string, line int) {
	pc, file, line, _ := runtime.Caller(2)
	function = runtime.FuncForPC(pc).Name()
	return
}

//New create a new xerror
func New(a ...interface{}) error {
	var e *xerror
	var f frame
	f.function, f.file, f.line = caller()

	as := a
	e, ok := as[0].(*xerror)
	if ok {
		f.msg = fmt.Sprint(as[1:]...)
	} else {
		f.msg = fmt.Sprint(a...)
		e = new(xerror)
		e.frames = make([]frame, 0, 10)
	}

	e.frames = append(e.frames, f)
	return e
}
