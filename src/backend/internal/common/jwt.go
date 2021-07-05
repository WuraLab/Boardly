package common

import {
	"github.com/dgrijalva/jwt-go"
}


func CreateToken(body string) (string, error) {
	var err error
	//SECRET KEY
	var SECRET_KEY := config.SERVER.JWT
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["body"] = body
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(SECRET_KEY))
	if err != nil {
	   return "", err
	}
	return token, nil
  }