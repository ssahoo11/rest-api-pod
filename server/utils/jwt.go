package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "$2a$14$EnBxC"

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		_, ok := tkn.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	validToken := parsedToken.Valid

	if !validToken {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token type")
	}
	// var user *models.User
	// user.Email = claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	// if !user.Exists() {
	// 	return errors.New("Invalid login")
	// }
	return userId, nil
}
