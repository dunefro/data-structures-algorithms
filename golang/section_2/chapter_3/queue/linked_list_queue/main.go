package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {
	var front *node
	var rear *node
	var opt int
	for {
		listOfOperations()
		fmt.Printf("Enter your option: ")
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1:
			var value int
			fmt.Printf("Enter the value of the node: ")
			fmt.Scanf("%d", &value)
			front, rear = insertNode(front, rear, value)
		case 2:
			var value int
			var err error
			front, rear, value, err = deleteNode(front, rear)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Deleted value from the queue is", value)
			}
		case 3:
			value, err := peekNode(front)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("The value in the front of the queue is", value)
			}
		case 4:
			displayQueue(front)
		case 5:
			return
		default:
			fmt.Printf("Unrecognized Option \"%d\"", opt)
		}

	}
}
