package main

import "fmt"

func display(start *node) {
	if start == nil {
		fmt.Println("Empty linked lists")
	} else {
		fmt.Println("Displaying linked lists")
		for i := start; i != nil; i = i.next {
			fmt.Printf("%d ", i.val)
		}
	}
}
