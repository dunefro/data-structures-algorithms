package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(cap(arr))
}
