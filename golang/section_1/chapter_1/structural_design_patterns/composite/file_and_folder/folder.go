package main

import "fmt"

type folder struct {
	component []component
	name      string
}

func (f *folder) search(keyword string) {
	fmt.Println("Searching Recursively in folder", f.name)
	for _, v := range f.component {
		v.search(keyword)
	}
}

// To add component to a folder
func (f *folder) add(c component) {
	f.component = append(f.component, c)
}
