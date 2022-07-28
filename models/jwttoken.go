package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type JwtToken struct {
	ID        string    `gorm:"index;size:21;primaryKey"`
	Key       string    `gorm:"size:65" binding:"required"`
	UserID    string    `gorm:"unique;not null;size:21" binding:"required"`
	Token     string    `gorm:"size:255" binding:"required"`
	IsActive  bool      `default:"true"`
	CreatedAt time.Time `form:"created_at" default:"now"`
	UpdatedAt time.Time `form:"updated_at" default:"now"`
	User      User      `gorm:"foreignKeys:UserID;references:ID"`
}

func (u *JwtToken) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	TokenID, _ := g.New(56)
	u.ID = id
	u.Key = TokenID
	u.IsActive = true
	return
}
