package print

import "fmt"

// hp is a pointer
type hp struct{}

func (h *hp) printFile() string {
	fmt.Println("printing by hp pointer")
	return "printing by hp pointer"
}
