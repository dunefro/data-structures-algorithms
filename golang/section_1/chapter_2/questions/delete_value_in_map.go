package main

import "fmt"

func main() {
	var del = make(map[string]int)
	del["k1"] = 1
	del["k2"] = 2
	del["k3"] = 3
	fmt.Println("Before Deleting", del)
	delete(del, "k3")
	fmt.Println("After deleting", del)
}
