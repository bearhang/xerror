package xerror

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	//ELongFile print the file name and its absolute path from root
	//such as C://golang/src/github.com/bearhang/xerror/xerror.go for windows
	//or /golang/src/github.com/bearhang/xerror/xerror.go for linux
	ELongFile = 1 << iota

	//EShortFile print only the file name without path
	EShortFile

	//ELongFunc print the function name and the relative path from GOPATH/src
	//eg:github.com/bearhang/xerror.function
	ELongFunc

	//EShortFunc print only the function name without path
	EShortFunc

	//EStdFlag defalt print style
	EStdFlag = EShortFile | EShortFunc
)

var gFlag = EStdFlag

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
		switch gFlag {
		case ELongFile:
			buf += fmt.Sprintf("%s:%d|%s\n",
				f.file, f.line, f.msg)
		case EShortFile:
			buf += fmt.Sprintf("%s:%d|%s\n",
				filepath.Base(f.file), f.line, f.msg)
		case ELongFunc:
			buf += fmt.Sprintf("%s:%d|%s\n",
				f.function, f.line, f.msg)
		case EShortFunc:
			buf += fmt.Sprintf("%s:%d|%s\n",
				filepath.Base(f.function), f.line, f.msg)
		case ELongFile | ELongFunc:
			buf += fmt.Sprintf("%s|%s:%d|%s\n",
				f.file, f.function, f.line, f.msg)
		case ELongFile | EShortFunc:
			buf += fmt.Sprintf("%s|%s:%d|%s\n",
				f.file, filepath.Base(f.function), f.line, f.msg)
		case EShortFile | ELongFunc:
			buf += fmt.Sprintf("%s|%s:%d|%s\n",
				filepath.Base(f.file), f.function, f.line, f.msg)
		case EShortFile | EShortFunc:
			buf += fmt.Sprintf("%s|%s:%d|%s\n",
				filepath.Base(f.file), filepath.Base(f.function), f.line, f.msg)
		default:
			panic(fmt.Sprintf("unsupport flag %x", gFlag))
		}
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

//SetFlags set print flag for Error
func SetFlags(flag int) error {
	if flag&ELongFile != 0 && flag&EShortFile != 0 {
		return New("confict flag ELongFile and EShortFile")
	}
	if flag&ELongFunc != 0 && flag&EShortFunc != 0 {
		return New("confict flag ELongFunc and EShortFunc")
	}
	gFlag = flag
	return nil
}
