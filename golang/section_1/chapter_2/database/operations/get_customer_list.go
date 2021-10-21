package operations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetCustomers() []Customer {
	var database *sql.DB
	database = getConnection()
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
