package users

import "time"

type Users struct {
	pseudo        string
	nom           string
	prenom        string
	gender        string
	email         string
	dateNaissance time.Time
	pays          string
	ville         string
	nationalite   string
}
