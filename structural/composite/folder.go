package composite

import "fmt"

type Folder struct {
	name      string
	component []Component
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching for keyword [%s] in floder [%s].\n", keyword, f.name)
	for _, component := range f.component {
		component.search(keyword)
	}
}

func (f *Folder) addComponent(component Component) {
	f.component = append(f.component, component)
}
