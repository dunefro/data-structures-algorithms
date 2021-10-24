package main

import "fmt"

type node struct {
	prev *node
	val  int
	next *node
}

func main() {
	var exit bool
	var start *node
	for {
		if exit {
			break
		}
		print_list()
		var opt int
		fmt.Printf("Enter you option: ")
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1:
			display(start)
		case 2:
			fmt.Printf("Enter the value for the node: ")
			var value int
			fmt.Scanf("%d", &value)
			start = addNodeToBeginning(start, value)
		case 3:
			fmt.Printf("Enter the value you wish to add at the end: ")
			var value int
			fmt.Scanf("%d", &value)
			start = addNodeToEnd(start, value)
		case 4:
			var nodeval int
			var value int
			var err error
			fmt.Printf("Enter the value for the node: ")
			fmt.Scanf("%d", &value)
			fmt.Printf("Enter the value before which you want your node to be added: ")
			fmt.Scanf("%d", &nodeval)
			start, err = addNodeBeforeValue(start, value, nodeval)
			if err != nil {
				fmt.Println(err.Error())
			}
		case 5:
			var nodeval int
			var value int
			var err error
			fmt.Printf("Enter the value for the node: ")
			fmt.Scanf("%d", &value)
			fmt.Printf("Enter the value after which you want your node to be added: ")
			fmt.Scanf("%d", &nodeval)
			start, err = addNodeAfterValue(start, value, nodeval)
			if err != nil {
				fmt.Println(err.Error())
			}
		case 6:
			var err error
			start, err = deleteNodeFromBeginning(start)
			if err != nil {
				fmt.Print(err.Error())
			}
		case 7:
			var err error
			start, err = deleteNodeFromEnd(start)
			if err != nil {
				fmt.Print(err.Error())
			}
		case 8:
			var err error
			var value int
			fmt.Printf("Enter the value you want to delete: ")
			fmt.Scanf("%d", &value)
			start, err = deleteNodeFromGivenValue(start, value)
			if err != nil {
				fmt.Print(err.Error())
			}
		case 9:
			start = deleteList(start)
		case 10:
			start = sortList(start)
		case 11:
			exit = true
		default:
			fmt.Println("Unrecognized option", opt)
		}
	}

}
