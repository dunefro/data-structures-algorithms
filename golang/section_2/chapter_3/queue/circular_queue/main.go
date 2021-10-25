package main

import "fmt"

const MAX = 10

var queue = make([]int, MAX)
var front = -1
var rear = -1

func main() {
	qptr := &queue
	for {
		listOfOperations()
		var opt int
		fmt.Printf("Enter your option: ")
		fmt.Scanf("%d ", &opt)
		switch opt {
		case 1:
			var value int
			var err error
			fmt.Printf("Enter the value of node: ")
			fmt.Scanf("%d", &value)
			qptr, front, rear, err = insertNode(qptr, front, rear, value)
			if err != nil {
				fmt.Println(err.Error())
			}
		case 2:
			var value int
			var err error
			qptr, front, rear, value, err = deleteNode(qptr, front, rear)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("The value from the queue is \"%d\"", value)
			}
		case 3:
			val, err := peek(qptr, front)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("The value in the front of the queue is \"%d\"", val)
			}
		case 4:
			err := displayQueue(qptr, front, rear)
			if err != nil {
				fmt.Println(err.Error())
			}
		case 5:
			return
		default:
			fmt.Printf("Unrecognized Option: \"%d\"\n", opt)
		}
	}
}
