package main

import "fmt"

func deleteNode(front *node, rear *node) (*node, *node, int, error) {

	if front == nil {
		return front, rear, 0, fmt.Errorf("EmptyQueue: Queue is empty. Nothing to delete")
	} else {
		value := front.val
		front = front.next
		if front == nil {
			rear = nil
		}
		return front, rear, value, nil
	}
}
