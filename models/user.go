package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID         string    `gorm:"size:21;primaryKey" form:"id" json:"id"`
	Email      string    `gorm:"unique;index;not null;size:50" form:"email" json:"email" binding:"required"`
	Password   string    `gorm:"size:60" form:"password" json:"-" binding:"required"`
	IsVerified bool      `form:"is_verified" json:"is_verified" default:"false"`
	CreatedAt  time.Time `form:"created_at" json:"created_at" default:"now"`
	UpdatedAt  time.Time `form:"updated_at" json:"updated_at" default:"now"`
}

func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
