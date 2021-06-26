package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Nama      string     `gorm:"type:varchar(255)" json:"nama"`
	Jalan     string     `gorm:"type:varchar(255)" json:"jalan"`
	Kecamatan string     `gorm:"type:varchar(255)" json:"kecamatan"`
	Provinsi  string     `gorm:"type:varchar(255)" json:"provinsi"`
	Dorayakis []Dorayaki `gorm:"many2many:shop_dorayakis" json:"inventory"`
}
