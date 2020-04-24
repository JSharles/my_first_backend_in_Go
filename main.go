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
	myRouter.HandleFunc("/get/{id}", s.GetUserByID)
	myRouter.HandleFunc("/signin", s.Signin)

	/////////////////////////////
	////////// EVENT ///////////
	///////////////////////////
	// Listen on port : ...
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {

	fmt.Println("Starting Moodoow Backend on port 5000")
	handleRequests()

}
