package composite

import "fmt"

// File 具体组件
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword [%s] in file [%s].\n", keyword, f.name)
}
