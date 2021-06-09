package application

import (
	"fmt"
	"net/http"
	"test-stone/config"
	"test-stone/domain"
	"test-stone/domain/tokenJWT"
	"test-stone/repository/infrastructure"

	"golang.org/x/crypto/bcrypt"
)

func DoLogin(accLogin *domain.Login) (*string, error) {
	response := "Senha incorreta "

	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := infrastructure.AccountRepositoryWithRDB(conn)
	acc, err := repo.LoginRepository(accLogin.CPF)

	if err != nil {
		return nil, err
	}

	if IsCorrectPassword(accLogin.Secret, acc.Secret) {
		token, err := tokenJWT.CreateToken(acc)
		if err != nil {
			return nil, err
		}

		return &token, nil
	}

	return &response, err
}

func IsCorrectPassword(SecretRequest string, SecretResponse string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(SecretResponse), []byte(SecretRequest))
	fmt.Println(err)
	return err == nil
}

func AuthLogin(r *http.Request) bool {

	login, err := tokenJWT.ExtractTokenMetadata(r)
	if err != nil {
		return err == nil
	}

	conn, err := config.ConnectDB()
	if err != nil {
		return err == nil
	}
	defer conn.Close()

	repo := infrastructure.AccountRepositoryWithRDB(conn)
	acc, err := repo.LoginRepository(login.CPF)

	if err != nil {
		return err == nil
	}

	if login.Secret == acc.Secret {

		return true
	} else {
		return false
	}

}
