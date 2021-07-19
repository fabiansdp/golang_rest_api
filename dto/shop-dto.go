package dto

import "github.com/fabiansdp/golang_rest_api/models"

type CreateShopInput struct {
	Nama      string `json:"nama" binding:"required"`
	Jalan     string `json:"jalan" binding:"required"`
	Kecamatan string `json:"kecamatan" binding:"required"`
	Provinsi  string `json:"provinsi" binding:"required"`
}

type UpdateShopInput struct {
	Nama      string `json:"nama"`
	Jalan     string `json:"jalan"`
	Kecamatan string `json:"kecamatan"`
	Provinsi  string `json:"provinsi"`
}

type Inventory struct {
	models.Dorayaki
	Quantity int `json:"quantity"`
}

type GetShopOutput struct {
	ShopInfo  models.Shop `json:"shop_info"`
	Inventory []Inventory `json:"inventory"`
}

type AddInventoryInput struct {
	DorayakiID uint `json:"dorayaki_id" binding:"required"`
	ShopID     uint `json:"shop_id" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required"`
}

type UpdateInventoryInput struct {
	DorayakiID uint `json:"dorayaki_id" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required"`
}

type MoveInventoryInput struct {
	DorayakiID  uint `json:"dorayaki_id" binding:"required"`
	RecipientID uint `json:"recipient_id" binding:"required"`
	Quantity    int  `json:"quantity" binding:"required"`
}
