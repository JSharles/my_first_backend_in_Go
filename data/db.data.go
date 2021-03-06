package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDb provides connexion infos to the DB
func ConnectDb() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "jnchrlsbrq"
	dbHost := "(localhost:3306)"
	dbName := "MOODOOW"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp"+dbHost+"/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to port 3306")
	}
	return db

}
