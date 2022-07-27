package models

import (
	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Category struct {
	ID          string `gorm:"size:21;primaryKey"`
	Category    string `gorm:"unique;size:50" form:"category" binding:"required"`
	Description string `gorm:"size:255" form:"description"`
	IsActive    bool   `form:"is_active" default:"false"`
	gorm.Model
}

func (u *Category) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
