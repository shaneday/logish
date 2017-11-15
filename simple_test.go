package logish

import (
	"os"
)

func ExampleSimple() {
	Default.dest = os.Stdout // Update to go-test capture
	Logf("tag1", "msg1a")
	Logf("tag1", "msg1b")
	Logf("", "") // force new line
	// Output:
	// tag1: msg1a msg1b
}

func ExampleSimpleIntermix() {
	Default.dest = os.Stdout // Update to go-test capture
	Logf("tag1", "msg1a")
	Logf("tag2", "msg2a")
	Logf("tag1", "msg1b")
	Logf("", "") // force new line
	// Output:
	// tag1: msg1a
	// tag2: msg2a
	// tag1: msg1b
}

func ExampleSimpleTagless() {
	Default.dest = os.Stdout // Update to go-test capture
	Logf("tag1", "msg1a")
	Logf("-notag", "tagless-msg")
	Logf("tag1", "msg1b")
	Logf("", "") // force new line
	// Output:
	// tag1: msg1a
	// tagless-msg
	// tag1: msg1b
}

func ExampleSimpleFullLine() {
	Default.dest = os.Stdout // Update to go-test capture
	Logf("tag1", "msg1a")
	Logf("", "Forced full line message")
	Logf("tag1", "msg1b")
	Logf("", "") // force new line
	// Output:
	// tag1: msg1a
	// Forced full line message
	// tag1: msg1b
}
