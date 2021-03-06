package files

import (
	"fmt"
)

type folder struct {
	children []iNode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *folder) clone() iNode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempNode []iNode
	for _, child := range f.children {
		c := child.clone()
		tempNode = append(tempNode, c)
	}
	cloneFolder.children = tempNode
	return cloneFolder
}
