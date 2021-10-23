package main

import "fmt"

func addNodeToBeginning(start *node, value int) *node {
	newNode := &node{
		prev: nil,
		val:  value,
		next: nil,
	}
	(*newNode).next = start
	start = newNode
	return start
}

func addNodeToEnd(start *node, value int) *node {
	newNode := &node{
		prev: nil,
		val:  value,
		next: nil,
	}
	if start != nil {
		ptr := start
		for {
			if ptr.next == nil {
				break
			}
			ptr = ptr.next
		}
		ptr.next = newNode
		newNode.prev = ptr
	} else {
		start = newNode
	}
	return start
}

func addNodeBeforeValue(start *node, value int, nodeval int) (*node, error) {
	newNode := &node{
		prev: nil,
		val:  value,
		next: nil,
	}
	ptr := start
	for {
		if ptr == nil {
			break
		}
		if ptr.val == nodeval {
			if ptr.prev != nil {
				ptr.prev.next = newNode
				newNode.prev = ptr.prev
			} else {
				start = newNode
			}
			newNode.next = ptr
			ptr.prev = newNode
			return start, nil
		}
		ptr = ptr.next

	}
	return start, fmt.Errorf("ValueNotFound: Not able to find the value %d in the list", nodeval)
}

func addNodeAfterValue(start *node, value int, nodeval int) (*node, error) {
	newNode := &node{
		prev: nil,
		val:  value,
		next: nil,
	}
	ptr := start
	for {
		if ptr == nil {
			break
		}
		if ptr.val == nodeval {
			if ptr.next != nil {
				(*newNode).next = ptr.next
				ptr.next.prev = newNode
			}
			(*newNode).prev = ptr
			ptr.next = newNode
			return start, nil
		}
		ptr = ptr.next
	}
	return start, fmt.Errorf("ValueNotFound: Not able to find the value %d in the list", nodeval)
}
