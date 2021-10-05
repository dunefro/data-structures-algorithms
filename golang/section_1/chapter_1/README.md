1. Linked, ordered, and unordered lists are grouped as heterogeneous data structures. Linear data structures are lists, sets, tuples, queues, stacks, and heaps. Trees, tables, and containers are categorized as nonlinear data structures. Two dimensional and multidimensional arrays are grouped as homogeneous data structures. Dynamic data structures are dictionaries, tree sets, and sequences.
2. In the `list.go` we don't need to explicitly convert the values to int. It will automatically be converted to the required values. Moreover we can also store string which means the below code snippet would work fine as well.
```
intList.PushBack("Hello")
intList.PushBack(23)
intList.PushBack(14)
```
3. A Heap is a special Tree-based data structure in which the tree is a complete binary tree. Generally, Heaps can be of two types:
    * a. Max-Heap: In a Max-Heap the key present at the root node must be greatest among the keys present at all of it’s children. The same property must be recursively true for all sub-trees in that Binary Tree.
    * b. Min-Heap: In a Min-Heap the key present at the root node must be minimum among the keys present at all of it’s children. The same property must be recursively true for all sub-trees in that Binary Tree. 
    ![Heap data structure](./images/heap.png)
A complete binary tree is the tree in which all the elements are filled except the termianl ones. Level 'r' of the tree will have 2^r nodes. If any element is filled at the 'i' position then its left successor will be at the position '2i' and the right onw will be at '2i+1'.
4. Structural patterns explain how to assemble objects and classes into larger structures while keeping these structures flexible and efficient.
    * a. Adapater: Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate. An example of adapter would be to convert all the data coming into XML format to JSON format for a third-party analytics tool.
    * b. 