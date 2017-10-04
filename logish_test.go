package logish

import "fmt"

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
