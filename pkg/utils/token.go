package utils

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTKey is secret key for jwtoken
var JWTKey = "my-secret-key"

// CustomClaims provides email information along with standard infos about jwt. claims
type CustomClaims struct {
	Email string
	jwt.StandardClaims
}

// JWToken is function for creating jwt and sending it back to the user
func JWToken(email string) (string, *Response) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "lenses.com",
		},
	})
	signedToken, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		return "", Back(http.StatusUnauthorized, "You are not authenticated. Please login again")
	}
	return signedToken, nil
}
