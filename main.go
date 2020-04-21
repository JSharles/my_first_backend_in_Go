package main

import (
	"MOODOOW/BACK-END-GOLANG/routes"
)

func main() {
	// dbCon := data.ConnectDb()
	// defer dbCon.Close()

	// insert, err := dbCon.Query("INSERT INTO users(pseudo, nom, prenom, gender, email, pays, ville, nationalite) VALUES ('pauloo', 'rey', 'paul', 'homme', 'rey.paul31@gmail.com', 'France', 'Gratentour', 'Francais');")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	routes.CreateUsers("Toto", "Toto", "Toto", "Male", "toto42@gmail.com", "1942-01-01", "toto land", "toto city", "unknown", "toto42")

}
