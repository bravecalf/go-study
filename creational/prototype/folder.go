package prototype

import "fmt"

// Folder 具体原型
type Folder struct {
	children []INode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, file := range f.children {
		file.print(indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	copyFolder := &Folder{name: f.name + "_clone"}
	tmpChildren := make([]INode, 0)
	for _, i := range f.children {
		tmpCopy := i.clone()
		tmpChildren = append(tmpChildren, tmpCopy)
	}
	copyFolder.children = tmpChildren
	return copyFolder
}
