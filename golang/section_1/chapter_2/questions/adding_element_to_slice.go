package main

import "fmt"

func main() {
	arr := make([]int, 5)
	slice := arr[0:2]
	slice[0] = 1
	slice[1] = 2
	s = append(slice, 3)
	fmt.Println(slice)

}
