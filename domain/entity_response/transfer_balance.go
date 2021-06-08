package entityresponse

import "time"

type Transfer_Balance struct {
	ID                   uint      `json:"id" gorm:"primaryKey autoIncrement"`
	CPF                  string    `json:"cpf"`
	Transferencia        string    `json:"transferencia"`
	AccountOriginId      int       `json:"account_origin_id"`
	AccountDestinationId int       `json:"account_destination_id"`
	Amount               string    `json:"amount"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}
