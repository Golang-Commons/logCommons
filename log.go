package logCommons

import (
	"fmt"
	"io"
	"log"
	"runtime"
)

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags = Ldate | Ltime
)

const (
	TextBlack   = 30
	TextRed     = 31
	TextGreen   = 32
	TextYellow  = 33
	TextBlue    = 34
	TextMagenta = 35
	TextCyan    = 36
	TextWhite   = 37
)

const (
	DEBUG   = 0
	INFO    = 1
	WARNING = 2
	ERROR   = 3
	PANIC   = 4
	FATAL   = 5
)

var levelArray = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"PANIC",
	"FATAL",
}

var levelColorArray = []int{
	TextBlue,
	TextGreen,
	TextMagenta,
	TextRed,
	TextBlack,
	TextBlack,
}

var (
	Flevel     = 0
	Fcolor     = 0
	Ffuncname  = 0
	Flevelname = 0
)

const (
	Lcolor = 1 << iota
	Lfuncname
	Llevelname
)

func SetLevel(level int) {
	Flevel = level
}

func SetFlag(flag int) {
	if flag&Lcolor != 0 {
		Fcolor = 1
	}
	if flag&Lfuncname != 0 {
		Ffuncname = 1
	}
	if flag&Llevelname != 0 {
		Flevelname = 1
	}

}

func SetFlags(flag int) {
	log.SetFlags(flag)
}

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func println(level int, a ...interface{}) {
	if level < Flevel {
		return
	}
	str := fmt.Sprintf("%s", a...)

	funcname := ""
	if Ffuncname == 1 {
		pc, _, _, _ := runtime.Caller(2)
		funcname = runtime.FuncForPC(pc).Name()
		str = funcname + " " + str
	}

	if Fcolor == 1 && Flevelname == 1 {
		str = fmt.Sprintf("\x1b[0;%dm[%s] \x1b[0m%s", levelColorArray[level], levelArray[level], str)
	} else if Flevelname == 1 {
		str = levelArray[level] + " " + str
	} else if Fcolor == 1 {
		str = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", levelArray[level], str)
	}
	log.Output(3, fmt.Sprintln(a...))
}

func Debug(a ...interface{}) {
	println(DEBUG, a...)
}

func Info(a ...interface{}) {
	println(INFO, a...)
}

func Warn(a ...interface{}) {
	println(WARNING, a...)
}

func Error(a ...interface{}) {
	println(ERROR, a...)
}

func Panic(a ...interface{}) {
	println(PANIC, a...)
}

func Fatal(a ...interface{}) {
	println(FATAL, a...)
}

func Println(a ...interface{}) {
	log.Output(2, fmt.Sprintln(a...))
}
