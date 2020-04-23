package users

type Users struct {
	ID            int    `json:"id"`
	Pseudo        string `json:"pseudo"`
	Nom           string `json:"nom"`
	Prenom        string `json:"prenom"`
	Gender        string `json:"gender"`
	Email         string `json:"email"`
	DateNaissance string `json:"date_naissance"`
	Pays          string `json:"pays"`
	Ville         string `json:"ville"`
	Nationalite   string `json:"nationalite"`
	Password      string `json:"password"`
}
