package print

import "testing"

// TODO: pretty test case
func TestPrint(t *testing.T) {

	l1 := linux{}
	l1.setPrinter(&hp{})
	l1.print()

	l2 := linux{}
	l2.setPrinter(&epson{})
	l2.print()

	m1 := mac{}
	m1.setPrinter(&hp{})
	m1.print()

	m2 := mac{}
	m2.setPrinter(&epson{})
	m2.print()
}
