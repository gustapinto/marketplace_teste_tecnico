package models

import "gorm.io/gorm"

type Inventory struct {
	Total     int `json:"total" gorm:"notNull"`
	Cut       int `json:"corte" gorm:"notNull"`
	Available int `json:"disponivel" gorm:"notNull"`
}

type Product struct {
	gorm.Model
	Code      *string   `json:"codigo" gorm:"uniqueIndex;notNull"`
	Name      *string   `json:"nome" gorm:"notNull"`
	Inventory Inventory `json:"estoque" gorm:"embedded;embeddedPrefix:inventory_;notNull"`
	PriceOf   float64   `json:"preco_de" gorm:"notNull"`
	PriceFor  float64   `json:"preco_por" gorm:"notNull"`
}
