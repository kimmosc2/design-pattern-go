package print

import "fmt"

type linux struct {
	printer
}

func (l *linux) print() string {
	fmt.Println("linux printing")
	return l.printer.printFile()
}

func (l *linux) setPrinter(p printer) {
	l.printer = p
}
