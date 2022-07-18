package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Category struct {
	ID          string    `gorm:"size:21;primaryKey" form:"id" json:"id"`
	Name        string    `gorm:"unique;size:50" form:"name" json:"name" binding:"required"`
	Description string    `gorm:"size:255" form:"description" json:"description"`
	IsActive    bool      `form:"is_active" json:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" json:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at" default:"now"`
}

func (u *Category) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
