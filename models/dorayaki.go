package models

type Dorayaki struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Rasa      string `gorm:"type:varchar(255)" json:"rasa"`
	Deskripsi string `gorm:"type:varchar(255)" json:"deskripsi"`
	Gambar    string `gorm:"type:varchar(255)" json:"gambar"`
}
