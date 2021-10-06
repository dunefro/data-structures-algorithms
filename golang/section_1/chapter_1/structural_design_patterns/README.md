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

### Lightning port
We have a `computer` interface which is implementing `insertIntoLightningPort()` function. Our purpose is to insert USB by client from mac or windows. For mac we can directly invoke the interface function but for windows we need to convert the lightning singal to USB and then insert it into the computer. For this reason client uses the `insertLightningConnectorIntoComputer()` function for type `Computer`. To run this execute the command
```
go run .
```

## Bridge
1. Bridge is a structural design pattern that divides business logic or huge class into separate class hierarchies that can be developed independently. To understand more https://refactoring.guru/design-patterns/bridge.
2. Moreover Birdge consists of two terms `abstraction` and `implementation` as part of its definition.
    * a. `Abstraction` is the high level layer which the client connects to. This layer is not supposed to do any work on its own rather this layer supasses the work to the `implementation` layer.
    * b. `Implementation` is the actual place where the process or the required action is performed.
3. Best example to understand this is the Azure cloud API management. The API management is responsible to handle various clients or requests coming from Azure CLI or Azure portal. This makes the API management an `abstraction` layer and the inner process or inner APIs as the `implementation` layer.
4. In the below picture 
    * a. The Abstraction provides high-level control logic. It relies on the implementation object to do the actual low-level work.
    * b. The Implementation declares the interface that’s common for all concrete implementations. An abstraction can only communicate with an implementation object via methods that are declared here. The abstraction may list the same methods as the implementation, but usually the abstraction declares some complex behaviors that rely on a wide variety of primitive operations declared by the implementation. 
    * c. Concrete Implementations contain platform-specific code.
    * d. Refined Abstractions provide variants of control logic. Like their parent, they work with different implementations via the general implementation interface.
    * e. Usually, the Client is only interested in working with the abstraction. However, it’s the client’s job to link the abstraction object with one of the implementation objects
    ![Bridge Structural design patter](../images/bridge.png)

### Printer
In this printer we are trying to implement both the printer `hp` and `epson` from different OS like `mac` and `windows`. For this their respective files are created. Our purpose it to make client not worried about connecting different computer to different printers.

## Composite
Composite is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects. https://refactoring.guru/design-patterns/composite

### Files and folders
Let’s try to understand the Composite pattern with an example of an operating system’s file system. In the file system, there are two types of objects: files and folders. There are cases when files and folders should be treated to be the same way. This is where the Composite pattern comes in handy.

Imagine that you need to run a search for a particular keyword in your file system. This search operation applies to both files and folders. For a file, it will just look into the contents of the file; for a folder, it will go through all files of that folder to find that keyword.