package models

import (
	"time"
)

type Shop struct {
	ID        uint       `gorm:"primaryIndex" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Nama      string     `gorm:"type:varchar(255)" json:"nama"`
	Jalan     string     `gorm:"type:varchar(255)" json:"jalan"`
	Kecamatan string     `gorm:"type:varchar(255)" json:"kecamatan"`
	Provinsi  string     `gorm:"type:varchar(255)" json:"provinsi"`
	Dorayakis []Dorayaki `gorm:"many2many:shop_dorayakis" json:"inventory,omitempty"`
}
