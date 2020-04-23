package data

import (
	"fmt"

	t "MOODOOW/BACK-END-GOLANG/types"

	"golang.org/x/crypto/bcrypt"
)

// SetNewUser : set user in database
func SetNewUser(user t.Users) {
	db := ConnectDb()
	defer db.Close()

	cryptPwd, err := HashPassword(user.Password)
	if err != nil {
		panic(err.Error())
	}

	insert, err := db.Prepare("INSERT INTO users(pseudo, nom, prenom, gender, email, date_naissance, pays, ville, nationalite, password) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	defer insert.Close()
	if err != nil {
		panic(err.Error())
	}

	_, err = insert.Exec(user.Pseudo, user.Nom, user.Prenom, user.Gender, user.Email, user.DateNaissance, user.Pays, user.Ville, user.Nationalite, cryptPwd)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("User %s was created", user.Pseudo)
	}
}

// GetAllUsers :  recover all users in database
func GetAllUsers() {
	db := ConnectDb()
	defer db.Close()

	var arrUsers []t.Users
	var users t.Users

	rows, err := db.Query("Select ID, pseudo, nom, prenom, gender, email, date_naissance, pays, ville, nationalite, password from users")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&users.ID, &users.Pseudo, &users.Nom, &users.Prenom, &users.Gender, &users.Email, &users.DateNaissance, &users.Pays, &users.Ville, &users.Nationalite, &users.Password)
		if err != nil {
			panic(err.Error())
		} else {
			arrUsers = append(arrUsers, users)
		}
	}
	for _, v := range arrUsers {
		fmt.Println(v)
	}
}

// GetUser :  recover one user where ID
func GetUser(id int) {
	db := ConnectDb()
	defer db.Close()

	var arrUser []t.Users
	var user t.Users

	row, err := db.Query("Select ID, pseudo, nom, prenom, gender, email, pays, ville, nationalite, password from users WHERE ID=?", id)
	if err != nil {
		panic(err.Error())
	}

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

// Update user

// Delete user

func DeleteUser(id int) {
	db := ConnectDb()
	defer db.Close()
	delUser, err := db.Prepare("DELETE FROM users WHERE ID=?")
	if err != nil {
		panic(err)
	}
	delUser.Exec(id)
	fmt.Printf("User with id %v was deleted successlfully.", id)
}

// Modify password

// Function to encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Function to compare password to saved hash

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
