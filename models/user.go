package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID         string    `gorm:"size:21" form:"id"`
	Email      string    `gorm:"unique;index;not null;size:50" form:"email"`
	Password   string    `gorm:"size:60" form:"password" json:"-"`
	IsVerified bool      `form:"is_verified" default:"false"`
	CreatedAt  time.Time `form:"created_at" default:"now"`
	UpdatedAt  time.Time `form:"updated_at" default:"now"`
}

func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
