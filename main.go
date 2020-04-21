package main

import (
	"MOODOOW/BACK-END-GOLANG/data"
)

func main() {
	dbCon := data.ConnectDb()
	defer dbCon.Close()

	insert, err := dbCon.Query("INSERT INTO users(pseudo, nom, prenom, gender, email, pays, ville, nationalite) VALUES ('pauloo', 'rey', 'paul', 'homme', 'rey.paul31@gmail.com', 'France', 'Gratentour', 'Francais');")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
