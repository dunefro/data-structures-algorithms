package main

import "fmt"

func main() {
	var demon = make(map[string]string)
	demon["key1"] = "val1"
	demon["key2"] = "val2"
	demon["key3"] = "val3"
	for k, v := range demon {
		fmt.Println(k, v)
	}
}
