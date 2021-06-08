package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey autoIncrement"`
	Name      string         `json:"name"`
	CPF       string         `json:"cpf" gorm:"unique"`
	Secret    string         `json:"secret,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
