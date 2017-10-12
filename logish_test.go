package logish

import (
	"encoding/hex"
	"fmt"
)

func ExampleLogf() {
	l := Logger{}
	defer l.Flush()
	l.Logf("Hello %d", 1)
	l.Logf("Hello %d", 2)
	fmt.Println("Println call")
	// Output:
	// Println call
	// Hello 1
	// Hello 2
}

func ExampleField() {
	l := Logger{}
	defer l.Flush()
	l.Field("field1", 1)
	l.Field("field2withLongName", 1.2)
	l.Field("field3", "abc")
	// Output:
	// field1:             1
	// field2withLongName: 1.2
	// field3:             "abc"
}

func ExampleFieldF() {
	l := Logger{}
	defer l.Flush()
	l.Fieldf("field1", "%2.3f", 1.2)
	l.Fieldf("field2", "%#[1]x %[1]d", 123)
	// Output:
	// field1: 1.200
	// field2: 0x7b 123
}

func ExampleClear() {
	l := Logger{}
	defer l.Flush()
	l.Logf("log1")
	l.Field("field1", 1)
	l.Fieldf("field2", "%x", 2)
	l.Clear()
	l.Logf("log2")
	// Output:
	// log2
}

func ExampleNilLogger() {
	var l *Logger
	defer l.Flush()
	l.Field("field1", 1)
	l.Fieldf("field1", "%d", 1)
	l.Logf("log1")
	l.Clear()
	// Output:
}

func ExampleWithHeader() {
	l := Logger{Header: "Head"}
	defer l.Flush()
	l.Field("field1", 1)
	// Output:
	// == Head ==
	// field1: 1
}

func ExampleOneline() {
	l := Logger{Header: "Head"}
	defer l.FlushOneline()
	l.Logf("log1")
	l.Fieldf("field1", "%2.3f", 1.2)
	l.Fieldf("field2", "%#[1]x(%[1]d)", 123)
	l.Fieldf("field3", "", "nofmt")
	// Output:
	// Head['log1' field1:1.200 field2:0x7b(123) field3:"nofmt"]
}

func ExampleNilOnelineLogger() {
	var l *Logger
	defer l.FlushOneline()
	l.Field("field1", 1)
	l.Logf("log1")
	l.Clear()
	// Output:
}

// Complex types are stored by reference, so will show the value at defer time.
func ExampleBugga() {
	l := Logger{}
	defer l.FlushOneline()
	x := []int{1}
	l.Fieldf("x", "%d", x)
	x[0] = 2
	// Output: x:[2]
}

// Realistic example
func ExampleTSDecode() {
	m := map[string]uint16{}
	l := &Logger{}
	defer l.Flush()
	ts, _ := hex.DecodeString("47410030075000007b0c7e00000001")

	field := func(l *Logger, m map[string]uint16, name string, val uint16) {
		l.Field(name, val)
		if m != nil {
			m[name] = val
		}
	}
	field(l, m, "SyncByte", uint16(ts[0]))
	field(l, m, "TransportErrorIndicator", uint16(ts[1]>>7&1))
	field(l, m, "PayloadUnitStartIndicator", uint16(ts[1]>>6&1))
	field(l, m, "TransportPriority", uint16(ts[1]>>5&1))
	field(l, m, "PID", uint16(ts[1]&0x1F)<<8|uint16(ts[2]))
	field(l, m, "TransportScramblingControl", uint16(ts[3]>>2&3))
	field(l, m, "AdaptationFieldControl", uint16(ts[3]>>4&3))
	field(l, m, "ContainsAdaptationField", uint16(ts[3]>>5&1))
	field(l, m, "ContainsPayload", uint16(ts[3]>>4&1))
	field(l, m, "ContinuityCounter", uint16(ts[3]&0xf))
	l.Logf("PID=%d", m["PID"])

	// Output:
	// PID=256
	// SyncByte:                   0x47
	// TransportErrorIndicator:    0x0
	// PayloadUnitStartIndicator:  0x1
	// TransportPriority:          0x0
	// PID:                        0x100
	// TransportScramblingControl: 0x0
	// AdaptationFieldControl:     0x3
	// ContainsAdaptationField:    0x1
	// ContainsPayload:            0x1
	// ContinuityCounter:          0x0
}
