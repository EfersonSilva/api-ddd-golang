package application

import (
	"errors"
	"net/http"
	"test-stone/config"
	"test-stone/domain"
	"test-stone/domain/tokenJWT"
	"test-stone/repository/infrastructure"
)

//Adiciona transferencia
func AddTransfer(transfer domain.Transfer) error {

	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repoTransfer := infrastructure.TrasnferRepositoryWithRDB(conn)
	repoAccount := infrastructure.AccountRepositoryWithRDB(conn)

	accountDestination, err := repoAccount.FindAccountById(transfer.AccountDestinationId)
	if err != nil {
		return err
	}

	accountOrigin, err := repoAccount.FindAccountById(transfer.AccountOriginId)
	if err != nil {
		return err
	}

	if accountOrigin.Balance <= transfer.Amount {
		return errors.New("invalid account balance")
	} else {
		accountOrigin.Balance -= transfer.Amount
		err := repoAccount.UpdateAccount(accountOrigin)
		if err != nil {
			return err
		}

		accountDestination.Balance += transfer.Amount
		err = repoAccount.UpdateAccount(accountDestination)
		if err != nil {
			return err
		}

		var comprovante domain.Transfer
		comprovante.AccountDestinationId = transfer.AccountDestinationId
		comprovante.AccountOriginId = transfer.AccountOriginId
		comprovante.Amount = transfer.Amount
		comprovante.CPF = accountOrigin.CPF
		comprovante.Transferencia = "debito"

		err = repoTransfer.InsertTransfer(&comprovante)
		if err != nil {
			return err
		}

		transfer.Transferencia = "credito"
		transfer.CPF = accountDestination.CPF
		err = repoTransfer.InsertTransfer(&transfer)
		if err != nil {
			return err
		}

		return err
	}
}

// Retorna transferencias
func GetTransfer(r *http.Request, limit int, page int) ([]domain.Transfer, error) {

	login, err := tokenJWT.ExtractTokenMetadata(r)
	if err != nil {
		return nil, err
	}

	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := infrastructure.TrasnferRepositoryWithRDB(conn)
	transfer, err := repo.ShowListTransfer(login.CPF)

	if err != nil {
		return nil, err
	}

	return transfer, nil
}
