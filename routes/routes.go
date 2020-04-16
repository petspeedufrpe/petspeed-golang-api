package routes

import (
	"fmt"
	"log"
	"net/http"
	"petspeed-golang-api/controller"
)

// Router é onde fica a porra das rotas
func Router() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/registerUser", controller.UserRegister)
	http.HandleFunc("/user/login", controller.Login)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Server is running at :3000...")
}
