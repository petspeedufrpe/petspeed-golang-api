package environment

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var environment = mysql.ConnectionURL{
	Host:     "host",
	User:     "user",
	Password: "password",
	Database: "database"}

func connectDatabase() (Database db.Database) {
	Database, error := mysql.Open(environment)
	if error != nil {
		log.Fatal(error.Error())
	}
	return Database
}
