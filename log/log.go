package log

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func logHandler(prefix, s string, a ...interface{}) {
	t := time.Now()
	_, fn, ln, _ := runtime.Caller(2)
	fn = path.Join(path.Base(path.Dir(fn)), path.Base(fn))
	timeString := color.HiBlackString("[%02d:%02d:%02d %s:%d]", t.Hour(), t.Minute(), t.Second(), fn, ln)
	fmt.Printf("%s %s: %s\n", timeString, prefix, fmt.Sprintf(s, a...))
}

// Error is used to log error messages
func Error(s string, a ...interface{}) {
	logHandler(color.RedString("Error"), s, a...)
}

// ErrorFatal is used to log error messages and exit
func ErrorFatal(s string, a ...interface{}) {
	logHandler(color.RedString("Error"), s, a...)
	os.Exit(1)
}

// Report is used to report an non-nil error
func Report(err error) {
	if err != nil {
		logHandler(color.RedString("Error"), err.Error())
	}
}

// ReportFatal is used to report an non-nil fatal error
func ReportFatal(err error) {
	if err != nil {
		logHandler(color.RedString("Error"), err.Error())
		os.Exit(1)
	}
}

// Success is used to log success messages
func Success(s string, a ...interface{}) {
	logHandler(color.GreenString("Success"), s, a...)
}

// Info is used to log info messages
func Info(s string, a ...interface{}) {
	logHandler(color.CyanString("Info"), s, a...)
}

// Warning is used to log warning messages
func Warning(s string, a ...interface{}) {
	logHandler(color.YellowString("Warning"), s, a...)
}

// Custom is used to log messages with a custom prefix
func Custom(prefix string, colourFunc func(s string, a ...interface{}) string, s string, a ...interface{}) {
	logHandler(colourFunc(prefix), s, a...)
}

// Header is used to print headers
func Header(s string, a ...interface{}) {
	fmt.Print(HeaderString(s, a...))
}

// HeaderString is used to return the string used to print headers
func HeaderString(s string, a ...interface{}) string {
	var underline string
	for range s {
		underline += "-"
	}
	return color.BlueString(s+"\n"+underline+"\n", a...)
}

// Description is used to print descriptions for items
func Description(item, s string, a ...interface{}) {
	fmt.Printf("%s %s\n", color.CyanString("- "+item+":"), color.WhiteString(fmt.Sprintf(s, a...)))
}

// Input is used to grab input via a custom prompt
func Input(destination *string, s string, a ...interface{}) {
	fmt.Print(color.MagentaString("  "+s+": ", a...))
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*destination = scanner.Text()
}
