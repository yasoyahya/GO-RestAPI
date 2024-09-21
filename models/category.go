package models

type Category struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaCategory string `gorm:"type:varchar(300)" json:"nama_category"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
}