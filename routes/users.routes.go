package routes

import (
	"MOODOOW/BACK-END-GOLANG/data"
	"MOODOOW/BACK-END-GOLANG/users"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUsers(pseudo, nom, prenom, gender, email, dateNaissance, pays, ville, nationalite, password string) {
	db := data.ConnectDb()
	defer db.Close()

	dateParse, err := time.Parse("2006-01-02", dateNaissance)
	if err != nil {
		panic(err.Error())
	}

	cryptPwd, _ := HashPassword(password)

	user := users.Users{
		Pseudo:        pseudo,
		Nom:           nom,
		Prenom:        prenom,
		Gender:        gender,
		Email:         email,
		DateNaissance: dateParse,
		Pays:          pays,
		Ville:         ville,
		Nationalite:   nationalite,
		Password:      cryptPwd,
	}

	insert, err := db.Prepare("INSERT INTO users(pseudo, nom, prenom, gender, email, date_naissance, pays, ville, nationalite, password) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	defer insert.Close()
	if err != nil {
		panic(err.Error())
	}

	_, err = insert.Exec(user.Pseudo, user.Nom, user.Prenom, user.Gender, user.Email, user.DateNaissance, user.Pays, user.Ville, user.Nationalite, user.Password)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("%s was added in the database\n", user.Pseudo)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
