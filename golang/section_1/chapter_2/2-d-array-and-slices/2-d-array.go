package main

import "fmt"

func main() {
	var arr [3][3]int

	// Creating identity matrix
	arr[0][0] = 1
	arr[1][1] = 1
	arr[2][2] = 1

	fmt.Println(arr)
}
