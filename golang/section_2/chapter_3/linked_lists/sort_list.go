package main

func sortList(start *node) *node {
	ptr1 := start
	for {
		if ptr1.next == nil {
			break
		}
		ptr2 := ptr1.next
		for {
			if ptr2 == nil {
				break
			}
			if ptr1.val > ptr2.val {
				ptr1.val, ptr2.val = ptr2.val, ptr1.val
			}
			ptr2 = ptr2.next
		}
		ptr1 = ptr1.next
	}
	return start
}
