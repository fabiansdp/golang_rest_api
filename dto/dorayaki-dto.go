package dto

type CreateDorayakiInput struct {
	Rasa      string `json:"rasa" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Gambar    string `json:"gambar" binding:"required"`
}

type UpdateDorayakiInput struct {
	Rasa      string `json:"rasa"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
}
