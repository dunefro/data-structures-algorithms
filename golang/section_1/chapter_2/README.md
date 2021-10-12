# Arrays
1. Arrays have fixed size.
2. To add more elements in an array we need to create a bigger array, copy the old elements into it and then add the new array elements.
3. If the array is big then the above operation might give performance issues.  

Creating arrays 
```
var arr = [5]int{1, 2, 3, 4, 5}
```
Checking type of `arr`
```
[5]int
```
# Slices
1.An array has a fix size and so elements can't be added to it. To solve this problem we create a slice which has a capacity of double the length of elemnts with which it is initialized.
2. The slice is just a reference to the pointer of array that is created in the background when slice is initialized.
3. Observe the program `slice-pointer.go`. In this when we change the element in slice `b` then the element in slice `a` also changes. This is because they both point to the same array.
4. Similar to this, the program `slice.go` states that the slice `slice` has a capacity similar to the underlying array `arr` starting from its lower bound which is pointed to the index 1 of array `arr`.
