package Jwt

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//IssueToken ...
func IssueToken(id, email string) (string, error) {
	type UserInfo struct {
		ID    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
		jwt.StandardClaims
	}

	expire := time.Now().Add(time.Second * 86400).Unix()

	// Create the Claims
	claims := &UserInfo{
		id,
		email,
		jwt.StandardClaims{
			Issuer:    "Tin",
			ExpiresAt: expire,
		},
	}
	return SignToken(claims)
}

//SignToken sign claims
func SignToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)

	if nil != err {
		log.Println("Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
		return ss, err
	}

	return ss, nil
}

//VerificationToken ...
func VerificationToken(tokenString string) (string, string, error) {
	type UserInfo struct {
		ID    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
		jwt.StandardClaims
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &UserInfo{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		return "", "", err
	}
	claims, ok := token.Claims.(*UserInfo)

	if !ok {
		return "", "", errors.New("invalid token")
	}
	fmt.Println("Claims id: ", claims.ID)
	fmt.Println("Claims email: ", claims.Email)
	return claims.ID, claims.Email, err
}

//VerificationRefreshToken ...
func VerificationRefreshToken(tokenString string) (string, string, error) {
	type UmbalaUserInfo struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		jwt.StandardClaims
	}
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &UmbalaUserInfo{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*UmbalaUserInfo)

	if !ok {
		return "", "", errors.New("invalid token")
	}
	return claims.ID, claims.Email, err
}
