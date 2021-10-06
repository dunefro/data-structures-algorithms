package main

import "fmt"

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request is from mac")
	m.printer.PrintFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}
