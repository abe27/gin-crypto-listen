package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Cryptocurrency struct {
	ID          string    `gorm:"size:21" form:"id"`
	Name        string    `gorm:"unique;size:50" form:"name"`
	Symbol      string    `gorm:"size:10" form:"symbol"`
	Address     string    `gorm:"size:60" form:"address"`
	Flag        string    `gorm:"null;size:255" form:"flag"`
	Description string    `gorm:"size:255" form:"description"`
	IsActive    bool      `form:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" default:"now"`
}

func (u *Cryptocurrency) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
