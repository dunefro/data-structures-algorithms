package main

import "fmt"

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("This request is from windows")
	w.printer.PrintFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}
