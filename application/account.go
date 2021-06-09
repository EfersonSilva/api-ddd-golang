package application

import (
	"test-stone/config"
	"test-stone/domain"
	"test-stone/repository/infrastructure"

	"github.com/biezhi/gorm-paginator/pagination"
)

//Adiciona conta
func CreateAccount(account domain.Account) error {

	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := infrastructure.AccountRepositoryWithRDB(conn)

	return repo.InsertAccount(&account)
}

// retorna conta por id
func GetAccountById(id int) (*domain.Account, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := infrastructure.AccountRepositoryWithRDB(conn)
	acc, err := repo.FindAccountById(id)
	if err != nil {
		return nil, err
	}

	acc.Secret = ""
	acc.Balance = 0
	return acc, err
}

//retorna saldo da conta por id
func GetBalanceAccountById(id int) (*domain.Account, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := infrastructure.AccountRepositoryWithRDB(conn)
	acc, err := repo.FindAccountById(id)

	if err != nil {
		return nil, err
	}

	return acc, nil
}

// Retorna todas as contas cadastradas
func GetAllAccounts(limit int, page int) ([]domain.Account, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var accounts []domain.Account
	pagination.Paging(&pagination.Param{
		DB:      conn.Preload("Topic"),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &accounts)

	for i, element := range accounts {
		element.Secret = ""
		element.Balance = 0
		accounts[i] = element
	}

	return accounts, nil
}

//Atualiza Conta
func UpdateAccount(account domain.Account) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := infrastructure.AccountRepositoryWithRDB(conn)

	return repo.UpdateAccount(&account)
}

// exclui conta
func DeleteAccount(id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := infrastructure.AccountRepositoryWithRDB(conn)
	return repo.DeleteAccount(id)
}
