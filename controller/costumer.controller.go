package controller

import (
	"github.com/petseepdufrpe/petspeed-golang-api/environment"
)

// CostumerRegister is the function to register a costumer in database.
func CostumerRegister(PersonID int64) {
	database := environment.ConnectDatabase()

	costumerStatement, _ := database.Prepare("insert into costumer(person_id) values(?)")
	costumerStatement.Exec(PersonID)

	defer database.Close()
	return
}
