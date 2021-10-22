package main

import "fmt"

func deleteNodeFromBeginning(start *node) (*node, error) {
	if start == nil {
		return start, fmt.Errorf("UnknownValue")
	}
	start = start.next
	return start, nil

}
func deleteNodeFromEnd(start *node) (*node, error) {
	ptr := start
	var pptr *node
	if start == nil {
		return start, fmt.Errorf("EmptyListException: ")
	} else {
		for {
			if ptr.next == nil {
				if pptr != nil {
					pptr.next = nil
				} else {
					start = nil
				}
				return start, nil
			}
			pptr = ptr
			ptr = ptr.next
		}
		return start, nil
	}
}

func deleteNodeFromGivenValue(start *node, value int) (*node, error) {
	var pptr *node
	ptr := start
	for {
		if ptr == nil {
			return start, fmt.Errorf("ValueNotFound: ")
		}
		if ptr.val == value {
			if pptr != nil {
				pptr.next = ptr.next
			} else {
				start = ptr.next
			}
			return start, nil
		}
		pptr = ptr
		ptr = ptr.next
	}
}

func deleteList(start *node) *node {
	start = nil
	return start
}
