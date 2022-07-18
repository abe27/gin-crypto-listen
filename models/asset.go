package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Asset struct {
	ID             string         `gorm:"size:21;primaryKey" form:"id"`
	CategoryID     string         `gorm:"size:21" form:"category_id"`
	CryptoID       string         `gorm:"size:21" form:"crypto_id"`
	Description    string         `gorm:"size:255" form:"description"`
	IsActive       bool           `form:"is_active" default:"false"`
	CreatedAt      time.Time      `form:"created_at" default:"now"`
	UpdatedAt      time.Time      `form:"updated_at" default:"now"`
	Category       Category       `gorm:"foreignKey:CategoryID;references:ID"`
	Cryptocurrency Cryptocurrency `gorm:"foreignKey:CryptoID;references:ID"`
}

func (u *Asset) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
