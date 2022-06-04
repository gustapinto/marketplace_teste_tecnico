package models

import "gorm.io/gorm"

type Inventory struct {
	Total     int `json:"total" gorm:"estoque_total;notNull"`
	Cut       int `json:"corte" gorm:"estoque_corte;notNull"`
	Available int `json:"disponivel" gorm:"estoque_disponivel;notNull"`
}

type Product struct {
	gorm.Model
	Code      *string   `json:"codigo" gorm:"codigo;uniqueIndex;notNull"`
	Name      *string   `json:"nome" gorm:"nome;notNull"`
	Inventory Inventory `json:"estoque" gorm:"embedded;notNull"`
	PriceOf   float64   `json:"preco_de" gorm:"preco_de;notNull"`
	PriceFor  float64   `json:"preco_por" gorm:"preco_por;notNull"`
}
