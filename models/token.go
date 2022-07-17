package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type ApiData struct {
	ID         string    `gorm:"size:21" form:"id"`
	UserID     string    `gorm:"size:21" form:"user_id"`
	ExchangeID string    `gorm:"size:21" form:"exchange_id"`
	SecretID   string    `gorm:"unique;size:60" form:"secret_id"`
	TokenID    string    `gorm:"unique;size:60" form:"token_id"`
	IsActive   bool      `gorm:"is_active" default:"false"`
	CreatedAt  time.Time `form:"created_at" default:"now"`
	UpdatedAt  time.Time `form:"updated_at" default:"now"`
	User       User      `gorm:"foreignKeys:UserID;references:ID"`
	Exchange   Exchange  `gorm:"foreignKeys:ExchangeID;references:ID"`
}

func (u *ApiData) BeforeCreate(t *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
