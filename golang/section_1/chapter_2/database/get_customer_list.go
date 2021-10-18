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

func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println(customers)
	InsertCustomer(Customer{
		CustomerId:   4,
		CustomerName: "Ron Weasely",
		SSN:          "3937849",
	})
	// time.Sleep(2 * time.Second)
	customers = GetCustomers()
	fmt.Println("After adding Ron")
	fmt.Println(customers)
}
