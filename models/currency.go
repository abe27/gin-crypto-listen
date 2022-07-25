package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Currency struct {
	ID          string    `gorm:"size:21;primaryKey" form:"id" json:"id"`
	Symbol      string    `gorm:"unique;size:50" form:"symbol" json:"symbol"`
	Flag        string    `gorm:"null;size:255" form:"flag" json:"flag"`
	Description string    `gorm:"size:255" form:"description" json:"description"`
	IsActive    bool      `form:"is_active" json:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" json:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at" default:"now"`
}

func (u *Currency) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
