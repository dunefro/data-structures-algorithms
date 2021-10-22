package main

import (
	"fmt"

	"sample.com/database/operations"
)

func main() {
	var customers []operations.Customer
	customers = operations.GetCustomers()
	fmt.Println(customers)
	operations.InsertCustomer(operations.Customer{
		CustomerId:   4,
		CustomerName: "Ron Weasely",
		SSN:          "3937849",
	})
	fmt.Println()
	customers = operations.GetCustomers()
	fmt.Println("After adding Ron")
	fmt.Println(customers)
	operations.UpdateCustomer(operations.Customer{
		CustomerId:   4,
		CustomerName: "Ron Weasely",
		SSN:          "0000",
	})
	fmt.Println()
	fmt.Println("Updating Ron's SSN value")
	customers = operations.GetCustomers()
	fmt.Println(customers)
	operations.DeleteCustomer(operations.Customer{
		CustomerId:   4,
		CustomerName: "Ron Weasely",
		SSN:          "0000",
	})
	fmt.Println()
	fmt.Println("After deleting Ron's entry")
	customers = operations.GetCustomers()
	fmt.Println(customers)
	fmt.Println()
	fmt.Println("Retrieving the Customer by ID")
	cust := operations.GetCustomerByID(2)
	fmt.Println(cust)
}
