package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func powerSeriesN(a int) (square int, cube int) {
	square = a * a

	cube = a * a * a

	return
}

func powerSeriesE(a int) (int, int, error) {
	var sq = a * a
	var cu = a * a * a

	return sq, cu, nil
}

func main() {
	fmt.Println(powerSeries(3))
	fmt.Println(powerSeriesN(4))
	fmt.Println(powerSeriesE(5))
}
