package models

import "gorm.io/gorm"

type Dorayaki struct {
	gorm.Model
	Rasa      string `json:"rasa"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
	TokoID    uint
}
