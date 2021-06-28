package models

type ShopDorayaki struct {
	DorayakiID uint `gorm:"primaryKey" json:"dorayaki_id"`
	ShopID     uint `gorm:"primaryKey" json:"shop_id"`
	Quantity   int  `gorm:"type:int(10); not null; default:0" json:"quantity"`
}
