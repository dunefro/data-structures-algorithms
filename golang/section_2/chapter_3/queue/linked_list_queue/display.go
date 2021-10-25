package main

import "fmt"

func displayQueue(front *node) {
	if front == nil {
		fmt.Println("Queue is empty")
		return
	}
	ptr := front
	for {
		if ptr == nil {
			break
		}
		fmt.Printf("%d ", ptr.val)
		ptr = ptr.next
	}
}
