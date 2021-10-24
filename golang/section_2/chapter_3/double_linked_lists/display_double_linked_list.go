package main

import "fmt"

func display(start *node) {
	if start == nil {
		fmt.Println("Empty linked list")
		return
	}
	for ptr := start; ptr != nil; ptr = ptr.next {
		fmt.Printf("%d ", ptr.val)
	}
}
