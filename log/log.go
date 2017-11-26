package log

import (
	"fmt"
	"log"
	"time"
)

//color
const (
	RED = uint8(iota + 91)
	GREEN
	YELLOW
	BLUE
	MAGENTA

	info = "[INFO]"
	trac = "[TRAC]"
	erro = "[ERRO]"
	warn = "[WARN]"
	succ = "[SUCC]"
	fata = "[FATA]"
)

// Trace means panic
func Trace(format string, a ...interface{}) {
	prefix := yellow(trac)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

// Info menas some info
func Info(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

// Success menas success
func Success(format string, a ...interface{}) {
	prefix := green(succ)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

// Warning means some warnings
func Warning(format string, a ...interface{}) {
	prefix := magenta(warn)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

// Error means some errors
func Error(format string, a ...interface{}) {
	prefix := red(erro)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

// Fatal means exit
func Fatal(format string, a ...interface{}) {
	prefix := red(fata)
	log.Fatal(formatLog(prefix), fmt.Sprintf(format, a...))
}

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", RED, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", GREEN, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", YELLOW, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", BLUE, s)
}

func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", MAGENTA, s)
}

func formatLog(prefix string) string {
	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
}
