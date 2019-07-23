package models

import "github.com/dgrijalva/jwt-go"

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

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}