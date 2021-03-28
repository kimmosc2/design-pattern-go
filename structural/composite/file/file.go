package file

import (
	"fmt"
	"strings"
)

type file struct {
	name    string
	content string
}

func (f *file) search(s string) (string, error) {
	fmt.Printf("Searching for keyword %s in file %s\n", s, f.name)
	if strings.Contains(f.content, s) {
		return f.name, nil
	}
	return "", fmt.Errorf("no such keywords in %s", f.name)
}
