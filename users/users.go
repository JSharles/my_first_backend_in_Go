package users

import "time"

type Users struct {
	Pseudo        string
	Nom           string
	Prenom        string
	Gender        string
	Email         string
	DateNaissance time.Time
	Pays          string
	Ville         string
	Nationalite   string
	Password      string
}
