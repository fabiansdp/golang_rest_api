package dto

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

type AddDorayakiInput struct {
	DorayakiID string `json:"dorayaki_id" binding:"required"`
	ShopID     string `json:"shop_id" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
}
