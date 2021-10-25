package main

import "fmt"

func insertNode(queue *[]int, front, rear, value int) (*[]int, int, int, error) {
	var err error
	if front == -1 {
		front = 0
		rear = 0
	} else if (front == 0 && rear == MAX-1) || (rear == front-1) {
		err = fmt.Errorf("QueueOverflow: Queue is full.")
		// return queue, front, rear,
	} else {
		if rear == MAX-1 {
			rear = 0
		} else {
			rear += 1
		}
	}
	(*queue)[rear] = value
	return queue, front, rear, err
}
