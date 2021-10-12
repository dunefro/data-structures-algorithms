package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3]
	fmt.Println("Length of array is ", len(slice))
	fmt.Println("Capacity of array is ", cap(slice))

}
