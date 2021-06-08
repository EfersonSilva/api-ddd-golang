package tokenJWT

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"test-stone/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(login *domain.Login) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "stone")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["login"] = login
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "Error in your key of token", err
	}
	return token, nil
}

func ExtractToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1], nil
	}
	return "", nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, err := ExtractToken(r)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*domain.Login, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		acc, ok := claims["login"].(map[string]interface{})
		if !ok {
			return nil, err
		}
		cpf := acc["cpf"].(string)
		secret := acc["secret"].(string)

		return &domain.Login{
			CPF:    cpf,
			Secret: secret,
		}, nil
	}
	return nil, err
}
