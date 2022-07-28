package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Asset struct {
	ID             string         `gorm:"size:21;primaryKey"`
	CategoryID     string         `form:"category"`
	CryptoID       string         `form:"crypto"`
	Description    string         `gorm:"size:255" form:"description"`
	IsActive       bool           `form:"is_active" default:"false"`
	Category       Category       `gorm:"foreignKey:CategoryID"`
	Cryptocurrency Cryptocurrency `gorm:"foreignKey:CryptoID"`
	CreatedAt      time.Time      `form:"created_at" default:"now"`
	UpdatedAt      time.Time      `form:"updated_at" default:"now"`
}

func (u *Asset) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
