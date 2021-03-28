package print

import "fmt"

// epson is a pointer
type epson struct {}

func (e *epson) printFile() string {
	fmt.Println("printing by epson printer")
	return "printing by epson printer"
}

