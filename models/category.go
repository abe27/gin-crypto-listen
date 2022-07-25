package models

import (
	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Category struct {
	ID          string `gorm:"size:21;primaryKey" form:"id" json:"id"`
	Category    string `gorm:"unique;size:50" form:"category" json:"category" binding:"required"`
	Description string `gorm:"size:255" form:"description" json:"description"`
	IsActive    bool   `form:"is_active" json:"is_active" default:"false"`
	gorm.Model
}

func (u *Category) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
