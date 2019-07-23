package models

type User struct {
	Id        string   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Credentials struct {
	Email string
	Password string
}