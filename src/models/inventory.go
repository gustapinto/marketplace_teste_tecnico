package models

import "gorm.io/gorm"

type Inventory struct {
	Total     int `json:"total" gorm:"estoque_total"`
	Cut       int `json:"corte" gorm:"estoque_corte"`
	Available int `json:"disponivel" gorm:"estoque_disponivel"`
}

type Product struct {
	gorm.Model
	Code      string    `json:"codigo" gorm:"codigo"`
	Name      string    `json:"nome" gorm:"nome"`
	Inventory Inventory `json:"inventory" gorm:"embedded"`
	PriceOf   float64   `json:"preco_de" gorm:"preco_de"`
	PriceFor  float64   `json:"preco_por" gorm:"preco_por"`
}
