package operations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteCustomer(customer Customer) {
	var database *sql.DB
	database = getConnection()
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
func DeleteCustomerByID(id int) {
	var database *sql.DB
	database = getConnection()
	defer database.Close()
	var err error
	var delete *sql.Stmt

	delete, err = database.Prepare("DELETE FROM customer WHERE id=?")
	if err != nil {
		fmt.Errorf("Not able to delete the customer")
		panic(err.Error())
	}
	delete.Exec(id)
}
