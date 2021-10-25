package main

import "fmt"

func listOfOperations() {
	fmt.Println()
	fmt.Println("*** MENU ***")
	fmt.Println("1. Insert")
	fmt.Println("2. Delete")
	fmt.Println("3. Peek")
	fmt.Println("4. Display")
	fmt.Println("5. Exit")
}

func peek(queue *[]int, front int) (int, error) {
	var err error
	var value int
	if front == -1 {
		err = fmt.Errorf("EmptyQueue: Queue is empty")
	} else {
		value = (*queue)[front]
	}
	return value, err
}

func displayQueue(queue *[]int, front int, rear int) error {
	var err error
	if front == -1 {
		err = fmt.Errorf("EmptyQueue: Queue is empty")
	} else {
		if rear >= front {
			for i := front; i <= rear; i++ {
				fmt.Printf("%d ", (*queue)[i])
			}
		} else {
			for i := front; i < MAX; i++ {
				fmt.Printf("%d ", (*queue)[i])
			}
			for i := 0; i <= rear; i++ {
				fmt.Printf("%d ", (*queue)[i])
			}
		}
	}
	return err
}
