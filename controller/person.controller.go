package controller

import (
	"github.com/petspeedufrpe/petspeed-golang-api/environment"
	"github.com/petspeedufrpe/petspeed-golang-api/model"
)

// VerifyDocument is the function to verify if the document is already registered in database
func VerifyDocument(document string) bool {
	database := environment.ConnectDatabase()
	var person model.Person

	verifyStatement := database.QueryRow("Select person_id, name, document, user_id from person where document = ?", document)

	verifyStatement.Scan(&person.PersonID, &person.Name, &person.Document, &person.UserID)
	if person.Document != "" {
		return true
	}
	defer database.Close()

	return false
}

// PersonRegister is the function to register the person on database.
func PersonRegister(name string, document string, userID int64) int64 {
	database := environment.ConnectDatabase()

	personStatement, _ := database.Prepare("insert into person(name, document, user_id) values(?,?,?)")
	personResponse, _ := personStatement.Exec(name, document, userID)
	personID, _ := personResponse.LastInsertId()

	defer database.Close()

	return personID
}
