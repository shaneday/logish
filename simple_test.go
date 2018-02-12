package logish

import (
	"os"
)

func setup() (i int) {
	// Update dest to the temp redirect that 'go test' has set up
	Default.dest = os.Stdout
	return 0 // allow 'defer teardown(setup())' idiom
}
func teardown(int) {
	// Force closing any partial line in progress
	Logf("", "")
}

func ExampleSimple() {
	defer teardown(setup())
	Logf("TS[1]", "PID=%x", 0x100)
	Logf("TS[1]", "PUSI=%d", 1)
	// Output:
	// TS[1]: PID=100 PUSI=1
}

func ExampleSimpleIntermix() {
	defer teardown(setup())
	Logf("TS[1]", "PID=%x PUSI=%d", 0x100, 1)
	Logf(" PES", "SID=%x PTS=%.3f", 0xe0, 1.442)
	Logf("TS[1]", "Feeding video pipeline %d bytes", 157)
	// Output:
	// TS[1]: PID=100 PUSI=1
	//  PES: SID=e0 PTS=1.442
	// TS[1]: Feeding video pipeline 157 bytes
}

func ExampleSimpleTagless() {
	defer teardown(setup())
	Logf("TS[1]", "PID=%x", 0x100)
	Logf("-PES", "Found video PES:")
	Logf("-PES", "SID=%x PTS=%.3f", 0xe0, 1.442)
	Logf("TS[1]", "PUSI=%d", 1)
	// Output:
	// TS[1]: PID=100
	// Found video PES: SID=e0 PTS=1.442
	// TS[1]: PUSI=1
}

func ExampleSimpleFullLine() {
	defer teardown(setup())
	Logf("TS[1]", "PID=%x", 0x100)
	Logf("", "Warning: packet payload is encrypted")
	Logf("TS[1]", "PUSI=%d", 1)
	// Output:
	// TS[1]: PID=100
	// Warning: packet payload is encrypted
	// TS[1]: PUSI=1
}

func ExampleSimpleNewlines() {
	defer teardown(setup())
	Logf("TS[1]", "PID=%x", 0x100)
	Logf("TS[1]", "(Message ending a line)\n")
	Logf("TS[1]", "PUSI=%d", 1)
	Logf("TS[1]", "\n(Message on new line)")
	Logf("TS[1]", "AF=%d", 1)
	Logf("TS[1]", "\n(Message on line alone)\n")
	Logf("TS[1]", "PCR=%.3f", 0.7)
	// Output:
	// TS[1]: PID=100 (Message ending a line)
	// TS[1]: PUSI=1
	// TS[1]: (Message on new line) AF=1
	// TS[1]: (Message on line alone)
	// TS[1]: PCR=0.700
}

func ExampleDoubleNewlineBug() {
	defer teardown(setup())
	Logf("tag", "one\n")
	Logf("tag", "\ntwo")
	// Output:
	// tag: one
	// tag: two
}

func ExampleEmptyLogBug() {
	defer teardown(setup())
	Logf("tag", "one")
	// Bug was this triggering newlines before and after "" msg
	Logf("", "")
	Logf("tag", "two")
	// Output:
	// tag: one
	// tag: two
}

func ExampleNewlineSpaceBug() {
	defer teardown(setup())
	Logf("tag", "one")
	// Bug was this triggering extra space before "two" msg
	Logf("tag", "\n")
	Logf("tag", "two") // This was pre-padding with a space
	// Output:
	// tag: one
	// tag: two
}
