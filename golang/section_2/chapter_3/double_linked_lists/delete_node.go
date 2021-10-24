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
func deleteNodeFromGivenValue(start *node, nodeval int) (*node, error) {
	ptr := start
	for {
		if ptr == nil {
			break
		}
		if ptr.val == nodeval {
			if ptr.prev == nil {
				start = ptr.next
			} else {
				ptr.prev.next = ptr.next
			}
			if ptr.next != nil {
				ptr.next.prev = ptr.prev
			}
			return start, nil
		}
		ptr = ptr.next
	}
	return start, fmt.Errorf("ValueNotFound: Not able to find the value %d in the list", nodeval)
}
func deleteList(start *node) *node {
	return nil
}
