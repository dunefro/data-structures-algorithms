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

func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println(customers)
}
