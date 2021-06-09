package infrastructure

import (
	"test-stone/domain"

	"github.com/jinzhu/gorm"
)

type TransferRepositoryImpl struct {
	Conn *gorm.DB
}

func TrasnferRepositoryWithRDB(conn *gorm.DB) TransferRepository {
	return &TransferRepositoryImpl{Conn: conn}
}

type TransferRepository interface {
	InsertTransfer(transfer *domain.Transfer) error
	ShowListTransfer(cpf string) ([]domain.Transfer, error)
}

//registra Transferencia
func (r *TransferRepositoryImpl) InsertTransfer(transfer *domain.Transfer) error {

	if err := r.Conn.Save(&transfer).Error; err != nil {
		return err
	}

	return nil
}

//retorna lista de transferencia
func (r *TransferRepositoryImpl) ShowListTransfer(cpf string) ([]domain.Transfer, error) {

	r.Conn.LogMode(true)
	transfer := []domain.Transfer{}
	if err := r.Conn.Model(&domain.Transfer{}).Where("cpf = ?", cpf).Find(&transfer).Error; err != nil {
		return nil, err
	}
	return transfer, nil
}
