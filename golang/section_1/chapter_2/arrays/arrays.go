package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr)
	for i, v := range arr {
		fmt.Println("Index", i, "--- Val ---", v)
	}
}
