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
	Logf("tag1", "msg1a=%d", 123)
	Logf("tag1", "msg1b=%.3f", 32.1)
	// Output:
	// tag1: msg1a=123 msg1b=32.100
}

func ExampleSimpleIntermix() {
	defer teardown(setup())
	Logf("tag1", "msg1a")
	Logf("tag2", "msg2a")
	Logf("tag1", "msg1b")
	// Output:
	// tag1: msg1a
	// tag2: msg2a
	// tag1: msg1b
}

func ExampleSimpleTagless() {
	defer teardown(setup())
	Logf("tag1", "msg1a")
	Logf("-notag", "tagless-msg")
	Logf("tag1", "msg1b")
	// Output:
	// tag1: msg1a
	// tagless-msg
	// tag1: msg1b
}

func ExampleSimpleFullLine() {
	defer teardown(setup())
	Logf("tag1", "msg1a")
	Logf("", "Forced full line message")
	Logf("tag1", "msg1b")
	// Output:
	// tag1: msg1a
	// Forced full line message
	// tag1: msg1b
}

func ExampleSimpleNewlines() {
	defer teardown(setup())
	Logf("tag1", "msg0")
	Logf("tag1", "msg1a-nl\n")
	Logf("tag1", "msg1b")
	Logf("tag1", "\nnl-msg1c")
	Logf("tag1", "msg1d")
	Logf("tag1", "\nnl-msg1e-nl\n")
	Logf("tag1", "msg1f")
	// Output:
	// tag1: msg0 msg1a-nl
	// tag1: msg1b
	// tag1: nl-msg1c msg1d
	// tag1: nl-msg1e-nl
	// tag1: msg1f
}
