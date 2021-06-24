package models

import "gorm.io/gorm"

type Dorayaki struct {
	gorm.Model
	Rasa      string `gorm:"type:varchar(255)" json:"rasa"`
	Deskripsi string `gorm:"type:varchar(255)" json:"deskripsi"`
	Gambar    string `gorm:"type:varchar(255)" json:"gambar"`
}
