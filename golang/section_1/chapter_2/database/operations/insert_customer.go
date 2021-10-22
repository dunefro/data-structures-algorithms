package operations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = getConnection()
	defer database.Close()
	var err error
	var insert *sql.Stmt
	insert, err = database.Prepare("INSERT INTO customer (name, ssn) VALUES (?,?);")
	if err != nil {
		fmt.Errorf("Failed to insert the customer details")
		panic(err.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
}
