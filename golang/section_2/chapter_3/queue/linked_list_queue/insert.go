package main

func insertNode(front *node, rear *node, value int) (*node, *node) {
	newNode := &node{
		val:  value,
		next: nil,
	}
	if front == nil {
		front = newNode
	} else {
		rear.next = newNode
	}
	rear = newNode
	return front, rear
}
