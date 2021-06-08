package infrastructure

import (
	"test-stone/domain"

	"github.com/jinzhu/gorm"
)

type AccountRepositoryImpl struct {
	Conn *gorm.DB
}

type AccountRepository interface {
	InsertAccount(account *domain.Account) error
	FindAccountById(id int) (*domain.Account, error)
	FindAccountByCpf(cpf string) (*domain.Account, error)
	FindAllAccount() ([]domain.Account, error)
	UpdateAccount(acc *domain.Account) error
	DeleteAccount(id int) error
	LoginRepository(cpf string) (*domain.Login, error)
}

func AccountRepositoryWithRDB(conn *gorm.DB) AccountRepository {
	return &AccountRepositoryImpl{Conn: conn}
}

//adicionaa conta
func (r *AccountRepositoryImpl) InsertAccount(account *domain.Account) error {

	if err := r.Conn.Save(&account).Error; err != nil {
		return err
	}

	return nil
}

//retorna conta por id
func (r *AccountRepositoryImpl) FindAccountById(id int) (*domain.Account, error) {
	account := &domain.Account{}
	if err := r.Conn.Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	account.Secret = ""
	return account, nil
}

//retorna conta por cpf
func (r *AccountRepositoryImpl) FindAccountByCpf(cpf string) (*domain.Account, error) {
	account := &domain.Account{}
	if err := r.Conn.Where("cpf = ?", cpf).First(&account).Error; err != nil {
		return nil, err
	}
	account.Secret = ""
	return account, nil
}

//retorna todas contas
func (r *AccountRepositoryImpl) FindAllAccount() ([]domain.Account, error) {
	accounts := []domain.Account{}
	if err := r.Conn.Preload("Topic").Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

// Update is update news
func (r *AccountRepositoryImpl) UpdateAccount(acc *domain.Account) error {
	if err := r.Conn.Model(&acc).Update(acc).Error; err != nil {
		return err
	}
	return nil
}

// exclui conta
func (r *AccountRepositoryImpl) DeleteAccount(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	news := domain.Account{}
	if err := tx.First(&news, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&news).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//retorna conta por id
func (r *AccountRepositoryImpl) LoginRepository(cpf string) (*domain.Login, error) {
	acc := &domain.Account{}
	if err := r.Conn.Where("cpf = ?", cpf).First(&acc).Error; err != nil {
		return nil, err
	}
	var login domain.Login
	login.CPF = acc.CPF
	login.Secret = acc.Secret

	return &login, nil
}
