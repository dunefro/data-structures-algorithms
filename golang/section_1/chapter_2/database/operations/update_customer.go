package operations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = getConnection()
	defer database.Close()
	var err error
	var update *sql.Stmt
	update, err = database.Prepare("UPDATE customer SET ssn=? WHERE name=?")
	if err != nil {
		fmt.Errorf("Not able to update the name of the customer")
	}
	update.Exec(customer.SSN, customer.CustomerName)
}
