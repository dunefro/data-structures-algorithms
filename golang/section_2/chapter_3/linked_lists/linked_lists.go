package main

import "fmt"

type node struct {
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
			fmt.Printf("Enter the value you wish to add in the beginning: ")
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
			fmt.Printf("Enter the value before which you want you are node to be added: ")
			fmt.Scanf("%d", &nodeval)
			start, err = addNodeBeforeValue(start, value, nodeval)
			if err != nil {

			}
		case 5:
			var nodeval int
			var value int
			var err error
			fmt.Printf("Enter the value for the node: ")
			fmt.Scanf("%d", &value)
			fmt.Printf("Enter the value after which you want you are node to be added: ")
			fmt.Scanf("%d", &nodeval)
			start, err = addNodeAfterValue(start, value, nodeval)
			if err != nil {
				fmt.Print(err.Error())
				fmt.Println("Not able to find a node with a value", value, "in the list")
			}
		case 11:
			exit = true
		}
	}

}
