package operations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() (database *sql.DB) {
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
