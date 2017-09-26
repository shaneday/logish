package logish

import "fmt"

type fieldItem struct {
	label   string
	vformat string
	value   interface{}
}
type Logger struct {
	messages []string
	fields   []fieldItem
	Header   string
}

func (l *Logger) Logf(format string, a ...interface{}) {
	if l == nil {
		return
	}
	str := fmt.Sprintf(format, a...)
	l.messages = append(l.messages, str)
}

func (l *Logger) Field(label string, a interface{}) {
	if l == nil {
		return
	}
	e := fieldItem{label: label, value: a}
	l.fields = append(l.fields, e)
}
func (l *Logger) Fieldf(label, vformat string, a interface{}) {
	if l == nil {
		return
	}
	e := fieldItem{label: label, vformat: vformat, value: a}
	l.fields = append(l.fields, e)
}

func (l *Logger) Clear() {
	if l == nil {
		return
	}
	l.messages = []string{}
	l.fields = []fieldItem{}
}

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
