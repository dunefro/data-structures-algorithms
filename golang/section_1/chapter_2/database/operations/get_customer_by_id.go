package operations

import (
	"database/sql"
	"fmt"
)

// type Customer struct {
// 	CustomerId   int
// 	CustomerName string
// 	SSN          string
// }

func GetCustomerByID(id int) Customer {
	var database *sql.DB
	database = getConnection()
	defer database.Close()
	var err error

	var rows *sql.Rows

	rows, err = database.Query("SELECT * FROM customer WHERE id=? ", id)

	if err != nil {
		fmt.Errorf("Not able to execute the query to find the customer with the id ", id)
		panic(err.Error())
	}
	var customer Customer
	for rows.Next() {
		var customerID int
		var customerName string
		var SSN string
		err = rows.Scan(&customerID, &customerName, &SSN)
		if err != nil {
			fmt.Errorf("Not able to take the customer credentials out")
			panic(err.Error())
		}
		customer = Customer{
			CustomerId:   customerID,
			CustomerName: customerName,
			SSN:          SSN,
		}
	}

	return customer
}
