package domain

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func DoLogin(cpf string, secret string) *Login {
	login := &Login{
		CPF:    cpf,
		Secret: secret,
	}

	return login
}
