package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	d "MOODOOW/BACK-END-GOLANG/data"
	t "MOODOOW/BACK-END-GOLANG/types"
)

var jwtKey = []byte("moodoow")

// Signin : Function to get user connexion info and compare to the DB
func Signin(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newUser t.UserConnectRequest
	json.Unmarshal(reqBody, &newUser)

	fmt.Println("Connection attempt from : ", newUser.Pseudo)

	connection := d.ConnectUser(newUser.Pseudo)

	if !connection.Ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authInfo := t.AuthInfo{
		ID:     connection.ID,
		Pseudo: connection.Pseudo,
		Pays:   connection.Pays,
		Ville:  connection.Ville,
	}

	token := GenerateToken(w, authInfo)
	answer := t.SigninAnswer{
		Token: token,
	}

	answer.ProfileOk = d.ProfilOk(connection.ID)

	json.NewEncoder(w).Encode(answer)
}

// GenerateToken : Function to generate the token
func GenerateToken(w http.ResponseWriter, authInfo t.AuthInfo) string {
	expirationTime := time.Now().Add(2 * time.Minute)

	authInfo.StandardClaims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authInfo)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return ""
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path:    "/",
		Expires: expirationTime,
	})
	return tokenString

}
