package main

import "fmt"

func main() {
	var arr = [7]int{1, 2, 3, 4, 5, 6, 7}

	// Creating two slices
	a := arr[0:4]
	b := arr[1:5]
	fmt.Println(a, b)

	// changing the element in a would change it in b as well because both refer to the same array in background
	b[0] = 29
	fmt.Println(a, b)
}
