package main

import "fmt"

type Iprocess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func main() {
	var processor Iprocess = Adapter{}
	processor.process()
	// fmt.Println(processor.adaptee.adapterType)
	fmt.Println(Adapter{}.adaptee.adapterType)

}
