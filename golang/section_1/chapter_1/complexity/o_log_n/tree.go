package main

import "fmt"

type tree struct {
	leftNode  *tree
	value     int
	rightNode *tree
}

func (t *tree) insert(val int) {
	if t != nil {
		if t.leftNode == nil {
			t.leftNode = &tree{
				leftNode:  nil,
				value:     val,
				rightNode: nil,
			}
		} else {
			if t.rightNode == nil {
				t.rightNode = &tree{
					leftNode:  nil,
					value:     val,
					rightNode: nil,
				}
			} else {
				if t.leftNode != nil {
					t.leftNode.insert(val)
				} else {
					t.rightNode.insert(val)
				}
			}
		}
	} else {
		t = &tree{
			leftNode:  nil,
			value:     val,
			rightNode: nil,
		}
	}
}

func print(tree *tree) {
	if tree != nil {
		fmt.Println("Value:", tree.value)
		if tree.leftNode != nil {
			fmt.Println("Tree left Node")
			print(tree.leftNode)
		}
		if tree.rightNode != nil {
			fmt.Println("Tree right node")
			print(tree.rightNode)
		}
	} else {
		fmt.Println("NIL")
	}
}

func main() {
	var tree *tree = &tree{nil, 1, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.leftNode.insert(7)
	print(tree)
}
