package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"petspeed-golang-api/environment"
	"petspeed-golang-api/model"

	"golang.org/x/crypto/bcrypt"
)

//UserRegister is the function to register a user in database
func UserRegister(writer http.ResponseWriter, request *http.Request) {
	database := environment.ConnectDatabase()
	bodyJSON, bodyData := ioutil.ReadAll(request.Body)
	if bodyData != nil {
		panic(bodyData)
	}
	var userPerson model.UserPerson
	bodyData = json.Unmarshal(bodyJSON, &userPerson)

	verifyEmail := VerifyEmail(userPerson.Email)
	verifyPerson := VerifyDocument(userPerson.Document)
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(userPerson.Password), 10)

	if verifyEmail == true || verifyPerson == true {
		error, _ := json.Marshal("This e-mail or document is already registered in our database.")
		writer.Write(error)
		defer database.Close()
	} else {

		userStatement, _ := database.Prepare("insert into user(email, password) values(?,?)")
		userResponse, _ := userStatement.Exec(&userPerson.Email, encryptedPassword)
		userID, _ := userResponse.LastInsertId()

		personID := PersonRegister(userPerson.Name, userPerson.Document, userID)
		if personID != 0 {
			CostumerRegister(personID)
		}
		defer database.Close()
	}
	return

}

//VerifyEmail is the function to verify se ja tem a porra de um email cadastrado
func VerifyEmail(email string) bool {
	database := environment.ConnectDatabase()
	var user model.User

	verifyStatement := database.QueryRow("select id, email, password from user where email = ?", email)
	verifyStatement.Scan(&user.ID, &user.Email, &user.Password)
	if user.Email != "" {
		return true
	}
	defer database.Close()

	return false
}
