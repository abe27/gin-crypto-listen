package models

import (
	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Exchange struct {
	ID          string `gorm:"size:21;primaryKey" form:"id" json:"id"`
	Exchange    string `gorm:"not null;unique;size:50" form:"exchange" json:"exchange" binding:"required"`
	Flag        string `gorm:"null;size:255" form:"flag" json:"flag"`
	Description string `gorm:"size:255" form:"description" json:"description"`
	IsActive    bool   `json:"is_active" form:"is_active" default:"false"`
	gorm.Model
}

func (u *Exchange) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
