package environment

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var environment = mysql.ConnectionURL{
	Host:     "haut.id",
	User:     "hautid_desenvolvimento",
	Password: "desenvolvimento2019",
	Database: "hautid_desenvolvimento"}

func connectDatabase() (Database db.Database) {
	Database, error := mysql.Open(environment)
	if error != nil {
		log.Fatal(error.Error())
	}
	return Database
}
