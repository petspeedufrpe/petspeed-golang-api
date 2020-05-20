package routes

import (
	"log"
	"net/http"

	"github.com/petspeedufrpe/petspeed-golang-api/controller"
)

// Router é onde fica a porra das rotas
func Router() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/registerUser", controller.UserRegister)
	http.HandleFunc("/user/login", controller.Login)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
