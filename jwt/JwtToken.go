package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userName string, password string) (string, error) {
	fmt.Println("username : ", userName)
	fmt.Println("password : ", password)

	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("token : ", token)
	tokenString, err := token.SignedString([]byte(password))
	if err != nil {
		fmt.Println(errors.Unwrap(err))
		return "", err
	}
	fmt.Println("Token created : ", tokenString)

	return tokenString, nil
}

func ValidateAccessToken(tokenFromRequest string, password string) (bool, string) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenFromRequest, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(password), nil
	})
	if err != nil {
		fmt.Println(err)
		if err == jwt.ErrSignatureInvalid {
			return false, err.Error()
		}
		return false, err.Error()
	}
	if !tkn.Valid {
		return false, "Token not valid"
	}

	return true, "Token is valid"
}
