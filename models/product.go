package models

type Product struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"varchar(255)" json:"nama_product"`
	Description string `gorm:"text" json:"deskripsi"`
}