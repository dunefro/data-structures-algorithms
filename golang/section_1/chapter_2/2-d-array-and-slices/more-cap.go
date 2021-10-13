package main

import "fmt"

func main() {
	var arr = [7]int{1, 2, 3, 4, 5, 6, 7}
	s := arr[3:7]
	fmt.Println(len(s))
	fmt.Println(cap(s))
	// s[7] = 8 this will give error
	s = append(s, 8)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
