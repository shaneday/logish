package logish

import (
	"fmt"
	"io"
	"os"
)

// A Simple logger provides minimal functionality above stdout
type Simple struct {
	currentTag string // !="" indicates a line is in progress
	dest       io.Writer
}

// Logf writes a log message fragment. Fragments with the same tag are appended
// to the line (space separated), while non-matching tags trigger a new line.
// Tags are printed at the start of each line, unless the first character of the
// tag is '-'. An empty tag forces the message on to a line by itself. If the
// format has a '\n' prefix then a new line is started even if the tag is
// unchanged, and a '\n' suffix causes a newline after this message.
func (o *Simple) Logf(tag, format string, a ...interface{}) {

	// Trigger newline before msg if format has '\n' prefix
	if len(format) > 0 && format[0] == '\n' {
		format = format[1:]
		fmt.Fprintf(o.dest, "\n")
		o.currentTag = ""
	}

	// Force newline after msg if format has \n suffix
	newlineSuffix := len(format) > 0 && format[len(format)-1] == '\n'
	if newlineSuffix {
		format = format[0 : len(format)-1]
	}

	// Easy case, append to existing line
	if tag != "" && tag == o.currentTag {
		fmt.Fprintf(o.dest, " "+format, a...)
		if newlineSuffix {
			fmt.Fprintf(o.dest, "\n")
			o.currentTag = ""
		}
		return
	}

	// End any existing line
	if o.currentTag != "" {
		fmt.Fprintf(o.dest, "\n")
	}

	// Prefix tag, unless disabled
	if len(tag) > 1 && tag[0] != '-' {
		fmt.Fprintf(o.dest, "%s: ", tag)
	}

	// Append the log message
	fmt.Fprintf(o.dest, format, a...)

	// Full line message, add newline
	if tag == "" || newlineSuffix {
		fmt.Fprintf(o.dest, "\n")
		tag = ""
	}

	// Record tag for next time
	o.currentTag = tag
}

// Default is a singleton simple logger used by logish.Logf()
var Default Simple

func init() {
	Default.dest = os.Stdout
}

// Logf writes to the default logger.
func Logf(tag, format string, a ...interface{}) {
	Default.Logf(tag, format, a...)
}
