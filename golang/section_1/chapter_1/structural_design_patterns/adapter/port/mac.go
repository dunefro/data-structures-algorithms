package main

import "fmt"

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connection is plugged into MAC machine")
}
