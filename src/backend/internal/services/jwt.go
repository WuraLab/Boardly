package services

import (
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// Define its own secret key
var anotherJwtKey = []byte(os.Getenv("ANOTHER_SECRET_KEY"))

// GenerateNonAuthToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateNonAuthToken(userID string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(1440 * time.Minute)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(anotherJwtKey)

	return tokenString, err
}

// DecodeNonAuthToken handles decoding a jwt token
func DecodeNonAuthToken(tkStr string) (string, error) {
	claims := &Claims{}

    // Decode token based on parameters provided, if it fails throw err
	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return anotherJwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

    // Return encoded email
	return claims.UserID, nil
}