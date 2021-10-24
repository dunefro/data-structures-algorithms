package main

import "fmt"

func deleteNodeFromBeginning(start *node) (*node, error) {
	if start != nil {
		start = start.next
		return start, nil
	}
	return start, fmt.Errorf("EmptyList: Not able to delete the node from the list")
}

func deleteNodeFromEnd(start *node) (*node, error) {
	if start == nil {
		return start, fmt.Errorf("EmptyList: Not able to delete the node from the list")
	} else {
		ptr := start
		for {
			if ptr.next == nil {
				if ptr.prev != nil {
					ptr.prev.next = nil
				} else {
					start = nil
				}
				return start, nil
			}
			ptr = ptr.next
		}
	}
}
