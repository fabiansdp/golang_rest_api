package models

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	Nama      string     `json:"nama"`
	Jalan     string     `json:"jalan"`
	Kecamatan string     `json:"kecamatan"`
	Provinsi  string     `json:"provinsi"`
	Dorayaki  []Dorayaki `json:"dorayaki"`
}
