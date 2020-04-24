package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	d "MOODOOW/BACK-END-GOLANG/data"
	e "MOODOOW/BACK-END-GOLANG/error"
	t "MOODOOW/BACK-END-GOLANG/types"
	s "strings"
)

// CreateUser : create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	e.HandleError(err)
	fmt.Printf("reqBody: %v\n", reqBody)

	var newUser t.Users

	fmt.Printf("Before : %v\n", newUser)

	json.Unmarshal(reqBody, &newUser)

	fmt.Printf("After: %v\n", newUser)

	fmt.Printf("%v, we are creating our account...\n", newUser.Pseudo)

	d.SetNewUser(newUser)
}

// DeleteUser : delete a user in the DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	fmt.Println(path)

	id := s.Split(path, "/")
	idInt, err := strconv.Atoi(id[2])
	e.HandleError(err)
	d.Delete(idInt)
}

// GetUserByID : processes the URL in order to get the user id and passes it to the GetUserInfo function located in the data package
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id := s.Split(path, "/")
	idInt, err := strconv.Atoi(id[2])
	e.HandleError(err)
	d.GetUserInfo(idInt)

}
