package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (h IntegerHeap) Len() int           { return len(h) }
func (h IntegerHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntegerHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntegerHeap) Minimum() int       { return h[0] }

func (h *IntegerHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntegerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func PopAllElements(h *IntegerHeap) {
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}
func DisplayAllElements(h IntegerHeap) {
	for _, v := range h {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func addElement(h *IntegerHeap) {

	var n int
	fmt.Printf("Enter the element you wish to add: ")
	fmt.Scanf("%d", &n)
	heap.Push(h, n)
}

func RemoveElement(h *IntegerHeap) {
	var index int
	fmt.Printf("Enter the index you want to remove")
	fmt.Scanf("%d", &index)
	fmt.Println("Element removed: ", heap.Remove(h, index))
}

func menu() {
	fmt.Println(" ********* MAIN MENU *********")
	fmt.Println("1. Add element")
	fmt.Println("2. Print Minimum")
	fmt.Println("3. Display All the elements")
	fmt.Println("4. Pop the last element")
	fmt.Println("5. Pop all the elements")
	fmt.Println("6. Remove the element")
	fmt.Println("7. Exit")
	fmt.Printf("Enter your option: ")
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{}
	heap.Init(intHeap)
	for {
		var opt int
		menu()
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1:
			addElement(intHeap)
		case 2:
			fmt.Println((*intHeap)[0])
		case 3:
			DisplayAllElements(*intHeap)
		case 4:
			fmt.Println("Last element is : ", heap.Pop(intHeap))
		case 5:
			PopAllElements(intHeap)
		case 6:
			RemoveElement(intHeap)
		case 7:
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Invalid Option")
		}
	}

}
