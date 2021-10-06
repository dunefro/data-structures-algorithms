package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Println("Searching for the keyword", keyword, "in file", f.name)
}

func (f *file) getName() string {
	return f.name
}
