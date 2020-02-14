package log

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func logHandler(prefix, s string, a ...interface{}) {
	t := time.Now()
	timeString := color.HiBlackString("[%02d:%02d:%02d]", t.Hour(), t.Minute(), t.Second())
	fmt.Printf("%s %s %s\n", timeString, prefix, fmt.Sprintf(s, a...))
}

// Error is used to log error messages
func Error(s string, a ...interface{}) {
	logHandler(color.RedString("Error:"), s, a...)
}

// ErrorFatal is used to log error messages and exit
func ErrorFatal(s string, a ...interface{}) {
	Error(s, a...)
	os.Exit(1)
}

// Success is used to log success messages
func Success(s string, a ...interface{}) {
	logHandler(color.GreenString("Success:"), s, a...)
}

// Info is used to log info messages
func Info(s string, a ...interface{}) {
	logHandler(color.CyanString("Info:"), s, a...)
}

// Warning is used to log warning messages
func Warning(s string, a ...interface{}) {
	logHandler(color.YellowString("Warning:"), s, a...)
}

// Custom is used to log messages with a custom prefix
func Custom(prefix string, colourFunc func(s string, a ...interface{}) string, s string, a ...interface{}) {
	logHandler(colourFunc(prefix), s, a...)
}
