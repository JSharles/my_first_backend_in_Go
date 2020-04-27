package types

// Users : Model to store users in Db
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

// UserConnectRequest : Model to make verification in the DB
type UserConnectRequest struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

// UserConnectAnswer : Model to answer the API connexion
type UserConnectAnswer struct {
	Ok     bool   `json:"ok"`
	ID     int    `json:"id"`
	Pseudo string `json:"pseudo"`
	Pays   string `json:"pays"`
	Ville  string `json:"ville"`
}

type SigninAnswer struct {
	Token     string `json:"token"`
	ProfileOk bool   `json:"profileOk"`
}
