package main

import "fmt"

type epson struct{}

func (e *epson) PrintFile() {
	fmt.Println("This is from printer: epson")
}
