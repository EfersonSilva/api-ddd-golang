package interfaces

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"test-stone/application"
	"test-stone/domain"
	entityresponse "test-stone/domain/entity_response"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//Login
	r.POST("/api/v1/login", Login)

	// Account
	r.GET("/api/v1/accounts", GetAllAccounts)
	r.GET("/api/v1/accounts/:id", GetAccountById)
	r.GET("/api/v1/account/:account_id/balance", GetBalanceAccountById)
	r.POST("/api/V1/accounts", CreateNewAccount)
	r.PUT("/api/v1/accounts", UpdateAccount)
	r.DELETE("/api/v1/accounts/:id", DeleteAccount)

	//Transfer
	r.POST("/api/v1/transfers", RegisterTransfer)
	r.GET("/api/v1/transfers", ShowTrasnfers)

	return r
}

//retorna todas as contas registradas
func GetAllAccounts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if !application.AuthLogin(r) {
		Error(w, http.StatusUnauthorized, errors.New("unauthorized"), " unauthorized")
		return
	}

	queryValues := r.URL.Query()

	limit := queryValues.Get("limit")
	page := queryValues.Get("page")

	//se vier com dados para paginação aciona esse bloco
	if limit != "" && page != "" {
		limit, _ := strconv.Atoi(limit)
		page, _ := strconv.Atoi(page)

		if limit != 0 && page != 0 {
			news, err := application.GetAllAccounts(limit, page)
			if err != nil {
				Error(w, http.StatusNotFound, err, err.Error())
				return
			}

			JSON(w, http.StatusOK, news)
			return
		}
	}
	acc, err := application.GetAllAccounts(15, 1) // 15, 1 Paginação
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, acc)
}

func GetAccountById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if !application.AuthLogin(r) {
		Error(w, http.StatusUnauthorized, errors.New("unauthorized"), " unauthorized")
		return
	}

	param := ps.ByName("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		Error(w, http.StatusOK, err, err.Error())
		return
	}

	account, err := application.GetAccountById(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, account)
}

func GetBalanceAccountById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if !application.AuthLogin(r) {
		Error(w, http.StatusUnauthorized, errors.New("unauthorized"), " unauthorized")
		return
	}

	param := ps.ByName("account_id")
	id, err := strconv.Atoi(param)

	if err != nil {
		Error(w, http.StatusBadRequest, err, err.Error())
		return
	}

	balance, err := application.GetBalanceAccountById(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, balance)

}

func UpdateAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if !application.AuthLogin(r) {
		Error(w, http.StatusUnauthorized, errors.New("unauthorized"), " unauthorized")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var account domain.Account
	err := decoder.Decode(&account)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	err = application.UpdateAccount(account)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

func CreateNewAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var account domain.Account

	if err := decoder.Decode(&account); err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	acc, _ := domain.NewAccount(account.Name, account.CPF, account.Secret, account.Balance)

	err := application.CreateAccount(*acc)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	acc.Secret = ""

	JSON(w, http.StatusCreated, acc)
}

//Exclui conta
func DeleteAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if !application.AuthLogin(r) {
		Error(w, http.StatusUnauthorized, errors.New("unauthorized"), " unauthorized")
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.DeleteAccount(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	decoder := json.NewDecoder(r.Body)

	var login domain.Login

	if err := decoder.Decode(&login); err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	var token entityresponse.Token
	resp, err := application.DoLogin(&login)
	if err != nil {
		JSON(w, http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token.Value = *resp
	JSON(w, http.StatusOK, token)
}

//Registra transferencia
func RegisterTransfer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var transfer domain.Transfer

	if err := decoder.Decode(&transfer); err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	err := application.AddTransfer(transfer)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusCreated, nil)
}

func ShowTrasnfers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if !application.AuthLogin(r) {
		Error(w, http.StatusUnauthorized, errors.New("unauthorized"), " unauthorized")
		return
	}

	queryValues := r.URL.Query()
	limit := queryValues.Get("limit")
	page := queryValues.Get("page")

	//se vier com dados para paginação aciona esse bloco
	if limit != "" && page != "" {
		limit, _ := strconv.Atoi(limit)
		page, _ := strconv.Atoi(page)

		if limit != 0 && page != 0 {
			transfers, err := application.GetTransfer(r, limit, page)
			if err != nil {
				Error(w, http.StatusNotFound, err, err.Error())
				return
			}

			JSON(w, http.StatusOK, transfers)
			return
		}
	}

	transfers, err := application.GetTransfer(r, 15, 1) // 15, 1 Paginação
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}
	JSON(w, http.StatusOK, transfers)
}
