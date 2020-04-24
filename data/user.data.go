package data

import (
	"fmt"

	e "MOODOOW/BACK-END-GOLANG/error"
	t "MOODOOW/BACK-END-GOLANG/types"

	"golang.org/x/crypto/bcrypt"
)

// SetNewUser : set user in database
func SetNewUser(user t.Users) {
	db := ConnectDb()
	defer db.Close()

	cryptPwd, err := HashPassword(user.Password)
	e.HandleError(err)

	insert, err := db.Prepare("INSERT INTO users(pseudo, nom, prenom, gender, email, date_naissance, pays, ville, nationalite, password) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	defer insert.Close()
	e.HandleError(err)
	_, err = insert.Exec(user.Pseudo, user.Nom, user.Prenom, user.Gender, user.Email, user.DateNaissance, user.Pays, user.Ville, user.Nationalite, cryptPwd)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("User %s was created", user.Pseudo)
	}
}

// GetUserInfo :  recover one user where ID
func GetUserInfo(id int) {
	db := ConnectDb()
	defer db.Close()

	var arrUser []t.Users
	var user t.Users

	row, err := db.Query("Select ID, pseudo, nom, prenom, gender, email, pays, ville, nationalite, password from users WHERE ID=?", id)
	e.HandleError(err)

	for row.Next() {
		err := row.Scan(&user.ID, &user.Pseudo, &user.Nom, &user.Prenom, &user.Gender, &user.Email, &user.Pays, &user.Ville, &user.Nationalite, &user.Password)
		if err != nil {
			panic(err.Error())
		} else {
			arrUser = append(arrUser, user)
		}
	}
	for _, v := range arrUser {
		fmt.Println(v)
	}
}

// GetUserInfoByPseudo : get user info using the pseudo
func GetUserInfoByPseudo(pseudo string) []t.Users {
	db := ConnectDb()
	defer db.Close()

	var arrUser []t.Users
	var user t.Users

	row, err := db.Query("Select ID, pseudo, nom, prenom, gender, email, pays, ville, nationalite, password from users WHERE pseudo=?", pseudo)
	e.HandleError(err)

	for row.Next() {
		err := row.Scan(&user.ID, &user.Pseudo, &user.Nom, &user.Prenom, &user.Gender, &user.Email, &user.Pays, &user.Ville, &user.Nationalite, &user.Password)
		if err != nil {
			panic(err.Error())
		} else {
			arrUser = append(arrUser, user)
		}
	}
	for _, v := range arrUser {
		fmt.Println(v)
	}

	return arrUser
}

// Update user

// Delete user
func Delete(id int) {
	db := ConnectDb()
	defer db.Close()
	delUser, err := db.Prepare("DELETE FROM users WHERE ID=?")
	e.HandleError(err)
	delUser.Exec(id)
	fmt.Printf("User with id %v was deleted successlfully.", id)

}

// Modify password

// HashPassword : Function to encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPwd : Function to compare password to saved hash
func CheckPwd(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
