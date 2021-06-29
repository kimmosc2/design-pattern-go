package print

import "fmt"

type mac struct {
	p printer
}

func (m *mac) print() string {
	fmt.Println("mac printing")
	return m.p.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.p = p
}


