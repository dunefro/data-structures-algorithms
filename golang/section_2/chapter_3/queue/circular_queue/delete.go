package main

import "fmt"

func deleteNode(queue *[]int, front, rear int) (*[]int, int, int, int, error) {
	var err error
	var value int
	if front == -1 {
		err = fmt.Errorf("QueueUnderflow: Queue is Empty")
		value = 0
	} else {
		fmt.Println(front)
		value = (*queue)[front]
		(*queue)[front] = 0
		if front == rear {
			front = -1
			rear = -1
		} else {
			if front == MAX-1 {
				front = 0
			} else {
				front++
			}
		}
	}
	return queue, front, rear, value, err
}
