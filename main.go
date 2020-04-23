package main

import (
	"fmt"
	"log"
	"net/http"

	s "MOODOOW/BACK-END-GOLANG/services"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter()

	/////////////////////////////
	/////////// USER ///////////
	///////////////////////////
	//Create an Account
	myRouter.HandleFunc("/register", s.CreateUser)
	myRouter.HandleFunc("/deleteuser/{id}", s.DeleteUser)

	/////////////////////////////
	////////// EVENT ///////////
	///////////////////////////
	// Listen on port : ...
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	//routes.CreateUsers("Pau", "Rey", "Paul", "Homme", "pau@paul.com", "1996-12-10", "france", "toulouse", "francaise", "paul")
	//routes.GetAllUsers()
	//routes.GetUser(2)

	fmt.Println("Starting Moodoow Backend on port 5000")
	handleRequests()

}
