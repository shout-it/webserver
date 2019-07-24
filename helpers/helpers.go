package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"os"
	"webserver/models"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}


func CreateSignedTokenStringFor(expirationTime int64,credentials models.Credentials) (string,error) {
	pwd, _ := os.Getwd()
	privateKey, err := ioutil.ReadFile(pwd+"/helpers/certificates/jwtRS256.key")
	if err != nil {
		return "",err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "",err
	}
	claims := &models.Claims {
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256,claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}


func ParseTokenFromSignedTokenString(tokenString string) (*jwt.Token,error){
	pwd, _ := os.Getwd()
	publicKey,err := ioutil.ReadFile(pwd+"/helpers/certificates/jwtRS256.key.pub")
	if err != nil {
		return nil,err
	}
	key,err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil,err
	}
	claims :=&models.Claims{}
	parsedToken,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return key, nil
	})
	if err != nil || parsedToken == nil || parsedToken != nil && !parsedToken.Valid{
		return nil,err
	}
	return parsedToken,nil
}