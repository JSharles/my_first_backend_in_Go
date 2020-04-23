package services

import (
	d "MOODOOW/BACK-END-GOLANG/data"
	t "MOODOOW/BACK-END-GOLANG/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	s "strings"
)

// CreateUser : create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("reqBody: %v\n", reqBody)

	var newUser t.Users

	fmt.Printf("Before : %v\n", newUser)

	err := json.Unmarshal(reqBody, &newUser)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("After: %v\n", newUser)

	fmt.Printf("%v, we are creating our account...\n", newUser.Pseudo)

	d.SetNewUser(newUser)
}

// DeleteUser : delete a user in the DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	fmt.Println(path)

	id := s.Split(path, "/")

	fmt.Println(id[2])
	idInt, err := strconv.Atoi(id[2])
	if err != nil {
		fmt.Println(err)
	}
	d.Delete(idInt)
}
