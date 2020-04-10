package environment

import (
	"database/sql"
	// just ignore this line
	_ "github.com/go-sql-driver/mysql"
)

// ConnectDatabase is the function to connect to our database on umbler.
func ConnectDatabase() (Database *sql.DB) {
	Database, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	return Database
}
