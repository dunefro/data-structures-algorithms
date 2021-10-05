# Structural Design Patterns
Structural patterns explain how to assemble objects and classes into larger structures while keeping these structures flexible and efficient.

## Adapter
1. Theory can be understood https://refactoring.guru/design-patterns/adapter
2. We have two programs to understand adapter in the `adapter/` folder.

### Iprocess
We have an `Iprocess` interface which is implementing `process()` function. The `Adapter` class(struct) is to be made of type `Iprocess` so it implements the `process()` function. The `Adapter` class has an attribute of type `Adaptee` which has a method called `convert()`, called inside of the process method.
```
fmt.Println(processor.adaptee.adapterType)
```
This would give an error because processor is of type `Iprocess` interface and this inteface doesn't have `adaptee` attribute. 
```
fmt.Println(Adapter{}.adaptee.adapterType)
```
This would give `0` as expected because `adaptee` is an attribute of `Adapter` class.