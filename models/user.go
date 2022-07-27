package models

import (
	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID         string `gorm:"size:21;primaryKey"`
	Email      string `gorm:"unique;index;not null;size:50" form:"email" binding:"required"`
	Password   string `gorm:"size:60" form:"password" json:"-" binding:"required"`
	IsVerified bool   `form:"is_verified" default:"false"`
	gorm.Model
}

func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
