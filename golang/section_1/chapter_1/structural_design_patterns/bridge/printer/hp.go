package main

import "fmt"

type hp struct{}

func (h *hp) PrintFile() {
	fmt.Println("This is from printer HP")
}
