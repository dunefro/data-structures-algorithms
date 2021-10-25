package main

import "fmt"

func peekNode(front *node) (int, error) {
	if front == nil {
		return 0, fmt.Errorf("EmptyQueue: Queue is empty")
	} else {
		return front.val, nil
	}
}
