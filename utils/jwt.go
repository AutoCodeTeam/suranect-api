package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type MyJWTClaims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

var secret = []byte("can-you-keep-a-secret?")

func CreateToken(sub string, userInfo interface{}) (string, error) {
	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Generate your own secret key!
	// Choose an expiration time. Shorter the better
	exp := time.Now().Add(time.Hour * 24)
	// Add your claims
	token.Claims = &MyJWTClaims{
		&jwt.RegisteredClaims{
			// Set the exp and sub claims. sub is usually the userID
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   sub,
		},
		userInfo,
	}
	// Sign the token with your secret key
	val, err := token.SignedString(secret)

	if err != nil {
		// On error return the error
		return "", err
	}
	// On success return the token string
	return val, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
