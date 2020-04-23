package services

import (
	d "MOODOOW/BACK-END-GOLANG/data"
	u "MOODOOW/BACK-END-GOLANG/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CreateUsers : create a new user
func CreateUsers(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("reqBody: %v\n", reqBody)

	var newUser u.Users

	fmt.Printf("Before : %v\n", newUser)

	err := json.Unmarshal(reqBody, &newUser)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("After: %v\n", newUser)

	fmt.Printf("%v, we are creating our account...\n", newUser.Pseudo)

	d.SetNewUser(newUser)
}
