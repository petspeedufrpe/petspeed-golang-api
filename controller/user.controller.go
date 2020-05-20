package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/petspeedufrpe/petspeed-golang-api/environment"
	"github.com/petspeedufrpe/petspeed-golang-api/model"
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

//Login is the function to login a user
func Login(writer http.ResponseWriter, request *http.Request) {
	database := environment.ConnectDatabase()
	bodyJSON, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var userCostumer model.UserCostumer
	var user model.User
	err = json.Unmarshal(bodyJSON, &user)
	userPassword := findUserByEmail(user.Email)
	encrypted := bcrypt.CompareHashAndPassword(userPassword, []byte(user.Password))
	if encrypted == nil {
		loginQuery := database.QueryRow("select u.id, u.email, u.password, p.person_id, p.name, p.document, p.user_id, c.costumer_id, c.person_id from user as u left join person as p on (p.user_id = u.id) left join costumer as c on (c.person_id = p.person_id) where email = ?", user.Email)
		loginQuery.Scan(&userCostumer.ID, &userCostumer.Email, &userCostumer.Password, &userCostumer.PersonID, &userCostumer.Name, &userCostumer.Document, &userCostumer.UserID, &userCostumer.CostumerID, &userCostumer.CostumerPersonID)
		userJSON, _ := json.Marshal(userCostumer)
		writer.Write(userJSON)

	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
	defer database.Close()
	return
}
func findUserByEmail(email string) []byte {
	database := environment.ConnectDatabase()
	var user model.User
	findQuery := database.QueryRow("select id, email, password from user where email = ?", email)
	findQuery.Scan(&user.ID, &user.Email, &user.Password)
	if user.Password != "" {
		return []byte(user.Password)
	}
	defer database.Close()
	return nil
}
