package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderId      uint       `gorm:"primaryKey" json:"order_id"`
	CustomerName string     `gorm:"not null" json:"customer_name" form:"customer_name" valid:"required~Your customer name is required"`
	OrderedAt    *time.Time `gorm:"not null" json:"ordered_at,omitempty"`
	ItemRefer    int        `json:"item_id"`
	Items        []Item     `gorm:"foreignKey:ItemRefer" json:"items"`
	// Item []Item `gorm:"constraint:OnUpdate:CASCADE" json:"items"`
}
