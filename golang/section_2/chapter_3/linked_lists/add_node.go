package main

import "fmt"

func addNodeToBeginning(start *node, value int) *node {
	newNode := &node{
		val:  value,
		next: nil,
	}
	if start != nil {
		tmp := start
		(*newNode).next = tmp
	}
	start = newNode
	return start
}

func addNodeToEnd(start *node, value int) *node {
	newNode := &node{
		val:  value,
		next: nil,
	}
	ptr := start
	if ptr == nil {
		start = newNode
	} else {
		for {
			if ptr.next == nil {
				break
			}
			ptr = ptr.next
		}
		ptr.next = newNode
	}
	return start
}

func addNodeAfterValue(start *node, value int, nodeval int) (*node, error) {
	newNode := &node{
		val:  value,
		next: nil,
	}
	ptr := start
	for {
		if ptr == nil {
			break
		}
		if nodeval == ptr.val {
			(*newNode).next = ptr.next
			ptr.next = newNode
			return start, nil
		}
		ptr = ptr.next
	}
	return start, fmt.Errorf("UnknownValue: Value not found.")
}

func addNodeBeforeValue(start *node, value int, nodeval int) (*node, error) {
	newNode := &node{
		val:  value,
		next: nil,
	}
	var pptr *node
	ptr := start
	for {
		if ptr == nil {
			break
		}
		if ptr.val == nodeval {
			(*newNode).next = ptr
			if pptr != nil {
				pptr.next = newNode
			} else {
				start = newNode
			}
			return start, nil
		}
		pptr = ptr
		ptr = ptr.next
	}
	return start, fmt.Errorf("UnknownValue: Value not found.")
}
