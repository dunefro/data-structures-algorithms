package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) addElement(val int) {
	if !set.containsElement(val) {
		set.integerMap[val] = true
	}
}

func (set *Set) deleteElement(val int) {
	delete(set.integerMap, val)
}
func (set *Set) containsElement(element int) bool {
	_, check := set.integerMap[element]
	return check
}

func main() {
	var set *Set
	set = &Set{}
	set.New()
	fmt.Println("--> Creating first set")
	set.addElement(1)
	set.addElement(2)
	set.addElement(3)
	set.addElement(4)
	fmt.Println("--> Printing first set")
	fmt.Println(set)
	fmt.Println("--> Checking if 1 exists")
	fmt.Println(set.containsElement(1))
	fmt.Println("--> Deleting 1")
	set.deleteElement(1)
	fmt.Println("--> Checking if 1 exists")
	fmt.Println(set.containsElement(1))
	fmt.Println("--> Creating second set")
	secondSet := &Set{}
	secondSet.New()
	secondSet.addElement(2)
	secondSet.addElement(4)
	secondSet.addElement(5)
	fmt.Println("--> Printing Second set")
	fmt.Println(secondSet)
	fmt.Println("--> Intersection of two sets")
	setIntersection := set.intersection(secondSet)
	fmt.Println(setIntersection)
	fmt.Println("--> Union of two sets")
	setUnion := set.union(secondSet)
	fmt.Println(setUnion)
}
