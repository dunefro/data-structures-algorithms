package main

import "fmt"

func main() {
	var rows int
	var cols int
	rows = 2
	cols = 3
	var twodslices = make([][]int, rows)
	for i := 0; i < len(twodslices); i++ {
		twodslices[i] = make([]int, cols)
	}
	fmt.Println(twodslices)
}
