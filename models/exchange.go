package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Exchange struct {
	ID          string    `gorm:"size:21" form:"id"`
	Name        string    `gorm:"unique;size:50" form:"name"`
	Description string    `gorm:"size:255" form:"description"`
	IsActive    bool      `form:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" default:"now"`
}

func (u *Exchange) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
