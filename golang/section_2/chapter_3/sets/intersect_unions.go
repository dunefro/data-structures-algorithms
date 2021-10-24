package main

func (set *Set) intersection(iset *Set) *Set {
	newSet := &Set{}
	newSet.New()
	for element, _ := range set.integerMap {
		if iset.containsElement(element) {
			newSet.addElement(element)
		}
	}
	return newSet
}

func (set *Set) union(iset *Set) *Set {
	newSet := &Set{}
	newSet.New()
	for element, _ := range set.integerMap {
		newSet.addElement(element)
	}
	for element, _ := range iset.integerMap {
		if !newSet.containsElement(element) {
			newSet.addElement(element)
		}
	}
	return newSet
}
