package models

import (
	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Investment struct {
	ID         string  `gorm:"size:21;primaryKey"`
	OrderNo    string  `gorm:"not null;unique;size:25" form:"order_no"`
	UserID     string  `gorm:"size:21" form:"user_id"`
	ExchangeID string  `gorm:"size:21" form:"exchange_id"`
	AssetID    string  `gorm:"size:21" form:"asset_id"`
	CurrencyID string  `gorm:"size:21" form:"currency_id"`
	Cost       float64 `gorm:"null" form:"cost" default:"0"`
	Invest     float64 `gorm:"null" form:"invest" default:"0"`
	Price      float64 `gorm:"null" form:"price" default:"0"`
	IsClosed   bool    `form:"is_closed" default:"false"`
	IsActive   bool    `form:"is_active" default:"false"`
	gorm.Model
	User     User     `gorm:"foreignKeys:UserID;references:ID"`
	Exchange Exchange `gorm:"foreignKeys:ExchangeID;references:ID"`
	Asset    Asset    `gorm:"foreignKeys:AssetID;references:ID"`
	Currency Currency `gorm:"foreignKeys:CurrencyID;references:ID"`
}

func (u *Investment) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
