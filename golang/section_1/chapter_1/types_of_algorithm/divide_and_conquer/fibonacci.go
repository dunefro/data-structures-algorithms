package main

import "fmt"

func fib(k int) int {
	if k <= 1 {
		return 1
	} else {
		return fib(k-1) + fib(k-2)
	}
}

func main() {
	fmt.Printf("Enter the value of n: ")
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}
}
