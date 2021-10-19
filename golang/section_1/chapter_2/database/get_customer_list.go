package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()
	var err error
	var update *sql.Stmt
	update, err = database.Prepare("UPDATE customer SET ssn=? WHERE id=?")
	if err != nil {
		fmt.Errorf("Not able to update the name of the customer")
	}
	update.Exec(customer.SSN, customer.CustomerId)
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "qwerty"
	databaseName := "crm"
	// databaseCon := databaseUser + ":" + databasePass + "@/localhost:3306" + databaseName
	databaseCon := databaseUser + ":" + databasePass + "@/" + databaseName

	database, err := sql.Open(databaseDriver, databaseCon)
	if err != nil {
		fmt.Errorf("Not able to open error")
		panic(err.Error())
	}
	return
}

func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()
	var err error
	var rows *sql.Rows

	// rows, err = database.Query("SELECT * from customers ORDER BY customerID DESC")
	rows, err = database.Query("SELECT * from customer")

	if err != nil {
		panic(err.Error())
	}
	var customers []Customer

	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		e := rows.Scan(&customerId, &customerName, &ssn)
		if e != nil {
			panic(e.Error())
		}
		customer := Customer{
			CustomerId:   customerId,
			CustomerName: customerName,
			SSN:          ssn,
		}
		customers = append(customers, customer)
	}
	return customers
}

func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()
	var err error
	var insert *sql.Stmt
	insert, err = database.Prepare("INSERT INTO customer VALUES (?,?,?);")
	if err != nil {
		fmt.Errorf("Failed to insert the customer details")
		panic(err.Error())
	}
	insert.Exec(customer.CustomerId, customer.CustomerName, customer.SSN)
}

func DeleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()
	var err error
	var delete *sql.Stmt

	delete, err = database.Prepare("DELETE FROM customer WHERE id=?")
	if err != nil {
		fmt.Errorf("Not able to delete the customer")
		panic(err.Error())
	}
	delete.Exec(customer.CustomerId)
}

func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println(customers)
	InsertCustomer(Customer{
		CustomerId:   4,
		CustomerName: "Ron Weasely",
		SSN:          "3937849",
	})
	fmt.Println()
	customers = GetCustomers()
	fmt.Println("After adding Ron")
	fmt.Println(customers)
	UpdateCustomer(Customer{
		CustomerId:   4,
		CustomerName: "Ron Weasely",
		SSN:          "0000",
	})
	fmt.Println()
	fmt.Println("Updating Ron's SSN value")
	customers = GetCustomers()
	fmt.Println(customers)
	DeleteCustomer(Customer{
		CustomerId:   4,
		CustomerName: "Ron Weasely",
		SSN:          "0000",
	})
	fmt.Println()
	fmt.Println("After deleting Ron's entry")
	customers = GetCustomers()
	fmt.Println(customers)

}
