package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	d "MOODOOW/BACK-END-GOLANG/data"
	t "MOODOOW/BACK-END-GOLANG/types"
)

// Signin : Function to get user connexion info and compare to the DB
func Signin(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newUser t.UserConnectRequest

	json.Unmarshal(reqBody, &newUser)

	arrUser := d.GetUserInfoByPseudo(newUser.Pseudo)
	hash := arrUser[0].Password

	connOK := d.CheckPwd(newUser.Password, hash)
	fmt.Println(connOK)

}
