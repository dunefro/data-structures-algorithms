package main

import "fmt"

func main() {
	var languages = map[int]string{
		1: "string1",
		2: "string2",
		3: "string3",
		4: "string4",
	}
	var products = make(map[int]string)
	products[1] = "chair"
	products[2] = "table"

	for i, v := range languages {
		fmt.Println(i, v)
	}
	fmt.Println("the second product is", products[1])
	fmt.Println(products)
}
