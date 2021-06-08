package domain

import (
	"time"
)

type Transfer struct {
	ID                   uint      `json:"id" gorm:"primaryKey autoIncrement"`
	CPF                  string    `json:"cpf"`
	Transferencia        string    `json:"transferencia"`
	AccountOriginId      int       `json:"account_origin_id"`
	AccountDestinationId int       `json:"account_destination_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}
