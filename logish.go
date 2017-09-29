package logish

import "fmt"

type fieldItem struct {
	label   string
	vformat string
	value   interface{}
}

// A Logger records log entries until they are printed or discarded
type Logger struct {
	messages []string
	fields   []fieldItem
	Header   string
}

// Logf adds a basic printf style log entry
func (l *Logger) Logf(format string, a ...interface{}) {
	if l == nil {
		return
	}
	str := fmt.Sprintf(format, a...)
	l.messages = append(l.messages, str)
}

// Field records a field/value entry
func (l *Logger) Field(label string, a interface{}) {
	if l == nil {
		return
	}
	e := fieldItem{label: label, value: a}
	l.fields = append(l.fields, e)
}

// Fieldf records a field/value entry, with printf style formatting code
func (l *Logger) Fieldf(label, vformat string, a interface{}) {
	if l == nil {
		return
	}
	e := fieldItem{label: label, vformat: vformat, value: a}
	l.fields = append(l.fields, e)
}

// Clear discards all entried. Call when everything is good and logging is not required
func (l *Logger) Clear() {
	if l == nil {
		return
	}
	l.messages = []string{}
	l.fields = []fieldItem{}
}

// Exit is to be called by defer, it prints the recorded entries
func (l *Logger) Exit() {
	if l == nil {
		return
	}
	if l.Header != "" {
		fmt.Println("== " + l.Header + " ==")
	}
	for _, str := range l.messages {
		fmt.Println(str)
	}
	maxWidth := 0
	for _, e := range l.fields {
		if len(e.label) > maxWidth {
			maxWidth = len(e.label)
		}
	}
	for _, e := range l.fields {
		if e.vformat == "" {
			e.vformat = "%#v"
		}
		val := fmt.Sprintf(e.vformat, e.value)
		fmt.Printf("%-*s %s\n", maxWidth+1, e.label+":", val)
	}
}

// ExitOneline is like Exit, but prints a compact format
func (l *Logger) ExitOneline() {
	if l == nil {
		return
	}
	if l.Header != "" {
		fmt.Printf("%s[", l.Header)
	}
	for i, str := range l.messages {
		fmt.Printf("'%s'", str)
		if i < len(l.messages) || l.Header == "" {
			fmt.Printf(" ")
		}
	}
	for i, e := range l.fields {
		if e.vformat == "" {
			e.vformat = "%#v"
		}
		val := fmt.Sprintf(e.vformat, e.value)
		fmt.Printf("%s:%s", e.label, val)
		if i < len(l.fields)-1 || l.Header == "" {
			fmt.Printf(" ")
		}
	}
	if l.Header != "" {
		fmt.Printf("]")
	}
}
