package models

import (
	n "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Interesting struct {
	ID          string  `gorm:"size:21;primaryKey" form:"id" json:"id"`
	ExchangeID  string  `gorm:"size:21" form:"exchange_id"`
	AssetID     string  `gorm:"size:21" form:"asset_id"`
	CurrencyID  string  `gorm:"size:21" form:"currency_id"`
	Price       float64 `gorm:"null" form:"price" default:"0"`
	Description string  `gorm:"size:255" form:"description"`
	IsNotifies  bool    `form:"is_notifies" default:"false"`
	IsActive    bool    `form:"is_active" default:"false"`
	gorm.Model
	Exchange Exchange `gorm:"foreignKeys:ExchangeID;references:ID"`
	Asset    Asset    `gorm:"foreignKeys:AssetID;references:ID"`
	Currency Currency `gorm:"foreignKeys:CurrencyID;references:ID"`
}

func (u *Interesting) BeforeCreate(tb *gorm.DB) (err error) {
	id, _ := n.New(21)
	u.ID = id
	return
}
