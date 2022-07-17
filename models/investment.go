package models

import (
	"time"

	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Investment struct {
	ID         string    `gorm:"size:21" form:"id"`
	UserID     string    `gorm:"size:21" form:"user_id"`
	User       User      `gorm:"foreignKeys:UserID;references:ID"`
	ExchangeID string    `gorm:"size:21" form:"exchange_id"`
	Exchange   Exchange  `gorm:"foreignKeys:ExchangeID;references:ID"`
	AssetID    string    `gorm:"size:21" form:"asset_id"`
	Asset      Asset     `gorm:"foreignKeys:AssetID;references:ID"`
	CurrencyID string    `gorm:"size:21" form:"currency_id"`
	Currency   Currency  `gorm:"foreignKeys:CurrencyID;references:ID"`
	Cost       float64   `gorm:"null" form:"cost" default:"0"`
	Invest     float64   `gorm:"null" form:"invest" default:"0"`
	Price      float64   `gorm:"null" form:"price" default:"0"`
	IsStatus   bool      `form:"is_status" default:"false"`
	IsActive   bool      `form:"is_active" default:"false"`
	CreatedAt  time.Time `form:"created_at" default:"now"`
}

func (u *Investment) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
