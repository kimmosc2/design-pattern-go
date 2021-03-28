package file

import "fmt"

type folder struct {
	name      string
	component []component
}

func (f *folder) search(s string) (string, error) {
	for _, comp := range f.component {
		if result, err := comp.search(s); err == nil {
			return f.name + "/" + result, nil
		}
	}
	return "", fmt.Errorf("no such keyword in folder: %s", f.name)
}

func (f *folder) add(comp component) {
	f.component = append(f.component, comp)
}
