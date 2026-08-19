package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(
	userID string,
	firstName string,
	lastName string,
	email string,
	role string,
	secret string,
) (string, error) {

	claims := jwt.MapClaims{
		"userId":    userID,
		"firstName": firstName,
		"lastName":  lastName,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"role":      role,
		"email":     email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}

func ValidateToken(
	tokenString string,
	secret string,
) (*jwt.Token, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrTokenSignatureInvalid
			}

			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return token, nil
}
