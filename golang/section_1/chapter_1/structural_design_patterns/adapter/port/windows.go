package main

import "fmt"

type windows struct {
}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB port is plugged into windows machine")
}
