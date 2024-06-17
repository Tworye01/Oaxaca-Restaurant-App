package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

// Creates a JWT token using the given username and password, these are used as claims for the token.
func CreateToken(username string, password string) (string, error) {
	// Creates new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			// Set username and experation time for token
			"username": username,
			"password": password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	// Signs token with secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	// Returns generated token
	return tokenString, nil
}

// Verify that a given token is valid, if it not valid an error will be returned.
func VerifyToken(tokenString string) error {
	// Parse and verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	// Returns an error if token is invalid
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// If token is valid we continue processing the request
	return nil
}

// Extracts username from JWT, through the JWT claims.
func GetUsernameFromToken(tokenString string) (string, string, error) {
	// Verify token
	err := VerifyToken(tokenString)
	if err != nil {
		return "", "", err
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", "", fmt.Errorf("error whilst parsing token")
	}

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("error, invalid token")
	}

	// Extract username from claims
	username, ok := claims["username"].(string)
	if !ok {
		return "", "", fmt.Errorf("error, username not found")
	}
	password, ok := claims["password"].(string)
	if !ok {
		return "", "", fmt.Errorf("error, username not found")
	}
	return username, password, nil
}
